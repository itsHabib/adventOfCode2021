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

const input = "day2/input.txt"

type moves struct {
	forward int
	down    int
	up      int
}

func main() {
	if err := part1(); err != nil {
		log.Fatalf("unable to complete part 1: %v", err)
	}
	if err := part2(); err != nil {
		log.Fatalf("unable to complete part 2: %v", err)
	}
}

// increment each direction as we find them, get the depth by subtracting up and
// down.
func part1() error {
	s, err := inputScanner("day2/input.txt")
	if err != nil {
		log.Fatalf("unable to get input scanner: %v", err)
	}

	var m moves
	for s.Scan() {
		l := s.Text()
		dir, amount, err := getMove(strings.TrimSpace(l))
		if err != nil {
			return fmt.Errorf("unable to get move from input line: %w", err)
		}
		switch strings.ToLower(dir) {
		case "forward":
			m.forward += amount
		case "down":
			m.down += amount
		case "up":
			m.up += amount
		default:
			return fmt.Errorf("unexpected direction type: %s", dir)
		}
	}

	/*
		Calculate the horizontal position and depth you would have after following
		the planned course. What do you get if you multiply your final horizontal
		 position by your final depth?
	*/
	// multiply by -1 since we are in a submarine and down is positive
	depth := (m.up - m.down) * -1

	fmt.Println("depth: ", depth)
	fmt.Println("horizontal: ", m.forward)
	fmt.Println("answer: ", depth*m.forward)

	return nil
}

type moves2 struct {
	aim        int
	depth      int
	horizontal int
}

// track depth and aim as we encounter each direction
func part2() error {
	s, err := inputScanner("day2/input.txt")
	if err != nil {
		log.Fatalf("unable to get input scanner: %v", err)
	}

	var m moves2
	for s.Scan() {
		l := s.Text()
		dir, amount, err := getMove(strings.TrimSpace(l))
		if err != nil {
			return fmt.Errorf("unable to get move from input line: %w", err)
		}
		switch strings.ToLower(dir) {
		case "forward":
			m.horizontal += amount
			m.depth += m.aim * amount
		case "down":
			m.aim += amount
		case "up":
			m.aim -= amount
		default:
			return fmt.Errorf("unexpected direction type: %s", dir)
		}
	}

	fmt.Println("depth: ", m.depth)
	fmt.Println("horizontal: ", m.horizontal)
	fmt.Println("answer: ", m.depth*m.horizontal)

	return nil
}

func getMove(l string) (string, int, error) {
	parts := strings.Split(l, " ")
	if len(parts) != 2 {
		return "", 0, fmt.Errorf("unexpected line form: %s", l)
	}

	amount, err := strconv.Atoi(parts[1])
	if err != nil {
		return "", 0, fmt.Errorf("unable to create int ")
	}

	return parts[0], amount, nil
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
