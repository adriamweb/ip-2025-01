package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		S        float64
		contador float64 = 1
	)

	for i := 1; i <= 51; i++ {

		if i%2 == 0 {
			S -= 1.0 / math.Pow(contador, 3)
		} else {
			S += 1.0 / math.Pow(contador, 3)
			contador += 2
		}
	}

	fmt.Print("\n∏: ", math.Pow((S*32), (1.0/3.0)))

}
