package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	// Create a empty LotteryOptions
	var opts LotteryOptions
	// Using Flags get the value from the Arguments
	flag.IntVar(&opts.MinNumber, "min-number", 1, "Minimun Number of the Lottery")
	flag.IntVar(&opts.MaxNumber, "max-number", 38, "Maximun Number of the Lottery")
	flag.IntVar(&opts.NumberPerCombination, "number-per-combination", 6, "Number Per Combination")
	flag.IntVar(&opts.CombinationsLength, "combinations-length", 4, "How many combination per Ticket")
	flag.Parse()

	lottery, err := NewLottery(opts)

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
