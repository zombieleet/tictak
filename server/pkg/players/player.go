package players

import (
	"net"
	"sync"
)

type PlayerInfo struct {
	Address      string `json:"addres"`
	Name         string `json:"name"`
	CurrentScore uint8  `json:"player_score"`
	Connection   net.Conn
}

type Players struct {
	PlayerOne *PlayerInfo `json:"player_one"`
	PlayerTwo *PlayerInfo `json:"player_two"`
}

type PlayerConnectedConnection struct {
	Address    string
	Connection net.Conn
}

// PlayersConnected
// a struct with private fields
// representing connected players
type PlayersConnected struct {
	// list of connected players (network address)
	Players []PlayerConnectedConnection
	mutex   sync.RWMutex
}

func CreateConnectedPlayers() *PlayersConnected {
	return &PlayersConnected{
		Players: make([]PlayerConnectedConnection, 1),
	}
}

// AddPlayer
// Adds a player to the list of players that is connected to the server
func (pNoRoom *PlayersConnected) AddPlayer(conn net.Conn, address string) {
	pNoRoom.mutex.Lock()
	defer pNoRoom.mutex.Unlock()
	pNoRoom.Players = append(pNoRoom.Players, PlayerConnectedConnection{
		Address:    address,
		Connection: conn,
	})
}

// RemovePlayer
// Removes a player from the list of players connected to the server
// TODO: when the client uses `q` | EOL to disconnect from the servr, this method will be called
func (pNoRoom *PlayersConnected) RemovePlayer() {}
