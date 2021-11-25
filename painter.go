package main

import (
	"fmt"
	"github.com/mattn/go-runewidth"

	"github.com/gdamore/tcell"
)

const (
	defaultColor = tcell.ColorDefault
	bgColor      = tcell.ColorDefault
	snakeColor   = tcell.ColorLightGreen
	fullBlock    = '█'
)

type screenDimensions struct {
	width, height, midY, left, right, top, bottom int
}

func (g *Game) render(screen tcell.Screen) {
	screen.Clear()
	var (
		w, h   = screen.Size()
		midY   = h / 2
		left   = (w - g.stage.width) / 2
		right  = (w + g.stage.width) / 2
		top    = midY - (g.stage.height / 2)
		bottom = midY + (g.stage.height / 2) + 1
	)

	screenDimensions := screenDimensions{w, h, midY, left, right, top, bottom}

	renderArena(screen, g.stage, screenDimensions)
	renderSnake(screen, left, bottom, g.stage.snake)
	renderFood(screen, left, bottom, g.stage.food)
	renderPoints(screen, screenDimensions, g.score)
	renderInstructions(screen, screenDimensions)

	screen.Show()
}

func renderSnake(sc tcell.Screen, left, bottom int, s *snake) {
	for _, b := range s.body {
		sc.SetContent(left+b.x, bottom-b.y, fullBlock, nil, style(snakeColor, snakeColor))
	}
}

func renderFood(screen tcell.Screen, left, bottom int, f *food) {
	screen.SetCell(left+f.coord.x, bottom-f.coord.y, tcell.StyleDefault.Foreground(tcell.ColorTomato), fullBlock)
}

func renderArena(screen tcell.Screen, a *stage, d screenDimensions) {
	for i := d.top; i < d.bottom; i++ {
		screen.SetContent(d.left-1, i, '│', nil, defaultStyle())
		screen.SetContent(d.left+a.width, i, '│', nil, defaultStyle())
	}

	renderContentRecursively(screen, d.left, d.top, a.width, a.height+1, '│', fgStyle(tcell.ColorGrey))

	screen.SetContent(d.left-1, d.top, '┌', nil, defaultStyle())
	screen.SetContent(d.left-1, d.bottom, '└', nil, defaultStyle())
	screen.SetContent(d.left+a.width, d.top, '┐', nil, defaultStyle())
	screen.SetContent(d.left+a.width, d.bottom, '┘', nil, defaultStyle())

	renderContentRecursively(screen, d.left, d.top, a.width, 1, '─', defaultStyle())
	renderContentRecursively(screen, d.left, d.bottom, a.width, 1, '─', defaultStyle())
}

func defaultStyle() tcell.Style {
	return style(defaultColor, bgColor)
}

func fgStyle(fg tcell.Color) tcell.Style {
	style := tcell.StyleDefault.Foreground(fg)
	return style
}

func style(fg, bg tcell.Color) tcell.Style {
	style := tcell.StyleDefault.Foreground(fg).Background(bg)
	return style
}

func renderPoints(screen tcell.Screen, d screenDimensions, s int) {
	pointsStr := fmt.Sprintf("Points: %v", s)
	printMsg(screen, d.right+2, d.top+2, fgStyle(tcell.ColorGreenYellow), pointsStr)
}

func renderInstructions(screen tcell.Screen, d screenDimensions) {
	msgs := []string{"Key bindings:", "ENTER : new game", "ESC : quit"}

	for i, msg := range msgs {
		printMsg(screen, d.right+2, d.midY+i, fgStyle(tcell.ColorDimGrey), msg)
	}
}

func renderContentRecursively(screen tcell.Screen, x, y, w, h int, cell rune, style tcell.Style) {
	for ly := 0; ly < h; ly++ {
		for lx := 0; lx < w; lx++ {
			screen.SetContent(x+lx, y+ly, cell, nil, style)
		}
	}
}

func printMsg(screen tcell.Screen, x, y int, style tcell.Style, str string) {
	for _, c := range str {
		var comb []rune
		w := runewidth.RuneWidth(c)
		if w == 0 {
			comb = []rune{c}
			c = ' '
			w = 1
		}
		screen.SetContent(x, y, c, comb, style)
		x += w
	}

}
