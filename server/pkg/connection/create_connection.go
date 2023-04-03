package connection

import (
	"context"
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
	ctx              context.Context
	cancleFunc       context.CancelCauseFunc
	commsChan        chan any
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

	playersConnected := players.CreateConnectedPlayers()
	message := message.InitMessage(message.MessageOptions{gameServerOptions.Logger, playersConnected})

	gameServerOptions.Logger.Log(tcpAddress.String())

	room := room.CreateRooms(2, message)
	commsChan := make(chan any)

	cancelCtx, cancelFunc := context.WithCancelCause(context.Background())

	return &GameServer{
		listener:         tcpListener,
		logger:           gameServerOptions.Logger,
		playersConnected: playersConnected,
		rooms:            room.Rooms,
		message:          message,
		handlers: handlers.InitHandlers(handlers.HandlerOption{
			Room:          *room,
			CancelCtxFunc: cancelFunc,
			CommsChan:     commsChan,
		}),
		ctx:        cancelCtx,
		cancleFunc: cancelFunc,
		commsChan:  commsChan,
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

			clientAddress := newUserConnection.RemoteAddr().String()

			gameServer.logger.Log(fmt.Sprintf("%+v", newUserConnection))

			gameServer.playersConnected.AddPlayer(newUserConnection, clientAddress)
			gameServer.message.Unicast.SendRooms(newUserConnection, gameServer.rooms)

			for {

				readCount, err := newUserConnection.Read(c)

				if errors.Is(err, io.EOF) {
					break
				}

				command, payload := gameServer.message.ProcessMessage(string(c[:readCount]))

				gameServer.logger.Log(fmt.Sprintf("Handling command (%s) from address (%s)", command, clientAddress))

				go gameServer.handlers.HandleCommand(
					command,
					payload,
					clientAddress,
					newUserConnection,
				)

				select {
				case <-gameServer.ctx.Done():
					errorCause := context.Cause(gameServer.ctx)
					if err := errorCause; err != nil && !errors.Is(err, context.Canceled) {
						gameServer.logger.LogWithCtx(gameServer.ctx, errorCause.Error(), clientAddress, []any{command, clientAddress})
					}
				}

			}

		}()
	}
}
