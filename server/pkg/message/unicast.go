package message

import (
	"context"
	"fmt"
	"github.com/zombieleet/tictak/server/pkg/logger"
	"github.com/zombieleet/tictak/server/pkg/room"
	"net"
	"strings"
)

type unicast struct {
	logger *logger.Logger
}

func (ucast *unicast) SendRooms(conn net.Conn, rooms *room.Room) {

	var payload strings.Builder

	for roomId, roomInfo := range *rooms {
		payload.WriteString(fmt.Sprintf("%d. %s (Occupied=%t)\n", roomId, roomInfo.Name, roomInfo.Occupied))
	}

	parentCtx := context.Background()

	valueCtx := context.WithValue(
		parentCtx,
		"SEND_ROOMS",
		encodeContext(messageContext{
			Type:         "SEND_ROOMS",
			ContextValue: conn.RemoteAddr().String(),
		}),
	)

	logger := ucast.logger.GetGroup("ROOMS")

	logger.InfoCtx(valueCtx, "sending rooms payload")

	_, error := conn.Write([]byte(payload.String()))

	if error != nil {
		ucast.logger.NetworkError(error)
		return
	}

	logger.InfoCtx(valueCtx, "finish sending rooms payload")
}
