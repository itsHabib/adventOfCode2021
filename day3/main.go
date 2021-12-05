package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const input = "day3/input.txt"

func main() {
	if err := part1(); err != nil {
		log.Fatalf("unable to complete part 1: %v", err)
	}

	if err := part2(); err != nil {
		log.Fatalf("unable to complete part 2: %v", err)
	}
}

// create a counter array that holds a sum in which each index value describes
// the most common bit of the ith bit of the number. We form the gamma binary
// num and flip its bits to get epsilon.
func part1() error {
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
		n, err := strconv.ParseUint(l, 2, 64)
		if err != nil {
			return fmt.Errorf("unable to parse uint: %w", err)
		}
		updateCounter(counter, uint(n))
	}

	if err := s.Err(); err != nil {
		return fmt.Errorf("erorr while scanning: %w", err)
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
	fmt.Printf("answer: %d\n", epsilon*gamma)

	return nil
}

// We create a counter array at the specified index for both the oxygen and
// co2 reading for each bit. With the counter array we can form both
// the oxygen and co2 reading using the most/least common bits.
func part2() error {
	s, err := inputScanner(input)
	if err != nil {
		log.Fatalf("unable to get input scanner: %v", err)
	}

	counter := make([]int, 12) // assuming length from given input
	var nums []uint
	for s.Scan() {
		l := strings.TrimSpace(s.Text())

		n, err := strconv.ParseUint(l, 2, 64)
		if err != nil {
			return fmt.Errorf("unable to parse uint: %w", err)
		}

		nums = append(nums, uint(n))
	}

	onums := make([]uint, len(nums))
	copy(onums, nums)
	oxygen := getReading(counter, onums, func(sum int) uint {
		var criteria uint
		if sum >= 0 {
			criteria = 1
		}

		return criteria
	})

	co2Nums := make([]uint, len(nums))
	copy(co2Nums, nums)
	co2 := getReading(counter, co2Nums, func(sum int) uint {
		var criteria uint
		if sum < 0 {
			criteria = 1
		}

		return criteria
	})

	fmt.Printf("oxygen: %d\n", oxygen)
	fmt.Printf("co2: %d\n", co2)
	fmt.Printf("answer: %d\n", oxygen*co2)

	return nil
}

func getReading(counter []int, nums []uint, getCriteria func(int) uint) uint {
	shift := len(counter) - 1
	for i := 0; i < len(counter) && len(nums) > 1; i++ {
		// reset after each iteration
		resetCounter(counter, nums, i, shift)
		criteria := getCriteria(counter[i])

		// go through each num and filter out ones that do not have the criteria
		// bit set at i
		for j := 0; j < len(nums) && len(nums) > 1; j++ {
			// if the ith bit isn't set to criteria, we need to remove it
			if (nums[j]>>shift)&1 != criteria {
				nums = append(nums[:j], nums[j+1:]...)
				j--
			}
		}
		shift--
	}

	// assuming valid input
	return nums[0]
}

func resetCounter(counter []int, nums []uint, i int, shift int) {
	counter[i] = 0
	for j := range nums {
		switch (nums[j] >> shift) & 1 {
		case 0:
			counter[i]--
		case 1:
			counter[i]++
		}
	}
}

func updateCounter(counter []int, n uint) {
	shift := len(counter) - 1
	for i := range counter {
		switch (n >> shift) & 1 {
		case 0:
			counter[i]--
		case 1:
			counter[i]++
		}
		shift--
	}
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
