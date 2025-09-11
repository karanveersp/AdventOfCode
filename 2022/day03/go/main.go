package main

import (
	"errors"

	"github.com/karanveersp/gobelt/pkg/file"
)

// rucksacks with items
// one erroneous item
// every item type is identified by a single lowercase or uppercase letter.
// - all items of a given type must be in the same compartment, not in both
// - items in each compartment will be equal. Two halves on each line.
// - get the two sets, and do an intersection to determine which items appear in both.
// - a-z map to 1-26, and A-Z map to 27-52 as priorities. Get the priority of the intersect and sum.

// Part 2
// Elves are divided into groups of 3. Each elf carries a badge of a specific type.
// All 3 elves will have that type in their bag.

var PrioritiesMap = map[rune]int{}

func init() {
	// initialize the Priorities
	i := 1
	for c := 'a'; c <= 'z'; c++ {
		PrioritiesMap[c] = i
		i++
	}
	for c := 'A'; c <= 'Z'; c++ {
		PrioritiesMap[c] = i
		i++
	}
}

func FindIntersectionPriority(line string) (int, error) {
	runes := []rune(line)
	firstPart := runes[0 : len(runes)/2]
	secondPart := runes[len(runes)/2:]

	setFirst := map[rune]struct{}{}

	for _, c := range firstPart {
		setFirst[c] = struct{}{}
	}

	for _, c := range secondPart {
		_, ok := setFirst[c]
		if ok {
			return PrioritiesMap[c], nil
		}
	}
	return 0, errors.New("no intersection found")
}

func FindIntersectionSum(filePath string) (int, error) {
	sum := 0
	for line, err := range file.IterLines(filePath) {
		if err != nil {
			return 0, err
		}
		i, err := FindIntersectionPriority(line)
		if err != nil {
			return 0, err
		}
		sum += i
	}
	return sum, nil
}

func FindBadgesPrioritySum(filePath string) (int, error) {
	sum := 0
	var group []string
	for line, err := range file.IterLines(filePath) {
		if err != nil {
			return 0, err
		}

		if len(group) < 3 {
			// new group
			group = append(group, line)
			continue
		}
		intersection := FindIntersections(group...)
		sum += PrioritiesMap[rune(intersection[0])]

		// start new group
		group = []string{line}
	}
	intersection := FindIntersections(group...)
	sum += PrioritiesMap[rune(intersection[0])]
	return sum, nil
}

// FindIntersections will determine the common characters in the given strings
// and return them as a string.
func FindIntersections(s ...string) string {
	intersection := map[rune]struct{}{}
	for _, c := range s[0] {
		intersection[c] = struct{}{}
	}

	for _, line := range s[1:] {
		newIntersection := map[rune]struct{}{}
		for _, c := range line {
			// compare current string against previous intersection and add to new intersection.
			if _, ok := intersection[c]; ok {
				newIntersection[c] = struct{}{}
			}
		}
		// update intersection.
		intersection = newIntersection
	}
	res := ""
	for k, _ := range intersection {
		res += string(k)
	}
	return res
}
