package main

import (
	"testing"
)

func TestProcessManifest_AoC_Example(t *testing.T) {

	testManifest := []rotateStep{
		{'L', 68},
		{'L', 30},
		{'R', 48},
		{'L', 5},
		{'R', 60},
		{'L', 55},
		{'L', 1},
		{'L', 99},
		{'R', 14},
		{'L', 82},
	}

	wantPassCount, wantExactCount := 6, 3

	passCount, exactCount := processManifest(testManifest)

	if exactCount != wantExactCount {
		t.Errorf("Exact count: got %d, want %d", exactCount, wantExactCount)
	}

	if passCount != wantPassCount {
		t.Errorf("Pass count: got %d, want %d", passCount, wantPassCount)
	}
}

func TestProcessManifest_Left_Two_Overlaps(t *testing.T) {

	testManifest := []rotateStep{
		{'L', 50},
		{'L', 100},
	}
	wantPassCount, wantExactCount := 2, 2

	passCount, exactCount := processManifest(testManifest)

	if exactCount != wantExactCount {
		t.Errorf("Exact count: got %d, want %d", exactCount, wantExactCount)
	}

	if passCount != wantPassCount {
		t.Errorf("Pass count: 	got %d, want %d", passCount, wantPassCount)
	}
}

func TestProcessManifest_Right_Two_Overlaps(t *testing.T) {

	testManifest := []rotateStep{
		{'R', 50},
		{'R', 100},
	}
	wantPassCount, wantExactCount := 2, 2

	passCount, exactCount := processManifest(testManifest)

	if exactCount != wantExactCount {
		t.Errorf("Exact count: got %d, want %d", exactCount, wantExactCount)
	}

	if passCount != wantPassCount {
		t.Errorf("Pass count: 	got %d, want %d", passCount, wantPassCount)
	}
}

func TestProcessManifest_Switch_Directions(t *testing.T) {

	testManifest := []rotateStep{
		{'L', 50},
		{'L', 5},
		{'R', 60},
	}
	wantPassCount, wantExactCount := 2, 1

	passCount, exactCount := processManifest(testManifest)

	if exactCount != wantExactCount {
		t.Errorf("Exact count: got %d, want %d", exactCount, wantExactCount)
	}

	if passCount != wantPassCount {
		t.Errorf("Pass count: 	got %d, want %d", passCount, wantPassCount)
	}
}

func TestProcessManifest_Left_StepExceedsDial(t *testing.T) {

	testManifest := []rotateStep{
		{'L', 150},
	}
	wantPassCount, wantExactCount := 2, 1

	passCount, exactCount := processManifest(testManifest)

	if exactCount != wantExactCount {
		t.Errorf("Exact count: got %d, want %d", exactCount, wantExactCount)
	}

	if passCount != wantPassCount {
		t.Errorf("Pass count: 	got %d, want %d", passCount, wantPassCount)
	}
}

func TestProcessManifest_Right_StepExceedsDialMultiple(t *testing.T) {

	testManifest := []rotateStep{
		{'R', 50},
		{'R', 500}, //5
	}
	wantPassCount, wantExactCount := 6, 2

	passCount, exactCount := processManifest(testManifest)

	if exactCount != wantExactCount {
		t.Errorf("Exact count: got %d, want %d", exactCount, wantExactCount)
	}

	if passCount != wantPassCount {
		t.Errorf("Pass count: got %d, want %d", passCount, wantPassCount)
	}
}

func TestProcessManifest_Right_StepExceedsDial(t *testing.T) {

	testManifest := []rotateStep{
		{'R', 150},
	}
	wantPassCount, wantExactCount := 2, 1

	passCount, exactCount := processManifest(testManifest)

	if exactCount != wantExactCount {
		t.Errorf("Exact count: got %d, want %d", exactCount, wantExactCount)
	}

	if passCount != wantPassCount {
		t.Errorf("Pass count: got %d, want %d", passCount, wantPassCount)
	}
}

func TestProcessManifest_Mixed(t *testing.T) {

	testManifest := []rotateStep{
		{'L', 150},
		{'R', 50},
		{'R', 50},
		{'L', 25},
		{'R', 575},
	}
	wantPassCount, wantExactCount := 9, 2

	passCount, exactCount := processManifest(testManifest)

	if exactCount != wantExactCount {
		t.Errorf("Exact count: got %d, want %d", exactCount, wantExactCount)
	}

	if passCount != wantPassCount {
		t.Errorf("Pass count: got %d, want %d", passCount, wantPassCount)
	}
}
