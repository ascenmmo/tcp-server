// GENERATED BY 'T'ransport 'G'enerator. DO NOT EDIT.
package tcpGameServer

type fallbackServerSettings interface {
	GetConnectionsNum(err error) bool
	HealthCheck(err error) bool
	GetServerSettings(err error) bool
	CreateRoom(err error) bool
	GetDeletedRooms(err error) bool
}
