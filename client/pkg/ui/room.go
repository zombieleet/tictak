package ui

import (
	"fmt"
	"github.com/rivo/tview"
	"github.com/zombieleet/tictak/client/pkg/commands"
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
func (room *Room) CreateRoomListUI(commsChan chan string, data string) {
	rooms := strings.Split(data, "\n")
	listUI := room.ui.createList()
	for key, roomValue := range rooms {
		roomId := rune(key + 1)
		listUI.AddItem(roomValue, "", roomId, func(r string, id rune) func() {
			return func() {
				commsChan <- fmt.Sprintf("%s%d_", commands.Commands["CMD_ENTER_ROOM"], roomId)
			}
		}(roomValue, roomId))
	}

	roomsGrid := room.createRoomGrid()

	roomsGrid.AddItem(listUI, 0, 0, 2, 1, 0, 0, true)
	room.ui.ContentLayout.AddItem(roomsGrid, 0, 0, 2, 1, 0, 0, true)
}

func (room *Room) UpdateRoomListUI(_ chan string, d string) {
	fmt.Println(d)
}

func (room *Room) createRoomGrid() *tview.Grid {
	return room.ui.createGrid([]int{1}).SetColumns(0)
}
