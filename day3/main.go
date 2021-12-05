package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const input = "day3/input.txt"

func main() {
	s, err := inputScanner(input)
	if err != nil {
		log.Fatalf("unable to get input scanner: %v", err)
	}

	// each index will have a counter that determines whether 1 or 0 was the
	// most common bit. for every 1 encountered we add 1, every 0 we subtract
	// 1. if the sum > 0, 1 was the most common, 0 then it was a tie(shouldn't,
	// happen), sum < 0, 0 was the most common
	counter := make([]int, 12) // assuming length from given input
	for s.Scan() {
		l := s.Text()
		if err := updateCounter(counter, strings.TrimSpace(l)); err != nil {
			log.Fatalf("unable to update counter: %v", err)
		}
	}

	// form gamma
	var (
		gamma uint
		power uint = 1
	)
	for i := len(counter) - 1; i >= 0; i-- {
		if counter[i] == 0 {
			log.Fatal("unexpected tie of binary digits")
		}

		if counter[i] > 0 {
			gamma += power
		}

		power *= 2
	}

	fmt.Printf("gamma: %b\n", gamma)
	// form epsilon by flipping all bits of gamma, including leading zeros
	epsilon := gamma
	for i := 0; i < 12; i++ {
		epsilon = epsilon ^ (1 << i)
	}

	fmt.Printf("epsilon: %b\n", epsilon)
	fmt.Printf("answer: %d", epsilon*gamma)
}

func updateCounter(counter []int, line string) error {
	for i := range line {
		switch line[i] {
		case '0':
			counter[i]--
		case '1':
			counter[i]++
		default:
			return fmt.Errorf("unexpected character: %s", string(line[i]))
		}
	}

	return nil
}

func inputScanner(input string) (*bufio.Scanner, error) {
	path, err := filepath.Abs(input)
	if err != nil {
		return nil, fmt.Errorf("unable to get absolute path: %w", err)
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	return bufio.NewScanner(f), nil
}
