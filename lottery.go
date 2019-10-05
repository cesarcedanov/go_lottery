package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// LotteryOptions will define some rules for the Lottery
type LotteryOptions struct {
	MinNumber            int
	MaxNumber            int
	NumberPerCombination int
	CombinationsLength   int
}

// Lottery creates a new type of 'Lottery'
// which take some number from a range of number
type Lottery struct {
	AvailableNumbers []string
	LotteryOpts      LotteryOptions
}

// NewLottery will create a new Lottery struct
func NewLottery(opts LotteryOptions) (*Lottery, error) {
	if opts.MaxNumber < 1 {
		return nil, errors.New("Missing Max Number in the lottery options")
	}
	if opts.NumberPerCombination < 1 {
		return nil, errors.New("Missing Number per combination in the lottery options")
	}

	numbers := []string{}
	for i := opts.MinNumber; i <= opts.MaxNumber; i++ {
		numbers = append(numbers, fmt.Sprintf("%02d", i))
	}

	fmt.Println("Lottery Number:", numbers)
	return &Lottery{
		AvailableNumbers: numbers,
		LotteryOpts:      opts,
	}, nil
}

// Shuffle will shuffle the values in the array
func (l *Lottery) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range l.AvailableNumbers {
		newIndex := rand.Intn(len(l.AvailableNumbers) - 1)
		l.AvailableNumbers[i], l.AvailableNumbers[newIndex] = l.AvailableNumbers[newIndex], l.AvailableNumbers[i]
	}
}
