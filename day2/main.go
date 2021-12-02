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
	s, err := inputScanner("day2/input.txt")
	if err != nil {
		log.Fatalf("unable to get input scanner: %v", err)
	}

	var m moves
	for s.Scan() {
		l := s.Text()
		dir, amount, err := getMove(strings.TrimSpace(l))
		if err != nil {
			log.Fatalf("unable to get move from input line: %v", err)
		}
		switch strings.ToLower(dir) {
		case "forward":
			m.forward += amount
		case "down":
			m.down += amount
		case "up":
			m.up += amount
		default:
			log.Fatalf("unexpected direction type: %s", dir)
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
