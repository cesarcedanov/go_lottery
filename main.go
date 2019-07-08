package main

import (
	"fmt"
	"time"
)

func main() {

	lottery, err := NewLottery(
		LotteryOptions{
			MinNumber:            1,
			MaxNumber:            38,
			NumberPerCombination: 6,
			CombinationsLength:   4,
		},
	)

	if err != nil {
		fmt.Println("Unable to create the Lottery")
		return
	}

	var combinationArr []*Combination
	for i := 0; i < lottery.LotteryOpts.CombinationsLength; i++ {

		combination, err := NewCombination(
			lottery.AvailableNumbers,
			lottery.LotteryOpts.NumberPerCombination)

		if err != nil {
			fmt.Printf("%s", err.Error())
			return
		}
		combination.Sort()

		combinationArr = append(combinationArr, combination)
	}

	ticket, err := NewTicket(combinationArr)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return
	}
	if err := ticket.WriteLines(ticket.CreatedTime.Format(time.RFC850)); err != nil {
		return
	}
}
