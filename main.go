package main

import (
	"fmt"
	"log"
)

func main() {

	lottery, err := NewLottery(
		LotteryOptions{
			MinNumber:            1,
			MaxNumber:            38,
			NumberPerCombination: 7,
			CombinationsLength:   4,
		},
	)

	if err != nil {
		fmt.Println("Unable to create the Lottery")
		return
	}
	lottery.Shuffle()

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

	if err := ticket.WriteLines(ticket.CreatedTime.Format("Monday, 02-Jan-06 15.04.05")); err != nil {
		log.Fatal(err)
		return
	}
}
