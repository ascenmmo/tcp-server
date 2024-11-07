package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ascenmmo/tcp-server/env"
	"github.com/ascenmmo/tcp-server/pkg/api/types"
	"github.com/ascenmmo/tcp-server/pkg/clients/tcpGameServer"
	"github.com/ascenmmo/tcp-server/pkg/start"
	tokengenerator "github.com/ascenmmo/token-generator/token_generator"
	tokentype "github.com/ascenmmo/token-generator/token_type"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

var (
	clients = 1000
	msgs    = 10

	doRead  = true
	baseURl = fmt.Sprintf("http://%s:%s", env.ServerAddress, env.TCPPort)
	token   = env.TokenKey
)

var ctx, cancel = context.WithCancel(context.Background())
var min, max time.Duration
var maxMsgs int

type Message struct {
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}

type Request struct {
	Token string  `json:"token,omitempty"`
	Data  Message `json:"data,omitempty"`
}

type Response struct {
	Data Message
}

func TestConnection(t *testing.T) {
	//logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	logger := zerolog.Logger{}
	go start.StartTCP(
		context.Background(),
		env.ServerAddress,
		env.TCPPort,
		env.TokenKey,
		env.MaxRequestPerSecond,
		2,
		logger,
		false,
	)
	time.Sleep(time.Second * 2)

	for i := 0; i < clients; i++ {
		createRoom(t, createToken(t, i))
		go Listener(t, i)
		go Publisher(t, i)
	}
	<-ctx.Done()
	time.Sleep(time.Second * 5)

	fmt.Println(max, min, maxMsgs)
}

func Publisher(t *testing.T, i int) {
	for j := 0; j < msgs; j++ {
		if ctx.Err() != nil {
			return
		}
		msg := buildMessage(t, i, j)
		cli := tcpGameServer.New(baseURl)
		err := cli.GameConnections().SetSendMessage(ctx, createToken(t, i), msg)
		assert.Nil(t, err, "client.do expected nil")
	}
}

func Listener(t *testing.T, i int) {
	_ = listen(t, createToken(t, i))
	//fmt.Println("done pubSub", i, "with msgs", response)
	cancel()
}

func createToken(t *testing.T, i int) string {
	//z := 0
	//if i > clients/2 {
	//	z = 1
	//}
	gameID := uuid.NewMD5(uuid.UUID{}, []byte(strconv.Itoa(1)))
	roomID := uuid.NewMD5(uuid.UUID{}, []byte(strconv.Itoa(i)))
	userID := uuid.New()

	tokenGen, err := tokengenerator.NewTokenGenerator(token)
	assert.Nil(t, err, "init gen token expected nil")

	token, err := tokenGen.GenerateToken(tokentype.Info{
		GameID: gameID,
		RoomID: roomID,
		UserID: userID,
		TTL:    time.Second * 100,
	}, tokengenerator.AESGCM)
	assert.Nil(t, err, "gen token expected nil")

	return token
}

func createRoom(t *testing.T, userToken string) {
	cli := tcpGameServer.New(baseURl)
	_ = cli.ServerSettings().CreateRoom(context.Background(), userToken, types.CreateRoomRequest{})
}

func buildMessage(t *testing.T, i, j int) (reqMsg types.RequestSetMessage) {
	data := Message{
		Text:      fmt.Sprintf("отправила горутина:%d \t номер сообщения: %d", i, j),
		CreatedAt: time.Now(),
	}
	if j == msgs-1 {
		data = Message{
			Text:      "close",
			CreatedAt: time.Now(),
		}
	}

	reqMsg.Data = data
	return reqMsg
}

func listen(t *testing.T, token string) int {
	for {

		cli := tcpGameServer.New(baseURl)
		messages, err := cli.GameConnections().GetMessage(context.Background(), token)
		if err != nil {
			return 0
			//t.Error(err)
		}
		counter := len(messages.DataArray)
		maxMsgs += counter

		marshal, err := json.Marshal(messages.DataArray)
		assert.Nil(t, err, "client.do expected nil")

		r := []Message{}
		err = json.Unmarshal(marshal, &r)
		assert.Nil(t, err, "Unmarshal expected nil")
		for _, v := range r {
			if v.Text == "close" {
				return counter
			}

			sub := time.Now().Sub(v.CreatedAt)
			if min == 0 {
				min = sub
			}
			if min > sub {
				min = sub
			}

			if max < sub {
				max = sub
			}
		}
		time.Sleep(time.Second * 1)
	}
}
