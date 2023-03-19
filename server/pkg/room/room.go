// room package contains methods and types needed for hanlding room related operations
package room

import (
	"fmt"
	"github.com/zombieleet/tictak/server/pkg/players"
	"strconv"
	"sync"
)

// RoomInfo
// Holds information about a room
type RoomInfo struct {
	// boolean to indicate a room that is not yet filled
	Occupied              bool
	NumberOfPlayersInRoom uint8
	// can be public or private
	// by default it is set to public
	Visibility string
	Name       string

	// raw grid of the game
	Grid [][]string

	// players added to a room (max is 2)
	players *players.Players
}

type RoomMap map[uint8]*RoomInfo

// Room
// Holds a relationship between a rooms id and the rooms info
type Room struct {
	Rooms *RoomMap
	mutex *sync.Mutex
}

// CreateRooms
// create rooms with a max roomCount of 10
func CreateRooms(roomCount uint8) *Room {
	rooms := make(RoomMap)
	room := &Room{
		Rooms: &rooms,
		mutex: new(sync.Mutex),
	}

	if roomCount > 10 || roomCount < 1 {
		roomCount = 10
	}

	for roomIndex := uint8(1); roomIndex <= roomCount; roomIndex += 1 {

		(*room.Rooms)[roomIndex] = &RoomInfo{
			Occupied:              false,
			NumberOfPlayersInRoom: 0,
			Visibility:            "PUBLIC",
			players:               new(players.Players),
			Name:                  fmt.Sprintf("%s %d", "Room", roomIndex),
		}
	}

	return room
}

// EnterRoom registers a user in a room
func (r *Room) EnterRoom(args ...interface{}) {

	rawRoomId := args[0].(string)
	address := args[1].(string)

	conv, err := strconv.Atoi(rawRoomId)

	// TODO: send error report back to client
	// revert to room list ui
	if err != nil {
		return
	}

	roomId := uint8(conv)

	r.mutex.Lock()
	defer r.mutex.Unlock()

	room := (*r.Rooms)[roomId]

	if room.NumberOfPlayersInRoom == 1 {
		room.players.PlayerTwo = &players.PlayerInfo{
			Address:      address,
			Name:         "",
			CurrentScore: 0,
		}
		room.Occupied = true
		room.NumberOfPlayersInRoom = 2
		return
	}

	room.players = &players.Players{
		PlayerOne: &players.PlayerInfo{
			Address:      address,
			Name:         "",
			CurrentScore: 0,
		},
	}
	fmt.Printf("%+v", room.players)
	// broadcast an update to all players connected in the list of room
	// check the playerconnected struct
}
