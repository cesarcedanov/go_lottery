package main

import (
	"math/rand"
	"time"
)

// Shuffle will shuffle the values in the array
func Shuffle(availableNumbers []string) []string {
	rand.Seed(time.Now().UnixNano())
	for i := len(availableNumbers) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		availableNumbers[i], availableNumbers[j] = availableNumbers[j], availableNumbers[i]
	}
	return availableNumbers
}
