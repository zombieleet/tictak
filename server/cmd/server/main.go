package main

import (
	tLog "github.com/zombieleet/tictak/server/pkg/logger"
	"golang.org/x/exp/slog"
)

func main() {

	logger := tLog.NewLogger(tLog.LoggerOptions{
		LogLevel: slog.LevelDebug,
	})

}
