package main

type coord struct {
	x, y int
}

func (c *coord) isOnAnother(another coord) bool {
	return c.x == another.x && c.y == another.y
}
