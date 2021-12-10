package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

const input = "day7/input.txt"

func main() {
	data, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatalf("unable to read input file: %v", err)
	}

	nums, err := getNums(data)
	if err != nil {
		log.Fatalf("unable to get nums: %v", err)
	}

	fuel, pos, err := getMinFuel("part1", nums)
	if err != nil {
		log.Fatalf("unable to get min fuel for part 1: %v", err)
	}

	fmt.Printf("part 1 answer: pos: %d, fuel:%d\n", pos, fuel)

	fuel, pos, err = getMinFuel("part2", nums)
	if err != nil {
		log.Fatalf("unable to get min fuel for part 1: %v", err)
	}
	fmt.Printf("part 2 answer: pos: %d, fuel:%d\n", pos, fuel)
}

// for each fuel we iterate through the rest of the positions to get
// the least min fuels, O(k^2) where k is the # of unique numbers
func getMinFuel(part string, nums map[int]int) (int, int, error) {
	minFuel := math.MaxInt
	minPos := -1
	var (
		// get bounds to search
		min int = math.MaxInt
		max int = -1
	)
	for k := range nums {
		if k < min {
			min = nums[k]
			continue
		}

		if k > max {
			max = k
		}
	}
	// search within bounds. this is just an assumption of how the problem
	// should be done. its not really clear on how to actually do it based
	// on the directions
	for i := min; i < max; i++ {
		curFuel := 0
		for pos, count := range nums {
			if pos == i {
				continue
			}

			stepper, err := getStepper(part)
			if err != nil {
				return 0, 0, fmt.Errorf("unable to get stepper: %w", err)
			}
			curFuel += stepper(i, pos, count)
		}
		if curFuel < minFuel {
			minFuel = curFuel
			minPos = i
		}
		// if we didnt find a new pos after the previous check
		// we can stop
		if minPos < i {
			break
		}
	}

	return minFuel, minPos, nil
}

func getNums(data []byte) (map[int]int, error) {
	numsStr := strings.Split(strings.TrimSpace(string(data)), ",")
	// position -> count
	nums := make(map[int]int)
	for i := range numsStr {
		n, err := strconv.Atoi(numsStr[i])
		if err != nil {
			return nil, fmt.Errorf("unable to convert num: %w", err)
		}
		nums[n]++
	}

	return nums, nil
}

type stepper func(alignAt, pos, count int) int

func getStepper(part string) (stepper, error) {
	switch part {
	case "part1":
		return func(alignAt, pos, count int) int {
			return absSub(alignAt, pos) * count
		}, nil
	case "part2":
		return func(alignAt, pos, count int) int {
			steps := absSub(alignAt, pos)
			// n(n+1)/2
			return ((steps * (steps + 1)) / 2) * count
		}, nil
	default:
		return nil, errors.New("unsupported part")
	}
}

func absSub(i, j int) int {
	if i > j {
		return i - j
	}

	return j - i
}
