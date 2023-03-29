package main

import (
	"fmt"
	"gobasic/model"
)

func Swap(x, y *int) {
	if (*x >= *y) {
		aux := *x
		*x = *y
		*y = aux
	}
}

func main() {

	points := []point.Point{
		{1, 2},
		{5, 3},
		{9, 2},
		{1, 24},
	}

	
	fmt.Println("Initial status: ", points)
	
	for i, point := range points {

		Swap(&point.X, &point.Y)
		points[i] = point

	}

	fmt.Println("Final status: ", points)

}
