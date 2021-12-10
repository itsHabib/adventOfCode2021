package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const input = "day6/input.txt"

// use an array to track the number of lanternfish in which each index
// represents the ith latern fish and the value of at each index represents
// the total count in i.
func main() {
	inputB, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatalf("unable to read file")
	}

	numsStr := strings.Split(string(inputB), ",")
	fmt.Println(string(inputB))
	// 0 - 8
	nums := make([]int, 9)
	for i := range numsStr {
		n, err := strconv.Atoi(numsStr[i])
		if err != nil {
			log.Fatalf("unable to convert to num: %v", err)
		}
		nums[n]++
	}

	// simulate through the days
	const days = 256
	var new int
	for i := 0; i < days; i++ {
		if i == 80 {
			fmt.Printf("part 1 answer: %d\n", total(nums))
		}
		for j := range nums {
			if nums[j] == 0 {
				continue
			}

			if j == 0 && nums[j] > 0 {
				// add  amount to the new count
				new += nums[j]
				// set count to 0
				nums[j] = 0

				continue
			}
			//  add to the one below, reset count to zero
			nums[j-1] += nums[j]
			nums[j] = 0
		}

		// add new count to the 8s
		nums[8] += new
		// add to #6 key with the new amount
		nums[6] += new
		new = 0
	}

	fmt.Printf("part 2 answer: %d\n", total(nums))
}

func total(nums []int) int {
	var sum int
	for j := range nums {
		sum += nums[j]
	}

	return sum
}
