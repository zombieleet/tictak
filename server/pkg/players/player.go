package players

import (
	"sync"
)

type PlayerInfo struct {
	Address      string
	Name         string
	CurrentScore uint8
}

type Players struct {
	PlayerOne *PlayerInfo
	PlayerTwo *PlayerInfo
}

// PlayersConnected
// a struct with private fields
// representing connected players
type PlayersConnected struct {
	// list of connected players (network address)
	players []string
	mutex   sync.RWMutex
}

func CreateConnectedPlayers() *PlayersConnected {
	return &PlayersConnected{
		players: make([]string, 1),
	}
}

// AddPlayer
// Adds a player to the list of players is connected to the server
func (pNoRoom *PlayersConnected) AddPlayer(address string) {
	pNoRoom.mutex.Lock()
	defer pNoRoom.mutex.Unlock()
	pNoRoom.players = append(pNoRoom.players, address)
}

// RemovePlayer
// Removes a player from the list of players connected to the server
// TODO: when the client uses `q` | EOL to disconnect from the servr, this method will be called
func (pNoRoom *PlayersConnected) RemovePlayer() {}
