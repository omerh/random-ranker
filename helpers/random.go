package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomRanks for setting P_1 to count
func RandomRanks(p1 float64, count int) {
	fmt.Printf("P_1 is:  %v\n", p1)

	for a := 1; a < count; a++ {
		p1 = p1 * getRandomFloat()
		fmt.Printf("P_%v is:  %v\n", a+1, int(p1))
	}
}

func getRandomFloat() float64 {
	rand.Seed(time.Now().UnixNano())
	min := 0.01
	max := 0.05
	v := 1 - (min + rand.Float64()*(max-min))
	return v
}
