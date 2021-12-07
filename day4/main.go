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
	input       = "day4/input.txt"
	metaCellIdx = 5
)

// one extra row/column to store the # of numbers that have been marked
type grid [6][6]cell

func (g *grid) markCall(call int) {
	for i := 0; i < metaCellIdx; i++ {
		for j := 0; j < metaCellIdx; j++ {
			if g[i][j].num == call {
				g[i][j].marked = true
				g[i][metaCellIdx].num++
				g[metaCellIdx][j].num++
			}
		}
	}
}

func (g *grid) hasWon() bool {
	for i := 0; i < metaCellIdx; i++ {
		if g[i][metaCellIdx].num == 5 || g[metaCellIdx][i].num == 5 {
			return true
		}
	}
	return false
}

func (g *grid) score(justCalled int) int {
	var sum int
	for i := 0; i < metaCellIdx; i++ {
		for j := 0; j < metaCellIdx; j++ {
			if !g[i][j].marked {
				sum += g[i][j].num
			}
		}
	}

	return sum * justCalled
}

type cell struct {
	num    int
	marked bool
	meta   bool
}

// straight forward implementation for both parts. Only trick i used was
// having an extra row and column that contained the # of marked numbers in
// that row /column. That we can only reference those to see if a board has one
func main() {
	if err := part1(); err != nil {
		log.Fatalf("unable to complete part 1: %v", err)
	}

	if err := part2(); err != nil {
		log.Fatalf("unable to complete part 2: %v", err)
	}

}

func part1() error {
	s, err := inputScanner(input)
	if err != nil {
		log.Fatalf("unable to get input scanner: %v", err)
	}

	calls, err := getCallouts(s)
	if err != nil {
		return fmt.Errorf("unable to get calls: %w", err)
	}

	grids, err := getGrids(s)
	if err != nil {
		return fmt.Errorf("unable to get grids: %w", err)
	}

	if err := s.Err(); err != nil {
		log.Fatalf("encountered error while scanning: %v", err)
	}

	for i := range calls {
		for j := range grids {
			grids[j].markCall(calls[i])
			if grids[j].hasWon() {
				fmt.Printf("Grid %d has won!\n", j+1)
				fmt.Printf("pt 1 answer: %d\n", grids[j].score(calls[i]))
				return nil
			}
		}
	}

	return nil
}

func part2() error {
	s, err := inputScanner(input)
	if err != nil {
		log.Fatalf("unable to get input scanner: %v", err)
	}

	calls, err := getCallouts(s)
	if err != nil {
		return fmt.Errorf("unable to get calls: %w", err)
	}

	grids, err := getGrids(s)
	if err != nil {
		return fmt.Errorf("unable to get grids: %w", err)
	}

	var (
		lastGrid *grid
		lastCall int
		wins     int
	)
findLast:
	for i := range calls {
		n := calls[i]
		for j := 0; j < len(grids); j++ {
			grids[j].markCall(n)
			if grids[j].hasWon() {
				lastGrid = &grids[j]
				lastCall = n
				wins++
				switch len(grids) {
				case 1:
					break findLast
				default:
					grids = append(grids[:j], grids[j+1:]...)
					j--
				}
			}
		}
	}

	if lastGrid != nil {
		fmt.Printf("pt 2 answer: %d\n", lastGrid.score(lastCall))
	}

	return nil
}

func getGrids(s *bufio.Scanner) ([]grid, error) {
	var grids []grid
	var rows int
	var g grid
	for s.Scan() {
		// reset
		if rows == 5 {
			g[metaCellIdx] = metaRow()
			grids = append(grids, g)
			rows = 0
			g = grid{}
		}

		l := s.Text()
		if l == "" {
			continue
		}
		row, err := gridRow(strings.TrimSpace(l))
		if err != nil {
			return nil, fmt.Errorf("unable to get grid row: %w", err)
		}
		g[rows] = row
		rows++
	}

	if rows == 5 {
		grids = append(grids, g)
	}

	return grids, nil
}

func getCallouts(s *bufio.Scanner) ([]int, error) {
	var calls []int
	if s.Scan() {
		strs := strings.Split(strings.TrimSpace(s.Text()), ",")
		for i := range strs {
			if strs[i] == "" {
				continue
			}
			n, err := strconv.Atoi(strs[i])
			if err != nil {
				return nil, fmt.Errorf("expected a number: %v", err)
			}
			calls = append(calls, n)
		}
	}

	return calls, nil
}

func metaRow() [6]cell {
	return [6]cell{{meta: true}, {meta: true}, {meta: true}, {meta: true}, {meta: true}, {meta: true}}
}

func gridRow(s string) ([6]cell, error) {
	nums := strings.Split(s, " ")

	var cells [6]cell
	var ptr int
	for i := range nums {
		if nums[i] == "" {
			continue
		}
		n, err := strconv.Atoi(nums[i])
		if err != nil {
			return [6]cell{}, fmt.Errorf("unable to convert to number: %w", err)
		}
		cells[ptr] = cell{num: n}
		ptr++
	}

	// add one more for the meta cell
	cells[metaCellIdx] = cell{meta: true}

	return cells, nil
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
