package main

import (
	"fmt"
)

type LotteryOptions struct {
	MinNumber            int
	MaxNumber            int
	NumberPerCombination int
	CombinationsLength   int
}

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
	combinationArr := make([]Combination, 0, lottery.LotteryOpts.CombinationsLength)
	for i := 0; i < lottery.LotteryOpts.CombinationsLength; i++ {
		lottery.Shuffle()
		// lottery.ShowAllAvailableNumbers()

		combination, err := NewCombination(lottery.AvailableNumbers, lottery.LotteryOpts.NumberPerCombination)
		if err != nil {
			fmt.Printf("%s", err.Error())
			return
		}
		// combination.Sort()
		fmt.Println(combination.Numbers)
		combinationArr = append(combinationArr, combination)
	}
	for _, i := range combinationArr {
		fmt.Println(i)
	}

	// ticket, err := NewTicket(combinationArr)
	// if err != nil {
	// 	fmt.Printf("%s", err.Error())
	// 	return
	// }
	// if err := ticket.WriteLines(ticket.CreatedTime.Format(time.RFC850)); err != nil {
	// 	return
	// }
}
