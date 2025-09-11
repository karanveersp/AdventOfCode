package main

import "testing"

func TestSampleInput(t *testing.T) {
	carriers := Parse("test_input.txt")
	res := GetMaxCalories(carriers)
	expected := 24000
	if res != expected {
		t.Fatalf("got %v wanted %v", res, expected)
	}
}

func TestInput(t *testing.T) {
	carriers := Parse("input.txt")
	res := GetMaxCalories(carriers)
	expected := 68787
	if res != expected {
		t.Fatalf("got %v wanted %v", res, expected)
	}
}

func TestSampleInputPart2(t *testing.T) {
	carriers := Parse("test_input.txt")
	res := GetTopThreeCarriersTotal(carriers)
	expected := 45000
	if res != expected {
		t.Fatalf("got %v wanted %v", res, expected)
	}
}

func TestInputPart2(t *testing.T) {
	carriers := Parse("input.txt")
	res := GetTopThreeCarriersTotal(carriers)
	expected := 198041
	if res != expected {
		t.Fatalf("got %v wanted %v", res, expected)
	}
}
