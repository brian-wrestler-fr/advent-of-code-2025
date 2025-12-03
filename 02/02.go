package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/brian-wrestler-fr/advent-of-code-2025/pkg/util"
)

func main() {
	var rangeArray = strings.Split(util.ReadInput("./input"), ",")
	resultProblem1 := processRangeArrayProblem1(rangeArray)
	resultProblem2 := processRangeArrayProblem2(rangeArray)

	fmt.Printf("[Problem 1] The sum of the invalid ids is: %d\n", resultProblem1)
	fmt.Printf("[Problem 2] The sum of the invalid ids is: %d\n", resultProblem2)
}

func checkArrayEquality(arr []string) bool {
	for element := range arr {
		// skip empty elements. SplitAfter can create these
		if arr[element] != "" && arr[element] != arr[0] {
			return false
		}
	}

	return true
}

func processRangeProblem2(start int, end int) int64 {
	var sum int64 = 0
	for id := start; id <= end; id++ {
		idStr := strconv.Itoa(id)

		for i := 0; i <= len(idStr)/2 && len(idStr) > 1; i++ {
			splits := strings.SplitAfter(idStr, string(idStr[0:i]))

			// Make sure they are all equal and over 1 character long
			if checkArrayEquality(splits) {
				fmt.Printf("%d is a invalid id\n", id)
				sum += int64(id)
				break
			}
		}
	}

	return sum
}

func processRangeArrayProblem2(rangeArray []string) int64 {
	var result int64 = 0
	for _, idRangeStr := range rangeArray {
		idRange := strings.Split(idRangeStr, "-")
		start, _ := strconv.Atoi(idRange[0])
		end, _ := strconv.Atoi(idRange[1])

		result += processRangeProblem2(start, end)
	}

	return result
}

func processRangeProblem1(start int, end int) int64 {
	var sum int64 = 0
	for id := start; id <= end; id++ {
		// Turn into a string so we can split it in half easily
		idStr := strconv.Itoa(id)
		firstHalf := idStr[:len(idStr)/2]
		secondHalf := idStr[len(idStr)/2:]

		if firstHalf == secondHalf {
			fmt.Printf("%d is a invalid id\n", id)
			sum += int64(id)
		}
	}

	return sum
}

func processRangeArrayProblem1(rangeArray []string) int64 {
	var result int64 = 0
	for _, idRangeStr := range rangeArray {
		idRange := strings.Split(idRangeStr, "-")
		start, _ := strconv.Atoi(idRange[0])
		end, _ := strconv.Atoi(idRange[1])

		result += processRangeProblem1(start, end)
	}

	return result
}
