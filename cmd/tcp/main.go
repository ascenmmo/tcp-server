package main

import (
	"context"
	"github.com/ascenmmo/tcp-server/env"
	"github.com/ascenmmo/tcp-server/pkg/start"
	"github.com/rs/zerolog"
	"os"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	ctx := context.Background()

	err := start.StartTCP(
		ctx,
		env.ServerAddress,
		env.TCPPort,
		env.TokenKey,
		env.MaxRequestPerSecond,
		10,
		logger,
		true,
	)

	if err != nil {
		logger.Fatal().Err(err).Msg("failed to start tcp server")
	}
}
