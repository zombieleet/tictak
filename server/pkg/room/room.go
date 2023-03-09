package room

import (
	"fmt"
	"github.com/zombieleet/tictak/server/pkg/players"
)

type RoomInfo struct {
	// boolean to indicate a room that is not yet filled
	Occupied              bool
	NumberOfPlayersInRoom uint8
	// can be public or private
	// by default it is set to public
	Visibility string
	Name       string
	players    *players.Players
}

type Room map[uint8]*RoomInfo

func CreateRooms(roomCount uint8) *Room {
	rooms := make(Room)

	if roomCount > 10 || roomCount < 1 {
		roomCount = 10
	}

	for room := uint8(1); room <= roomCount; room += 1 {
		rooms[room] = &RoomInfo{
			Occupied:              false,
			NumberOfPlayersInRoom: 0,
			Visibility:            "PUBLIC",
			players:               new(players.Players),
			Name:                  fmt.Sprintf("%s %d", "Room", room),
		}
	}

	return &rooms
}
