package connection

import (
	"errors"
	"fmt"
	"github.com/zombieleet/tictak/server/pkg/handlers"
	"github.com/zombieleet/tictak/server/pkg/logger"
	"github.com/zombieleet/tictak/server/pkg/message"
	"github.com/zombieleet/tictak/server/pkg/players"
	"github.com/zombieleet/tictak/server/pkg/room"
	"io"
	"net"
)

type GameServerOptions struct {
	HostName string
	Port     uint32
	Logger   *logger.Logger
}

type GameServer struct {
	listener         *net.TCPListener
	logger           *logger.Logger
	rooms            *room.RoomMap
	playersConnected *players.PlayersConnected
	message          *message.Message
	handlers         *handlers.Handler
}

// CreateGameServer
// creates the game server and returns a pointer that no public fields
// the private fields of `GameServer` will be used internally within the
// `conection` package
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

	room := room.CreateRooms(2)

	return &GameServer{
		listener:         tcpListener,
		logger:           gameServerOptions.Logger,
		playersConnected: players.CreateConnectedPlayers(),
		rooms:            room.Rooms,
		message:          message.InitMessage(message.MessageOptions{gameServerOptions.Logger}),
		handlers:         handlers.InitHandlers(handlers.HandlerOption{*room}),
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
			c := make([]byte, 1000)

			gameServer.logger.Log(fmt.Sprintf("%+v", newUserConnection))

			gameServer.playersConnected.AddPlayer(newUserConnection.RemoteAddr().String())
			gameServer.message.Unicast.SendRooms(newUserConnection, gameServer.rooms)

			clientAddress := newUserConnection.RemoteAddr().String()

			for {

				readCount, err := newUserConnection.Read(c)

				if errors.Is(err, io.EOF) {
					break
				}

				command, payload := gameServer.message.ProcessMessage(string(c[:readCount]))

				gameServer.logger.Log(fmt.Sprintf("Handling command (%s) from address (%s)", command, clientAddress))

				go gameServer.handlers.HandleCommand(command, payload, clientAddress)
			}

		}()
	}
}
