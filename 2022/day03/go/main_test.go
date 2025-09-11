package main

import (
	"testing"
)

func TestPart1Sample(t *testing.T) {
	sum, err := FindIntersectionSum("test_input.txt")
	if err != nil {
		t.Error(err)
	}
	if sum != 157 {
		t.Errorf("got %d, want %d", sum, 157)
	}
}

func TestPart1(t *testing.T) {
	sum, err := FindIntersectionSum("input.txt")
	if err != nil {
		t.Error(err)
	}
	if sum != 8202 {
		t.Errorf("got %d, want %d", sum, 157)
	}
}

func TestPart1_FindIntersectionPriority(t *testing.T) {
	line := "vJrwpWtwJgWrhcsFMMfFFhFp"
	i, err := FindIntersectionPriority(line)
	if err != nil {
		t.Errorf("FindIntersection() error = %v", err)
	}
	if i != 16 {
		t.Errorf("FindIntersection() want 16, got %v", i)
	}
}

func TestPart2_FindBadgesSample(t *testing.T) {
	sum, err := FindBadgesSum("test_input.txt")
	if err != nil {
		t.Error(err)
	}
	if sum != 70 {
		t.Errorf("got %d, want %d", sum, 70)
	}
}

func TestPart2_FindBadges(t *testing.T) {
	sum, err := FindBadgesSum("input.txt")
	if err != nil {
		t.Error(err)
	}
	if sum != 2864 {
		t.Errorf("got %d, want %d", sum, 2864)
	}
}

func TestPart2_FindIntersection(t *testing.T) {
	values := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
	}

	i := FindIntersections(values...)
	if i != "r" {
		t.Errorf("FindIntersections() want r, got %v", i)
	}

}

func TestPart2_FindIntersection2(t *testing.T) {
	values := []string{
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}

	i := FindIntersections(values...)
	if i != "Z" {
		t.Errorf("FindIntersections() want Z, got %v", i)
	}

}
