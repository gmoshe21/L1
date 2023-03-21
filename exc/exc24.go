package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) GetX() float64{
	return p.x
}

func (p *Point) GetY() float64{
	return p.y
}

func new(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func main() {
	p1 := new(7, 0)
	p2 := new(8, 0)

	fmt.Println("Distance beetween", p1.GetX(), ":", p1.GetY(),
				"and", p2.GetX(), ":", p2.GetY(), "=",
				math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2)))

}