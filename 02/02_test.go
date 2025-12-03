package main

import (
	"testing"
)

func TestProcessRanges_AoC_Example_1(t *testing.T) {
	rangeArray := []string{
		"11-22",
		"95-115",
		"998-1012",
		"1188511880-1188511890",
		"222220-222224",
		"1698522-1698528",
		"446443-446449",
		"38593856-38593862",
	}

	var want int64 = 1227775554

	result := processRangeArrayProblem1(rangeArray)

	if result != want {
		t.Errorf("The result was %d but wanted %d", result, want)
	}
}

func TestProcessRanges_AoC_Single_Digts(t *testing.T) {
	rangeArray := []string{
		"2-12",
	}

	var want int64 = 11

	result := processRangeArrayProblem2(rangeArray)

	if result != want {
		t.Errorf("The result was %d but wanted %d", result, want)
	}
}

func TestProcessRanges_AoC_Example_2(t *testing.T) {
	rangeArray := []string{
		"11-22",
		"95-115",
		"998-1012",
		"1188511880-1188511890",
		"222220-222224",
		"1698522-1698528",
		"446443-446449",
		"38593856-38593862",
		"565653-565659",
		"824824821-824824827",
		"2121212118-2121212124",
	}

	var want int64 = 4174379265

	result := processRangeArrayProblem2(rangeArray)

	if result != want {
		t.Errorf("The result was %d but wanted %d", result, want)
	}
}

func TestProcessRanges_AoC_Tripple(t *testing.T) {
	rangeArray := []string{
		"121312131213-121312131214",
	}

	var want int64 = 121312131213

	result := processRangeArrayProblem2(rangeArray)

	if result != want {
		t.Errorf("The result was %d but wanted %d", result, want)
	}
}
