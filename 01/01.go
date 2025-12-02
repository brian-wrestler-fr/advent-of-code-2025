package main

import (
	_ "embed"
	"fmt"
	"strconv"

	"github.com/brian-wrestler-fr/advent-of-code-2025/pkg/util"
)

const dialSize = 100

type rotateStep struct {
	direction rune
	stepCount int
}

func main() {
	var rawManifest = util.ReadLines("./input")
	var parsedManifest = parseManifest(rawManifest)

	passCount, exactCount := processManifest(parsedManifest)

	fmt.Printf("The dial pointed at 0 a total of %d times.\n", exactCount)
	fmt.Printf("The dial passed by 0 a total of %d times.\n", passCount)

}

func processManifest(manifest []rotateStep) (passCount int, exactCount int) {
	passCount, exactCount = 0, 0

	currentPoint := 50

	fmt.Printf("The dial starts by pointing at %d.\n", currentPoint)
	for index, step := range manifest {
		var passed int
		var newPoint int

		if step.direction == 'L' {
			passed, newPoint = leftRotate(currentPoint, step.stepCount)

		} else if step.direction == 'R' {
			passed, newPoint = rightRotate(currentPoint, step.stepCount)
		}

		passCount += passed
		currentPoint = newPoint

		if currentPoint == 0 {
			exactCount++
		}

		fmt.Printf("Step %d: The dial is rotated %s %s to point at %s. During its rotation it passed zero: %d times.\n", index+1, string(step.direction), strconv.Itoa(step.stepCount), strconv.Itoa(currentPoint), passed)
	}

	return passCount, exactCount
}

func parseManifest(rawManifest []string) []rotateStep {
	manifest := []rotateStep{}

	for _, element := range rawManifest {
		clicks, _ := strconv.Atoi(element[1:])

		rs := rotateStep{
			rune(element[0]),
			clicks,
		}

		manifest = append(manifest, rs)
	}

	return manifest
}

func leftRotate(currentPoint int, stepCount int) (passCount int, endingPoint int) {
	remainder := stepCount % dialSize
	passCount = stepCount / dialSize
	finalPoint := currentPoint - remainder

	if stepCount == 5 {
		// Debugging aid
		_ = finalPoint
	}

	if finalPoint < 0 {
		finalPoint = (finalPoint + dialSize)
		if currentPoint > 0 {
			passCount++
		}
	}

	if finalPoint == 0 && remainder > 0 {
		passCount++
	}

	return passCount, finalPoint
}

func rightRotate(currentPoint int, stepCount int) (passCount int, endingPoint int) {
	remainder := stepCount % dialSize
	passCount = stepCount / dialSize
	finalPoint := currentPoint + remainder

	if stepCount == 60 {
		// Debugging aid
		_ = finalPoint
	}

	if finalPoint >= dialSize {
		finalPoint = finalPoint % dialSize
		passCount++
	}

	return passCount, finalPoint
}
