package main

import "errors"

type direction int

const (
	RIGHT direction = 1 + iota
	LEFT
	UP
	DOWN
)

type snake struct {
	body      []coord
	direction direction
	length    int
}

func newSnake(d direction, b []coord) *snake {
	return &snake{
		length:    len(b),
		body:      b,
		direction: d,
	}
}

func (s *snake) changeDirection(d direction) {
	opposites := map[direction]direction{
		RIGHT: LEFT,
		LEFT:  RIGHT,
		UP:    DOWN,
		DOWN:  UP,
	}

	if o := opposites[d]; o != 0 && o != s.direction {
		s.direction = d
	}
}

func (s *snake) update(a *stage) error {
	head := s.head()
	nextPosition := coord{x: head.x, y: head.y}

	switch s.direction {
	case RIGHT:
		nextPosition.x++
	case LEFT:
		nextPosition.x--
	case UP:
		nextPosition.y++
	case DOWN:
		nextPosition.y--
	}

	if s.hitItself(nextPosition) {
		return s.die()
	}

	if nextPosition.x > a.width {
		nextPosition.x = 0
	} else if nextPosition.y > a.height {
		nextPosition.y = 0
	} else if nextPosition.x < 0 {
		nextPosition.x = a.width
	} else if nextPosition.y < 0 {
		nextPosition.y = a.height
	}

	if s.length > len(s.body) {
		s.body = append(s.body, nextPosition)
	} else {
		s.body = append(s.body[1:], nextPosition)
	}

	return nil
}

func (s *snake) hitItself(c coord) bool {
	for _, b := range s.body {
		if b.x == c.x && b.y == c.y {
			return true
		}
	}

	return false
}

func (s *snake) head() coord {
	return s.body[len(s.body)-1]
}

func (s *snake) die() error {
	return errors.New("Died")
}
