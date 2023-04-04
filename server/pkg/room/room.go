// room package contains methods and types needed for hanlding room related operations
package room

import (
	"context"
	"errors"
	"fmt"
	"github.com/zombieleet/tictak/server/pkg/message"
	"github.com/zombieleet/tictak/server/pkg/players"
	"net"
	"strconv"
	"strings"
	"sync"
)

// RoomInfo
// Holds information about a room
type RoomInfo struct {
	// boolean to indicate a room that is not yet filled
	Occupied              bool  `json:"occupied"`
	NumberOfPlayersInRoom uint8 `json:"num_players_in_room"`
	// can be public or private
	// by default it is set to public
	Visibility string `json:"visibility"`
	Name       string `json:"room_name"`

	// raw grid of the game
	grid [][]string

	GridSize uint8 `json:"grid_size"`

	// players added to a room (max is 2)
	Players *players.PlayersInRoom `json:"players_in_room"`
}

type RoomMap map[uint8]*RoomInfo

// Room
// Holds a relationship between a rooms id and the rooms info
type Room struct {
	Rooms   *RoomMap
	mutex   *sync.Mutex
	message *message.Message
}

// createGrid creates an x by x grid
func createGrid(size uint8) [][]string {
	grid := make([][]string, size)
	for row := uint8(0); row < size; row += 1 {
		grid[row] = make([]string, size)
		for col := uint8(0); col < size; col += 1 {
			grid[row][col] = "-"
		}
	}
	return grid
}

// CreateRooms
// create rooms with a max roomCount of 10
func CreateRooms(roomCount uint8, message *message.Message) *Room {
	rooms := make(RoomMap)
	room := &Room{
		Rooms:   &rooms,
		mutex:   new(sync.Mutex),
		message: message,
	}

	if roomCount > 10 || roomCount < 1 {
		roomCount = 10
	}

	for roomIndex := uint8(1); roomIndex <= roomCount; roomIndex += 1 {
		(*room.Rooms)[roomIndex] = &RoomInfo{
			Occupied:              false,
			NumberOfPlayersInRoom: 0,
			grid:                  createGrid(3),
			GridSize:              3,
			Visibility:            "PUBLIC",
			Players:               new(players.PlayersInRoom),
			Name:                  fmt.Sprintf("%s %d", "Room", roomIndex),
		}
	}

	return room
}

// EnterRoom registers a user in a room
func (r *Room) EnterRoom(
	_ net.Conn,
	commsChan chan string,
	cancelCtxFunc context.CancelCauseFunc,
	args ...interface{},
) {

	rawRoomId := args[0].(string)
	address := args[1].(string)

	conv, err := strconv.Atoi(rawRoomId)

	if err != nil {
		cancelCtxFunc(errors.Join(E_ROOM_ID_CONVERSION_ERROR, err, fmt.Errorf("roomid(%s)", rawRoomId)))
		return
	}

	roomId := uint8(conv)

	r.mutex.Lock()
	defer r.mutex.Unlock()

	room := (*r.Rooms)[roomId]

	if room.NumberOfPlayersInRoom == 1 {
		room.Players.PlayerTwo = &players.PlayerInfo{
			Address:      address,
			Name:         "",
			CurrentScore: 0,
		}
		room.Occupied = true
		room.NumberOfPlayersInRoom = 2

		go r.message.Unicast.SendPlayerTwoConnectedInfoFromPlayerOne()

	} else {
		room.Players = &players.PlayersInRoom{
			PlayerOne: &players.PlayerInfo{
				Address:      address,
				Name:         "",
				CurrentScore: 0,
			},
		}
		room.NumberOfPlayersInRoom = 1
	}

	//commsChan <- *room
	go r.message.Broadcast.SendPlayerCountInRoomToConnectedPlayers(roomId, room.NumberOfPlayersInRoom)
	//r.message.Broadcast.SendPlayerCountToConnectedPlayers(r.message.Broadcast, room.NumberOfPlayersInRoom)

	// broadcast an update to all players connected in the list of room
	// check the playerconnected struct
}

func (r *Room) SendRoomsInfo(conn net.Conn, cancelCtxFunc context.CancelCauseFunc) {
	var payload strings.Builder

	for roomId, roomInfo := range *r.Rooms {
		payload.WriteString(fmt.Sprintf("%d. %s (Occupied=%t)_", roomId, roomInfo.Name, roomInfo.Occupied))
	}
	go r.message.Unicast.SendRoomsInfoToNewClient(conn, payload.String())
}
