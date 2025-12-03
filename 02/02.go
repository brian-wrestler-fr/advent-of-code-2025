package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brian-wrestler-fr/advent-of-code-2025/pkg/util"
)

func main() {
	var rangeArray = strings.Split(util.ReadInput("./input"), ",")
	result := processRangeArray(rangeArray)

	fmt.Printf("The sum of the invalid ids is: %d\n", result)
}

// Process the range. We will convert each number to a string, split it in half, and then XOR  the halves. If 0 they are the same
func processRange(start int, end int) int64 {
	var sum int64 = 0
	for id := start; id <= end; id++ {
		// Turn into a string so we can split it in half easily
		idStr := strconv.Itoa(id)
		firstHalf := idStr[:len(idStr)/2]
		secondHalf := idStr[len(idStr)/2:]

		if firstHalf == secondHalf {
			fmt.Printf("%d is an invalid id\n", id)
			sum += int64(id)
		}
	}

	return sum
}

func processRangeArray(rangeArray []string) int64 {
	var result int64 = 0
	for _, idRangeStr := range rangeArray {
		idRange := strings.Split(idRangeStr, "-")
		start, _ := strconv.Atoi(idRange[0])
		end, _ := strconv.Atoi(idRange[1])

		result += processRange(start, end)
	}

	return result
}
