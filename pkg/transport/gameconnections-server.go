// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package transport

import (
	"context"
	"github.com/ascenmmo/tcp-server/pkg/restconnection"
	"github.com/ascenmmo/tcp-server/pkg/restconnection/types"
	"github.com/google/uuid"
)

type serverGameConnections struct {
	svc            restconnection.GameConnections
	setSendMessage GameConnectionsSetSendMessage
	getMessage     GameConnectionsGetMessage
	removeUser     GameConnectionsRemoveUser
}

type MiddlewareSetGameConnections interface {
	Wrap(m MiddlewareGameConnections)
	WrapSetSendMessage(m MiddlewareGameConnectionsSetSendMessage)
	WrapGetMessage(m MiddlewareGameConnectionsGetMessage)
	WrapRemoveUser(m MiddlewareGameConnectionsRemoveUser)

	WithTrace()
	WithMetrics()
	WithLog()
}

func newServerGameConnections(svc restconnection.GameConnections) *serverGameConnections {
	return &serverGameConnections{
		getMessage:     svc.GetMessage,
		removeUser:     svc.RemoveUser,
		setSendMessage: svc.SetSendMessage,
		svc:            svc,
	}
}

func (srv *serverGameConnections) Wrap(m MiddlewareGameConnections) {
	srv.svc = m(srv.svc)
	srv.setSendMessage = srv.svc.SetSendMessage
	srv.getMessage = srv.svc.GetMessage
	srv.removeUser = srv.svc.RemoveUser
}

func (srv *serverGameConnections) SetSendMessage(ctx context.Context, token string, message types.RequestSetMessage) (err error) {
	return srv.setSendMessage(ctx, token, message)
}

func (srv *serverGameConnections) GetMessage(ctx context.Context, token string) (messages types.ResponseGetMessage, err error) {
	return srv.getMessage(ctx, token)
}

func (srv *serverGameConnections) RemoveUser(ctx context.Context, token string, userID uuid.UUID) (err error) {
	return srv.removeUser(ctx, token, userID)
}

func (srv *serverGameConnections) WrapSetSendMessage(m MiddlewareGameConnectionsSetSendMessage) {
	srv.setSendMessage = m(srv.setSendMessage)
}

func (srv *serverGameConnections) WrapGetMessage(m MiddlewareGameConnectionsGetMessage) {
	srv.getMessage = m(srv.getMessage)
}

func (srv *serverGameConnections) WrapRemoveUser(m MiddlewareGameConnectionsRemoveUser) {
	srv.removeUser = m(srv.removeUser)
}

func (srv *serverGameConnections) WithTrace() {
	srv.Wrap(traceMiddlewareGameConnections)
}

func (srv *serverGameConnections) WithMetrics() {
	srv.Wrap(metricsMiddlewareGameConnections)
}

func (srv *serverGameConnections) WithLog() {
	srv.Wrap(loggerMiddlewareGameConnections())
}