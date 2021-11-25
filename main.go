package main

import (
	"time"

	"github.com/gdamore/tcell"
)

var (
	scoreChan    = make(chan int)
	keyEventChan = make(chan keyboardEvent)
)

func main() {
	createGame().Start()
}

func createInitialSnake() *snake {
	return newSnake(RIGHT, []coord{
		{x: 1, y: 10},
		{x: 1, y: 11},
	})
}

func initialScore() int {
	return 0
}

func initialGameStage() *stage {
	return newStage(createInitialSnake(), scoreChan, 20, 60)
}

func (g *Game) gameOver() {
	g.isOver = true
}

func (g *Game) updateInterval() time.Duration {
	ms := 100 - (g.score / 2)
	return time.Duration(ms) * time.Millisecond
}

func (g *Game) addPoints(p int) {
	g.score += p
}

func createGame() *Game {
	return &Game{stage: initialGameStage(), score: initialScore()}
}

func (g *Game) startNewGame() {
	g.stage = initialGameStage()
	g.score = initialScore()
	g.isOver = false
}

func (g *Game) Start() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := recover()
		if err != nil {
			screen.Fini()
			panic(err)
		}
	}()

	initError := screen.Init()
	if initError != nil {
		panic(initError)
	}

	screen.Clear()

	go listenToKeyEvents(screen, keyEventChan)

	g.render(screen)

mainloop:
	for {
		select {
		case p := <-scoreChan:
			g.addPoints(p)
		case e := <-keyEventChan:
			switch e.eventType {
			case MOVE:
				g.stage.snake.changeDirection(e.direction)
			case NEW:
				g.startNewGame()
			case END:
				break mainloop
			}
		default:
			if !g.isOver {
				if err := g.stage.update(); err != nil {
					g.gameOver()
					screen.Fini()
				}
			}

			g.render(screen)

			time.Sleep(g.updateInterval())
		}
	}
}
