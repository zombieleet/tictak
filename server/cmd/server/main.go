package main

import (
	"github.com/zombieleet/tictak/server/pkg/connection"
	tLog "github.com/zombieleet/tictak/server/pkg/logger"
	"golang.org/x/exp/slog"
)

func main() {

	logger := tLog.NewLogger(tLog.LoggerOptions{
		LogLevel: slog.LevelDebug,
	})

	gameServer := connection.CreateGameServer(
		connection.GameServerOptions{
			HostName: "0.0.0.0",
			Port:     1234,
			Logger:   logger,
		},
	)

	if gameServer == nil {
		panic("Cannot start game server, connection not established.")
	}

	gameServer.Start()

}
