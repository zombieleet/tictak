package message

import (
	"context"
	"github.com/zombieleet/tictak/server/pkg/command"
	"github.com/zombieleet/tictak/server/pkg/logger"
	"github.com/zombieleet/tictak/server/pkg/players"
	"strconv"
)

type broadcast struct {
	logger                       *logger.Logger
	connectedPlayersForBroadcast *players.PlayerConnectedToGame
}

func (bcast *broadcast) SendPlayerCountInRoomToConnectedPlayers(roomId uint8, numPlayersInRoom uint8) {
	parentCtx := context.Background()
	valueCtx := context.WithValue(
		parentCtx,
		"SEND_ROOMS_LIST_UPDATE",
		encodeContext(messageContext{
			Type: "SEND_ROOMS_LIST_UPDATE",
		}),
	)

	bcast.logger.LogWithCtx(
		valueCtx,
		"sending rooms list update",
		"ROOMS",
		[]any{},
	)

	for _, conn := range bcast.connectedPlayersForBroadcast.Players {
		go conn.Connection.Write(
			[]byte(command.Commands["CMD_SEND_ROOM_LIST_UPDATE_TO_CLIENTS"] +
				"room_id=" +
				strconv.FormatUint(uint64(roomId), 10) +
				strconv.FormatUint(uint64(numPlayersInRoom), 10) +
				"\n",
			),
		)
	}

}
