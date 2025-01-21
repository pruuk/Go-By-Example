// trying to produce a normal distribution using math and rand packages
package main

import (
	"fmt"
	"math"
	"math/rand"
)

// NormalDistribution returns a random number from a normal distribution
func NormalDistribution(mean, stdDev float64) float64 {
	return math.Round(mean + stdDev*rand.NormFloat64())
}

func main() {
	mean := 100.0
	stdDev := 10.0
	result := NormalDistribution(mean, stdDev)
	fmt.Println(result)
}
