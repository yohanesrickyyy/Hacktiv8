package main

import (
	"fmt"
	"math"
	"sync"
)

type Circle struct {
	Area          float64
	Circumference float64
}

func main() {
	diameters := []float64{5.0, 10.0, 15.0, 20.0} 

	resultChannel := make(chan Circle)

	var wg sync.WaitGroup

	for _, diameter := range diameters {
		wg.Add(1)
		go func(d float64) {
			defer wg.Done()
			resultChannel <- calculateCircle(d)
		}(diameter)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	for circle := range resultChannel {
		fmt.Printf("Diameter: %.2f, Luas: %.2f, Keliling: %.2f\n", diam(circle.Area), circle.Area, circle.Circumference)
	}
}

func calculateCircle(diameter float64) Circle {
	radius := diameter / 2
	area := math.Pi * math.Pow(radius, 2)
	circumference := 2 * math.Pi * radius
	return Circle{Area: area, Circumference: circumference}
}

func diam(n float64) float64 {
	return n / 2
}
