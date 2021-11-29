package main

import (
	"math/rand"
	"time"
)

type Game struct {
	stage  *stage
	score  int
	isOver bool
}

type stage struct {
	food       *food
	snake      *snake
	canEatFood func(*stage, coord) bool
	height     int
	width      int
	pointsChan chan int
}

func newStage(s *snake, p chan int, h, w int) *stage {
	a := &stage{
		snake:      s,
		height:     h,
		width:      w,
		pointsChan: p,
		canEatFood: canEatFood,
	}

	a.generateRandomFood()

	return a
}

func (a *stage) update() error {
	if err := a.snake.update(a); err != nil {
		return err
	}

	if a.canEatFood(a, a.snake.head()) {
		go a.increasePoints(a.food.points)
		a.snake.length++
		a.generateRandomFood()
	}

	return nil
}

func (a *stage) increasePoints(p int) {
	a.pointsChan <- p
}

func (a *stage) generateRandomFood() {
	var x, y int
	foodPosition := randomFoodPosition(a, x, y)
	a.food = newFood(foodPosition)
}

func randomFoodPosition(a *stage, x int, y int) coord {
	rand.Seed(time.Now().UnixNano())
	var foodPosition coord
	for {
		x = rand.Intn(a.width)
		y = rand.Intn(a.height)

		foodPosition = coord{x: x, y: y}
		if !a.isOccupied(foodPosition) {
			break
		}
	}
	return foodPosition
}

func canEatFood(a *stage, c coord) bool {
	return c.samePositionAs(a.food.coord)
}

func (a *stage) isOccupied(c coord) bool {
	return a.snake.hitItself(c)
}
