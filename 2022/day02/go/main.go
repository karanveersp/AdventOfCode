package main

import (
	"strings"

	"github.com/karanveersp/gobelt/pkg/file"
	"github.com/karanveersp/gobelt/pkg/slices"
)

type Shape int

// Values correspond to scores for each shape
const (
	Rock     Shape = 1
	Paper    Shape = 2
	Scissors Shape = 3
)

func (s Shape) Beats(o Shape) bool {
	if s == Rock && o == Scissors {
		return true
	}
	if s == Scissors && o == Paper {
		return true
	}
	if s == Paper && o == Rock {
		return true
	}
	return false
}

// BeatsBy returns the shape that beats the current shape.
func (s Shape) BeatBy() Shape {
	switch s {
	case Rock:
		return Paper
	case Paper:
		return Scissors
	default:
		return Rock
	}
}

func (s Shape) Score() int {
	return int(s)
}

type Outcome int

const (
	Lose Outcome = 0
	Draw Outcome = 3
	Win  Outcome = 6
)

var loseScore = 0
var drawScore = 3
var winScore = 6

var fstColMap = map[string]Shape{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var sndColMap = map[string]Shape{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var sndColMapOutcome = map[string]Outcome{
	"X": Lose,
	"Y": Draw,
	"Z": Win,
}

// RunRound returns p1 and p2 score from the round.
func RunRound(p1 Shape, p2 Shape) (int, int) {
	if p1 == p2 {
		return p1.Score() + drawScore, p2.Score() + drawScore
	}
	if p1.Beats(p2) {
		return p1.Score() + winScore, p2.Score() + loseScore
	}
	return p1.Score() + loseScore, p2.Score() + winScore
}

type RoundPicks struct {
	P1 Shape
	P2 Shape
}

func ParseRoundPicks(fname string) []RoundPicks {
	var picks = make([]RoundPicks, 0)
	for line, err := range file.IterLines(fname) {
		if err != nil {
			panic(err)
		}
		parts := strings.Split(line, " ")
		p1, p2 := fstColMap[parts[0]], sndColMap[parts[1]]
		picks = append(picks, RoundPicks{P1: p1, P2: p2})
	}
	return picks
}

func GetTotalScore(picks []RoundPicks) int {
	total := 0
	for _, pick := range picks {
		_, p2Score := RunRound(pick.P1, pick.P2)
		total += p2Score
	}
	return total
}

// Part 2

type RoundOutcomes struct {
	P1     Shape
	Result Outcome
}

func ParseRoundOutcomes(fname string) []RoundOutcomes {
	var picks = make([]RoundOutcomes, 0)
	for line, err := range file.IterLines(fname) {
		if err != nil {
			panic(err)
		}
		parts := strings.Split(line, " ")
		p1, outcome := fstColMap[parts[0]], sndColMapOutcome[parts[1]]
		picks = append(picks, RoundOutcomes{P1: p1, Result: outcome})
	}
	return picks
}

func GetTotalScoreForOutcome(picks []RoundOutcomes) int {
	total := 0
	for _, pick := range picks {
		switch pick.Result {
		case Lose:
			// find which shape loses to P1, and calculate score.
			shape, found := slices.Find([]Shape{Rock, Paper, Scissors}, pick.P1.Beats)
			if !found {
				panic("didn't find any shape that beats P1!")
			}
			total += shape.Score() + loseScore
		case Draw:
			total += pick.P1.Score() + drawScore
		case Win:
			total += pick.P1.BeatBy().Score() + winScore
		}
	}
	return total
}
