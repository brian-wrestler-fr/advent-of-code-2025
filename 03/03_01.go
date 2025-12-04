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

func processBatteryArrays(batteryArray []string) int {
	outputVoltage := 0

	for _, battery := range batteryArray {
		firstDigit := 0
		secondDigit := 0

		for i := 0; i < len(battery); i++ {
			cellInt, _ := strconv.Atoi(string(battery[i]))

			if firstDigit < cellInt && i < len(battery)-1 {
				firstDigit = cellInt
				secondDigit = 0
				continue
			}

			if secondDigit < cellInt {
				secondDigit = cellInt
			}
		}

		batteryTotal := firstDigit*10 + secondDigit

		fmt.Printf("In %s, largest joltage possible is: %d\n", battery, batteryTotal)
		outputVoltage += batteryTotal
	}

	fmt.Printf("Output Voltage: %d", outputVoltage)

	return outputVoltage
}
