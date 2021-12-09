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

type point struct {
	x int
	y int
}

type slope struct {
	x int
	y int
}

func (p *point) travel(slope *slope) {
	if slope == nil {
		return
	}

	p.x += slope.x
	p.y += slope.y
}

const input = "day5/input.txt"

// keep a track of all points we have seen. for each point we travel to the
// destination point and keep track of the points we see along the way.
// if we interact with any that have already been seen, that counts as an
// overlap
func main() {
	if err := part1(); err != nil {
		log.Fatalf("unable to complete part1: %v", err)
	}

	if err := part2(); err != nil {
		log.Fatalf("unable to complete part2: %v", err)
	}
}

func part1() error {
	s, err := inputScanner(input)
	if err != nil {
		log.Fatalf("unable to get input scanner: %v", err)
	}

	// track the vertices we've already seen. if we encounter one already seen
	// that counts as an overlap
	seen := make(map[point]int)
	for s.Scan() {
		l := s.Text()
		from, to, err := getPoints(strings.TrimSpace(l))
		if err != nil {
			log.Fatalf("unable to get points from line: %v", err)
		}

		// part 1 rule
		if from.x != to.x && from.y != to.y {
			continue
		}

		// form vertices to go from->to
		// add to seen map if we havent seen em
		lineSlope := getUnitSlope(from, to)
		to.travel(lineSlope)
		for *from != *to {
			_, ok := seen[*from]
			//fmt.Println(from.String())
			if ok {
				seen[*from]++
			} else {
				seen[*from] = 1
			}
			from.travel(lineSlope)
		}
	}

	if err := s.Err(); err != nil {
		log.Fatalf("encountered error while scanning: %v", err)
	}

	var twoOrMore int
	for _, v := range seen {
		if v > 1 {
			twoOrMore++
		}
	}

	fmt.Printf("answer: %d\n", twoOrMore)

	return nil
}

func part2() error {
	s, err := inputScanner(input)
	if err != nil {
		log.Fatalf("unable to get input scanner: %v", err)
	}

	// track the vertices we've already seen. if we encounter one already seen
	// that counts as an overlap
	seen := make(map[point]int)
	for s.Scan() {
		l := s.Text()
		from, to, err := getPoints(strings.TrimSpace(l))
		if err != nil {
			log.Fatalf("unable to get points from line: %v", err)
		}

		// form vertices to go from->to
		// add to seen map if we havent seen em
		lineSlope := getUnitSlope(from, to)
		to.travel(lineSlope)
		for *from != *to {
			_, ok := seen[*from]
			//fmt.Println(from.String())
			if ok {
				seen[*from]++
			} else {
				seen[*from] = 1
			}
			from.travel(lineSlope)
		}
	}

	if err := s.Err(); err != nil {
		log.Fatalf("encountered error while scanning: %v", err)
	}

	var twoOrMore int
	for _, v := range seen {
		if v > 1 {
			twoOrMore++
		}
	}

	fmt.Printf("answer: %d\n", twoOrMore)

	return nil
}

func getUnitSlope(from *point, to *point) *slope {
	if from == nil || to == nil {
		return nil
	}

	// divide by itself because we only want to move in 1s
	x := to.x - from.x
	if x != 0 {
		x /= absSub(to.x, from.x)
	}
	y := to.y - from.y
	if y != 0 {
		y /= absSub(to.y, from.y)
	}
	return &slope{
		// divide by itself because we only want to move in 1s
		x: x,
		y: y,
	}
}

func getPoints(line string) (*point, *point, error) {
	parts := strings.Split(line, " -> ")
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("unexpected form of line: %s", line)
	}

	from, err := getPoint(parts[0])
	if err != nil {
		return nil, nil, fmt.Errorf("unable to get from coord: %w", err)
	}

	to, err := getPoint(parts[1])
	if err != nil {
		return nil, nil, fmt.Errorf("unable to get to coord: %w", err)
	}

	return &from, &to, nil
}

func getPoint(coords string) (point, error) {
	parts := strings.Split(coords, ",")
	if len(parts) != 2 {
		return point{}, fmt.Errorf("unexpected form of point: %s", coords)
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return point{}, fmt.Errorf("unable to get x coord: %w", err)
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return point{}, fmt.Errorf("unable to get y coord: %w", err)
	}

	return point{
		x: x,
		y: y,
	}, nil
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

func absSub(i, j int) int {
	if i > j {
		return i - j
	}

	return j - i
}
