package main

import (
	"fmt"
	"gobasic/model"
)

func main() {

	points := []point.Point{
		{1, 2},
		{5, 3},
		{9, 2},
		{1, 24},
	}

	
	fmt.Println("Initial status: ", points)
	
	for i, point := range points {

		point.Swap()
		points[i] = point

	}

	fmt.Println("Final status: ", points)

}
