package message

import (
	"context"
	"github.com/zombieleet/tictak/server/pkg/command"
	"github.com/zombieleet/tictak/server/pkg/logger"
	"net"
)

type Command string

const (
	CMD_SEND_ROOMS Command = "SEND_ROOMS"
)

type unicast struct {
	logger *logger.Logger
}

// SendRooms
// Sends the list of rooms and their avaialability information to
// a client
func (ucast *unicast) SendRoomsInfoToNewClient(conn net.Conn, payload string) {

	connAddress := conn.RemoteAddr().String()
	parentCtx := context.Background()

	valueCtx := context.WithValue(
		parentCtx,
		"SEND_ROOMS",
		encodeContext(messageContext{
			Type:         "SEND_ROOMS",
			ContextValue: connAddress,
		}),
	)

	ucast.logger.LogWithCtx(
		valueCtx,
		"sending rooms payload",
		"ROOMS",
		[]any{"address", connAddress},
	)

	_, error := conn.Write([]byte(command.Commands["CMD_SEND_ROOMS"] + payload + "\n"))

	if error != nil {
		ucast.logger.NetworkError(error)
		return
	}

	ucast.logger.LogWithCtx(
		valueCtx,
		"finish sending rooms payload",
		"ROOMS",
		[]any{"address", connAddress},
	)
}

func (ucast *unicast) SendPlayerTwoConnectedInfoFromPlayerOne() {}
