package main

import (
	"fmt"
)

var combinationArr []Combination

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


	for i := 0; i < lottery.LotteryOpts.CombinationsLength; i++ {

		combination, err := NewCombination(
			lottery.AvailableNumbers,
			lottery.LotteryOpts.NumberPerCombination)

		if err != nil {
			fmt.Printf("%s", err.Error())
			return
		}
		fmt.Println(combinationArr.Numbers)

		combinationArr = append(combinationArr, combination)
	}
	fmt.Println(combinationArr)
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
