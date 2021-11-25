package main

const foodPoints = 1

type food struct {
	points int
	coord  coord
}

func newFood(coord coord) *food {
	return &food{
		points: foodPoints,
		coord:  coord,
	}
}
