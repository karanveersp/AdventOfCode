package main

import (
	"fmt"
	"iter"
	"slices"
	"strconv"
	"time"

	futil "github.com/karanveersp/gobelt/pkg/file"
)

func main() {
	for n, err := range futil.IterLines("test_input.txt") {
		if err != nil {
			panic(err)
		}
		fmt.Println(n)
		time.Sleep(1 * time.Second)
	}
}

func InfiniteNumbers() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; ; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func Parse(filePath string) []int {
	var count = 0
	carriers := make([]int, 0)
	for line, err := range futil.IterLines(filePath) {
		if err != nil {
			panic(err)
		}
		i, err := strconv.Atoi(line)
		if err != nil {
			// value is empty line
			carriers = append(carriers, count)
			count = 0
			continue
		}
		count += i
	}
	carriers = append(carriers, count) // add the last carrier count
	return carriers
}

func GetMaxCalories(inventory []int) int {
	return slices.Max(inventory)
}

func GetTopThreeCarriersTotal(inventory []int) int {
	slices.Sort(inventory)
	size := len(inventory)
	return inventory[size-1] + inventory[size-2] + inventory[size-3]
}
