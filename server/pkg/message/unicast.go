package message

import (
	"context"
	"fmt"
	"github.com/zombieleet/tictak/server/pkg/commands"
	"github.com/zombieleet/tictak/server/pkg/logger"
	"github.com/zombieleet/tictak/server/pkg/room"
	"net"
	"strings"
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
func (ucast *unicast) SendRooms(conn net.Conn, rooms *room.RoomMap) {

	var payload strings.Builder

	for roomId, roomInfo := range *rooms {
		payload.WriteString(fmt.Sprintf("%d. %s (Occupied=%t)_", roomId, roomInfo.Name, roomInfo.Occupied))
	}

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

	_, error := conn.Write([]byte(commands.Commands["CMD_SEND_ROOMS"] + payload.String() + "\n"))

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
