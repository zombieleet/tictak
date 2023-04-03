package message

import (
	"github.com/zombieleet/tictak/server/pkg/logger"
	"github.com/zombieleet/tictak/server/pkg/players"
)

type broadcast struct {
	logger                       *logger.Logger
	connectedPlayersForBroadcast []players.PlayerConnectedConnection
}

func (bcast *broadcast) SendPlayerCountToConnectedPlayers() {
	fmt.Println(bcast.connectedPlayersForBroadcast)
}
