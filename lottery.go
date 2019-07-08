package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Lottery creates a new type of 'Lottery'
// which take some number from a range of number
type Lottery struct {
	AvailableNumbers []string
	LotteryOpts      LotteryOptions
}

// NewLottery will create a new Lottery struct
func NewLottery(opts LotteryOptions) (Lottery, error) {
	if opts.MaxNumber < 1 {
		return Lottery{}, errors.New("Missing Max Number in the lottery options")
	}
	if opts.NumberPerCombination < 1 {
		return Lottery{}, errors.New("Missing Number per combination in the lottery options")
	}
	numbers := []string{}
	for i := opts.MinNumber; i <= opts.MaxNumber; i++ {
		numbers = append(numbers, fmt.Sprint(i))
	}
	return Lottery{
		AvailableNumbers: numbers,
		LotteryOpts:      opts,
	}, nil
}

// ShowAllAvailableNumbers shows up all the Availablesadas Numbers in the lottery
func (l Lottery) ShowAllAvailableNumbers() {
	fmt.Println(strings.Join(l.AvailableNumbers, " "))
}

// Shuffle will shuffle the values in the array
func (l Lottery) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := len(l.AvailableNumbers) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		l.AvailableNumbers[i], l.AvailableNumbers[j] = l.AvailableNumbers[j], l.AvailableNumbers[i]
	}
}
