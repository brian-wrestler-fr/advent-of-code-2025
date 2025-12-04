package main

import (
	"testing"
)

func TestProcessBatteryArrays(t *testing.T) {
	batteryArray := []string{
		"987654321111111",
		"811111111111119",
		"234234234234278",
		"818181911112111"}

	var want int64 = 3121910778619

	got := processBatteryArrays(batteryArray)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestProcessBatteryArrays_Problem(t *testing.T) {
	batteryArray := []string{
		"811111111111119",
	}

	var want int64 = 811111111119

	got := processBatteryArrays(batteryArray)

	if got != want {
		t.Errorf("got %d, wanted %d\n", got, want)
	}
}
