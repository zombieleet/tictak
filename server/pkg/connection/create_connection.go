package connection

import (
	"errors"
	"fmt"
	"github.com/zombieleet/tictak/server/pkg/logger"
	"github.com/zombieleet/tictak/server/pkg/message"
	"github.com/zombieleet/tictak/server/pkg/room"
	"net"
)

type GameServerOptions struct {
	HostName string
	Port     int32
	Logger   *logger.Logger
}

type GameServer struct {
	listener *net.TCPListener
	logger   *logger.Logger
	rooms    *room.Room
	message  *message.Message
}

func CreateGameServer(gameServerOptions GameServerOptions) *GameServer {
	address := fmt.Sprintf("%s:%d", gameServerOptions.HostName, gameServerOptions.Port)
	tcpAddress, error := net.ResolveTCPAddr("tcp", address)

	if error != nil {
		gameServerOptions.Logger.NetworkError(errors.Join(E_NOT_RESOLVED, error))
		return nil
	}

	tcpListener, error := net.ListenTCP("tcp", tcpAddress)

	if error != nil {
		gameServerOptions.Logger.NetworkError(error)
		return nil
	}

	gameServerOptions.Logger.Log(tcpAddress.String())

	return &GameServer{
		listener: tcpListener,
		logger:   gameServerOptions.Logger,
		rooms:    room.CreateRooms(2),
		message:  message.InitMessage(message.MessageOptions{gameServerOptions.Logger}),
	}
}

func (gameServer *GameServer) Start() {

	for {

		newUserConnection, error := gameServer.listener.Accept()

		if error != nil {
			gameServer.logger.NetworkError(errors.Join(E_NOT_ACCEPT, error))
			continue
		}

		go func() {
			gameServer.logger.Log(fmt.Sprintf("%+v", newUserConnection))
			gameServer.logger.Log(fmt.Sprintf("%+v", gameServer.rooms))

			gameServer.message.Unicast.SendRooms(newUserConnection, gameServer.rooms)
		}()
	}
}
