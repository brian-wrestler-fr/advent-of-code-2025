package main

import (
	"fmt"
	"strconv"

	"github.com/brian-wrestler-fr/advent-of-code-2025/pkg/util"
)

func main() {
	batteryArray := util.ReadLines("./input")
	processBatteryArrays(batteryArray)
}

func processBatteryArrays(batteryArray []string) int64 {
	var outputVoltage int64 = 0

	for _, battery := range batteryArray {
		var digits [12]int

		for i := range battery {
			cellInt, _ := strconv.Atoi(string(battery[i]))
			cellsLeft := len(battery) - (i + 1)

			for j := range digits {
				cellsToFill := 12 - (j + 1)
				if digits[j] < cellInt && cellsLeft >= cellsToFill {
					digits[j] = cellInt

					newSlice := make([]int, cellsLeft)
					if cellsLeft > 0 {
						copy(digits[j+1:], newSlice)
					}

					break
				}
			}

		}

		var batteryTotal int64 = 0
		for i := 0; i < len(digits); i++ {
			batteryTotal = batteryTotal*10 + int64(digits[i])
		}

		fmt.Printf("In %s, largest joltage possible is: %d\n", battery, batteryTotal)
		outputVoltage += batteryTotal
	}

	fmt.Printf("Output Voltage: %d\n", outputVoltage)

	return outputVoltage
}
