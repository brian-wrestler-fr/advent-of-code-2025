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
	var zeroCount = processManifest(parsedManifest)

	fmt.Printf("The dial pointed at 0 a total of %d times.\n", zeroCount)
}

func processManifest(manifest []rotateStep) int {
	zeroCount := 0

	currentPoint := 50

	fmt.Printf("The dial starts by pointing at %d.\n", currentPoint)
	for index, step := range manifest {
		if step.direction == 'L' {
			currentPoint = leftRotate(currentPoint, step.stepCount)
		} else if step.direction == 'R' {
			currentPoint = rightRotate(currentPoint, step.stepCount)
		}

		if currentPoint == 0 {
			zeroCount++
		}

		fmt.Printf("Step %d: The dial is rotated %s %s to point at %s.\n", index+1, string(step.direction), strconv.Itoa(step.stepCount), strconv.Itoa(currentPoint))
	}

	return zeroCount
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

func leftRotate(currentPoint int, stepCount int) int {
	// Rotate left. Ignore rap around for now
	unModdedPoint := currentPoint - stepCount

	// If positive we did not wrap.  Return it
	if unModdedPoint >= 0 {
		return unModdedPoint
	}

	// We wrapped mod it down to the remainder
	moddedPoint := -unModdedPoint % dialSize

	// If we modded to 0 return 0
	if moddedPoint == 0 {
		return moddedPoint
	}

	// Remove the remainder from 100 to get the wrapped point
	return 100 - moddedPoint

}

func rightRotate(currentPoint int, stepCount int) int {
	unModdedPoint := currentPoint + stepCount

	return unModdedPoint % dialSize

}
