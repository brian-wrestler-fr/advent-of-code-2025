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

	want := 357

	got := processBatteryArrays(batteryArray)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
