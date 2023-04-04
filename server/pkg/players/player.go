package players

import (
	"net"
	"sync"
)

// PlayerInfo holds info about a connected player in a room
type PlayerInfo struct {
	Address      string `json:"addres"`
	Name         string `json:"name"`
	CurrentScore uint8  `json:"player_score"`
	Connection   net.Conn
}

// PlayersInRoom maps playerone and playertwo to list of players in a room
type PlayersInRoom struct {
	PlayerOne *PlayerInfo `json:"player_one"`
	PlayerTwo *PlayerInfo `json:"player_two"`
}

// PlayerConnectedConnection holds connection information for players in the game
type PlayerConnectedConnection struct {
	Address    string
	Connection net.Conn
}

// PlayersConnected holds list of players that are in the game (tcp client connections)
type PlayerConnectedToGame struct {
	// list of connected players (network address)
	Players []PlayerConnectedConnection
	mutex   sync.RWMutex
}

func CreateConnectedPlayers() *PlayerConnectedToGame {
	return &PlayerConnectedToGame{
		Players: make([]PlayerConnectedConnection, 0),
	}
}

// AddPlayer
// Adds a player to the list of players that is connected to the server
func (pNoRoom *PlayerConnectedToGame) AddPlayer(conn net.Conn, address string) {
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
func (pNoRoom *PlayerConnectedToGame) RemovePlayer() {}
