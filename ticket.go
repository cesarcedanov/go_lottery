package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

// Ticket should contains one to multiple combinations
type Ticket struct {
	Combinations []*Combination
	CreatedTime  time.Time
}

// NewTicket will create a Ticket with some combination of some number
// and a Time when was created.
func NewTicket(combinations []*Combination) (*Ticket, error) {
	if len(combinations) < 1 {
		return nil, errors.New("Can NOT create a Ticket without Combiniations")
	}
	return &Ticket{
		Combinations: combinations,
		CreatedTime:  time.Now(),
	}, nil
}

// AppendCombination will add a new combination to the array of combination in the ticket
func (t *Ticket) AppendCombination(combination *Combination) {
	t.Combinations = append(t.Combinations, combination)
}

// WriteLines will write the array of string into a file
func (t *Ticket) WriteLines(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)

	fmt.Fprintln(w, "Created At: ", time.Now().Format(time.RFC850))

	for _, combination := range t.Combinations {
		fmt.Fprintln(w, combination.ToString())
	}
	return w.Flush()
}

func (t *Ticket) SaveTicket(filename string) error {
	combination := "1,4,12,18,26,37"
	err := ioutil.WriteFile(filename, []byte(combination), 0)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func newTicketFromFile(filename string) *Combination {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return &Combination{
		Numbers: s,
	}
}
