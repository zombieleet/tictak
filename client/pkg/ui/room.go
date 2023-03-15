package ui

import (
	"context"
	_ "fmt"
	"strconv"
	"strings"
)

type Room struct {
	ui *UI
}

func InitRoomsUI(ui *UI) *Room {
	return &Room{ui}
}

func (room *Room) CreateRoomListUI(ctx context.Context, data string) {
	rooms := strings.Split(data, "\n")
	listUI := room.ui.createList()
	for key, room := range rooms {
		listUI.AddItem(room, "", []rune(strconv.Itoa(key))[0], nil)
	}
	room.ui.Layout.AddItem(listUI, 1, 0, 1, 1, 0, 0, true)
}

func (room *Room) UpdateRoomListUI(_ context.Context, _ string) {}
