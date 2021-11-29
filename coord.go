package main

type coord struct {
	x, y int
}

func (c *coord) samePositionAs(another coord) bool {
	return c.x == another.x && c.y == another.y
}
