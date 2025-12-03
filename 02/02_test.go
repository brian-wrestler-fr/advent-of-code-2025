package main

import (
	"testing"
)

func TestProcessRanges_AoC_Example(t *testing.T) {

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

	result := processRangeArray(rangeArray)

	if result != want {
		t.Errorf("The result was %d but wanted %d", result, want)
	}

}
