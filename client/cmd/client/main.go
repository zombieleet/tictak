package main

import (
	"github.com/zombieleet/tictak/client/pkg/connection"
)

func main() {
	gameClient, err := connection.CreateGameClient(
		connection.GameClientOptions{
			HostName: "0.0.0.0",
			Port:     1234,
		},
	)

	if err != nil {
		panic(err)
	}

	gameClient.Connect()
}
