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

const (
	input = "day1/input.txt"
)

func main() {
	part1Num, err := part1()
	if err != nil {
		log.Fatalf("unable to complete part 1: %v", err)
	}
	fmt.Println(part1Num)

	part2Num, err := part2()
	if err != nil {
		log.Fatalf("unable to complete part 2: %v", err)
	}
	fmt.Println(part2Num)
}

// go through each num keeping track of the last. if the last isnt nil and
// the current number is bigger, add 1 to the increasing var.
func part1() (int, error) {
	s, err := inputScanner(input)
	if err != nil {
		return 0, fmt.Errorf("unable to get scanner: %w", err)
	}

	var last *int
	var increasing int

	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		num, err := strconv.Atoi(line)
		if err != nil {
			return 0, fmt.Errorf("unable to convert: %w", err)
		}

		if last != nil && num > *last {
			increasing++
		}
		last = &num
	}

	if err := s.Err(); err != nil {
		return 0, fmt.Errorf("encountered error while scanning: %w", err)
	}

	return increasing, nil
}

// create a running sliding window and variables to track state as we go through
// the list of numbers. As we progress through each number add to the sum
// of all the running windows and move the slide after each number.
func part2() (int, error) {
	s, err := inputScanner(input)
	if err != nil {
		return 0, fmt.Errorf("unable to get scanner: %w", err)
	}

	// current tracks the current window
	var current int
	// windowIdx tracks the movement of the sliding window
	var windowIdx int
	// running holds the mapping of indices to the windows list
	running := []int{-1, -1, -1}
	// windows holds the sums
	var windows []int

	for s.Scan() {
		l := s.Text()
		num, err := strconv.Atoi(l)
		if err != nil {
			return 0, fmt.Errorf("unable to convert: %w", err)
		}

		running[windowIdx] = current

		for i := range running {
			if running[i] == -1 {
				continue
			}

			if running[i] >= len(windows) {
				windows = append(windows, 0)
			}

			windows[running[i]] += num
		}

		current++
		windowIdx = (windowIdx + 1) % 3
	}

	var increasing int
	var last *int
	for i := range windows {
		if last != nil && windows[i] > *last {
			increasing++
		}
		last = &windows[i]
	}

	return increasing, nil
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
