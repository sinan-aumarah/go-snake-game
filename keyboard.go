package main

import "github.com/gdamore/tcell"

type keyboardEventType int

const (
	MOVE keyboardEventType = 0 + iota
	NEW
	END
)

type keyboardEvent struct {
	eventType keyboardEventType
	key       tcell.Key
	direction direction
}

func listenToKeyEvents(screen tcell.Screen, evChan chan keyboardEvent) {

	for {
		switch ev := screen.PollEvent(); ev.(type) {
		case *tcell.EventKey:
			keyEvent := ev.(*tcell.EventKey)
			switch keyEvent.Key() {
			case tcell.KeyLeft:
				evChan <- keyboardEvent{eventType: MOVE, key: keyEvent.Key(), direction: LEFT}
			case tcell.KeyDown:
				evChan <- keyboardEvent{eventType: MOVE, key: keyEvent.Key(), direction: DOWN}
			case tcell.KeyRight:
				evChan <- keyboardEvent{eventType: MOVE, key: keyEvent.Key(), direction: RIGHT}
			case tcell.KeyUp:
				evChan <- keyboardEvent{eventType: MOVE, key: keyEvent.Key(), direction: UP}
			case tcell.KeyEsc:
				evChan <- keyboardEvent{eventType: END, key: keyEvent.Key()}
			case tcell.KeyEnter:
				evChan <- keyboardEvent{eventType: NEW, key: keyEvent.Key()}
			}
		}
	}

}
