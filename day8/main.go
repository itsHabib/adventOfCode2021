package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const input = "day8/input.txt"

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
		return fmt.Errorf("unable to get scanner: %w", err)
	}

	counter := make(map[int]int)
	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		// assuming clean output
		// there are 10 unique words, followed by a  |, then the output. We are
		// only interested in the output for part 1, so we discard the rest
		output := strings.Split(l, " ")[11:]

		for i := range output {
			counter[getNum(len(output[i]))]++
		}
	}

	var sum int
	for k, v := range counter {
		if k == -1 {
			continue
		}

		sum += v
	}

	fmt.Printf("answer: %d\n", sum)

	return nil
}

// use the length of segments left after plucking out the unique numbers
func part2() error {
	s, err := inputScanner(input)
	if err != nil {
		return fmt.Errorf("unable to get scanner: %w", err)
	}

	var sum int
	for s.Scan() {
		parts := strings.Split(strings.TrimSpace(s.Text()), " ")
		input := parts[:10]
		output := parts[11:]

		unique := formunique(input)
		n, err := getOutputNum(unique, output)
		if err != nil {
			return fmt.Errorf("unable to get output num: %w", err)
		}

		sum += n
	}

	if err := s.Err(); err != nil {
		return fmt.Errorf("encountered scan err: %w", err)
	}

	fmt.Printf("answer: %d\n", sum)

	return nil
}

func formunique(input []string) map[int]string {
	// sort by length asc
	sort.Slice(input, func(i, j int) bool {
		return len(input[i]) < len(input[j])
	})

	unique := make(map[int]string)
	for i := range input {
		// found all unique
		if len(unique) == 4 {
			break
		}
		letters := strings.Split(input[i], "")
		sort.Slice(letters, func(i, j int) bool {
			return letters[i] < letters[j]
		})
		str := strings.Join(letters, "")

		switch len(input[i]) {
		case 2:
			unique[1] = str
		case 3:
			unique[7] = str
		case 4:
			unique[4] = str
		case 7:
			unique[8] = str
		}
	}

	return unique
}

func getOutputNum(unique map[int]string, output []string) (int, error) {
	var nums []string

	for i := range output {
		switch len(output[i]) {
		case 2:
			nums = append(nums, "1")
		case 3:
			nums = append(nums, "7")
		case 4:
			nums = append(nums, "4")
		case 5:
			// 2, 3, or 5
			// 2 can be made up of setting on #4(minus c) and adding 3 segments
			// 3 can be made up of #1 and adding 3 segments. (we can just check
			// if one is set)
			// 5 can be made up of #4(minus c) and adding 2 segments

			// try to strip #1 from the number if it contains it, otherwise
			// strip #4 characters
			if hasLetters(output[i], unique[1]) {
				// found a 3
				nums = append(nums, "3")
				continue
			}

			cut := pluckString(output[i], unique[4])
			switch len(cut) {
			case 2:
				// found a 5
				nums = append(nums, "5")
			case 3:
				// found a 2
				nums = append(nums, "2")
			}
		case 6:
			// 0, 6, or 9
			// 0 can be made up of setting on #1 and adding 4 segments
			// 6 can be made up of setting on #1(minus c)  and adding 5
			// segments
			// 9 can be made up of setting on #4 (we can just check if 4 is set)

			// first check if 4 is set, if so we found a 9
			if hasLetters(output[i], unique[4]) {
				nums = append(nums, "9")
				continue
			}

			cut := pluckString(output[i], unique[1])
			switch len(cut) {
			case 4:
				nums = append(nums, "0")
			case 5:
				nums = append(nums, "6")
			}
		case 7:
			nums = append(nums, "8")
		}
	}

	n, err := strconv.Atoi(strings.Join(nums, ""))
	if err != nil {
		return 0, fmt.Errorf("unable to create num: %w", err)
	}

	return n, nil
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

func getNum(length int) int {
	m := map[int]int{
		2: 1,
		3: 7,
		4: 4,
		7: 8,
	}

	n, ok := m[length]
	if !ok {
		return -1
	}

	return n
}

func hasLetters(str, has string) bool {
	m := make(map[uint8]struct{})
	for i := range has {
		m[has[i]] = struct{}{}
	}

	for i := range str {
		if _, ok := m[str[i]]; ok {
			delete(m, str[i])
		}

		if len(m) == 0 {
			return true
		}
	}

	return false
}

func pluckString(start, toRemove string) string {
	s := new(strings.Builder)

	m := make(map[uint8]struct{})
	for i := range toRemove {
		// no duplicate letters so no issue there
		m[toRemove[i]] = struct{}{}
	}

	for i := range start {
		if _, ok := m[start[i]]; !ok {
			// shouldn't error so not checking
			s.WriteRune(rune(start[i]))
		}
	}

	return s.String()
}
