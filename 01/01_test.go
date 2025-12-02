package main

import (
	"testing"
)

func Test_example(t *testing.T) {

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
	want := 3

	var result = processManifest(testManifest)

	if result != want {
		t.Errorf("got %d, want %d", result, want)
	}
}

func Test_leftRotate(t *testing.T) {

	testManifest := []rotateStep{
		{'L', 150},
	}
	want := 1

	var result = processManifest(testManifest)

	if result != want {
		t.Errorf("got %d, want %d", result, want)
	}
}
