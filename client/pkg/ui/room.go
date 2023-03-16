package ui

import (
	"context"
	"fmt"
	"github.com/rivo/tview"
	"github.com/zombieleet/tictak/client/pkg/commands"
	"strconv"
	"strings"
)

type Room struct {
	ui *UI
}

func InitRoomsUI(ui *UI) *Room {
	return &Room{ui}
}

// CreateRoomListUI creates the ui for room list
// the room list ui is appended to ContentLayout on row 0 and column 0
func (room *Room) CreateRoomListUI(commsChan chan string, ctx context.Context, data string) {
	rooms := strings.Split(data, "\n")
	listUI := room.ui.createList()
	for key, roomValue := range rooms {
		roomId := []rune(strconv.Itoa(key + 1))[0]
		listUI.AddItem(roomValue, "", roomId, func(r string) func() {
			return func() {
				commsChan <- fmt.Sprintf("%s %d", commands.Commands["CMD_ENTER_ROOM"], roomId)
			}
		}(roomValue))
	}

	roomsGrid := room.createRoomGrid()

	roomsGrid.AddItem(listUI, 0, 0, 2, 1, 0, 0, true)
	room.ui.ContentLayout.AddItem(roomsGrid, 0, 0, 2, 1, 0, 0, true)
}

func (room *Room) UpdateRoomListUI(_ chan string, _ context.Context, _ string) {}

func (room *Room) createRoomGrid() *tview.Grid {
	return room.ui.createGrid([]int{1}).SetColumns(0)
}
