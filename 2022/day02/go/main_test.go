package main

import "testing"

func TestSampleInput(t *testing.T) {
	picks := ParseRoundPicks("test_input.txt")
	result := GetTotalScore(picks)
	if result != 15 {
		t.Errorf("Expected %d, got %d", 15, result)
	}
}

func TestInput(t *testing.T) {
	picks := ParseRoundPicks("input.txt")
	result := GetTotalScore(picks)
	if result != 13675 {
		t.Errorf("Expected %d, got %d", 13675, result)
	}
}

func TestSampleInputPart2(t *testing.T) {
	picks := ParseRoundOutcomes("test_input.txt")
	result := GetTotalScoreForOutcome(picks)
	if result != 12 {
		t.Errorf("Expected %d, got %d", 12, result)
	}
}

func TestInputPart2(t *testing.T) {
	picks := ParseRoundOutcomes("input.txt")
	result := GetTotalScoreForOutcome(picks)
	if result != 14184 {
		t.Errorf("Expected %d, got %d", 14184, result)
	}
}
