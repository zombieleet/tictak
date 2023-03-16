// client connection package
package connection

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/zombieleet/tictak/client/pkg/handler"
	"github.com/zombieleet/tictak/client/pkg/message"
	"net"
)

type GameClientOptions struct {
	HostName string
	Port     uint32
}

type GameClient struct {
	dialer  *net.TCPConn
	message *message.Message
	handler *handler.Handler

	ctx         context.Context
	cancelCause context.CancelCauseFunc
	commsChan   chan string
}

// CreageGameClient
// establishes a connection with the gameserver while setting up other
// packages needed by the gameclient
func CreateGameClient(gameClientOptions GameClientOptions) (*GameClient, error) {
	address := fmt.Sprintf("%s:%d", gameClientOptions.HostName, gameClientOptions.Port)

	tcpAddress, err := net.ResolveTCPAddr("tcp", address)

	if err != nil {
		return nil, err
	}

	tcpDialer, err := net.DialTCP("tcp", nil, tcpAddress)

	if err != nil {
		return nil, err
	}

	ctx, cancelCtxCause := context.WithCancelCause(context.Background())
	commsChan := make(chan string)

	return &GameClient{
		dialer:      tcpDialer,
		message:     message.NewMessage(),
		handler:     handler.InitHandlers(ctx, cancelCtxCause, commsChan),
		ctx:         ctx,
		cancelCause: cancelCtxCause,
		commsChan:   commsChan,
	}, nil
}

func (gameClient *GameClient) Connect() {

	reader := bufio.NewReader(gameClient.dialer)

	defer gameClient.dialer.Close()

	for {

		data, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			continue
		}

		command, payload := gameClient.message.ProcessMessage(data)

		go gameClient.handler.HandleUICommand(command, payload)

		select {
		case <-gameClient.ctx.Done():
			if err := gameClient.ctx.Err(); err != nil && !errors.Is(err, context.Canceled) {
				panic(context.Cause(gameClient.ctx))
			}
			return
		case payload := <-gameClient.commsChan:
			gameClient.dialer.Write([]byte(payload))
		}
	}
}
