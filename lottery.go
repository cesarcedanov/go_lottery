package main

import "fmt"

// Lottery creates a new type of 'Lottery'
// which take some number from a range of number
type Lottery struct {
	Numbers []int
}

// NewLottery will create a new Lottery struct
func NewLottery() Lottery {
	numbers := []int{}
	for i := 1; i <= 38; i++ {
		numbers = append(numbers, i)
	}
	return Lottery{
		Numbers: numbers,
	}
}

// ShowNumber shows up all the number in the lottery
func (l Lottery) ShowNumber() {
	for _, number := range l.Numbers {
		fmt.Println(number)
	}
}
