package main

import (
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/diegoalzate/advent-of-code-2024/internal"
)

type Location struct {
	values []int
}

func parse(lines []string) []Location {
	first := Location{}
	second := Location{}

	for _, line := range lines {

		fields := strings.Fields(line)

		firstVal, err := strconv.Atoi(fields[0])

		if err != nil {
			log.Fatal(err)
		}

		secondVal, err := strconv.Atoi(fields[1])

		if err != nil {
			log.Fatal(err)
		}

		first.values = append(first.values, firstVal)
		second.values = append(second.values, secondVal)
	}

	return []Location{first, second}
}

func (l *Location) sortValues() {
	slices.Sort(l.values)
}

func sumDiff(first Location, second Location) int {
	absoluteDiffArr := make([]int, len(first.values))

	first.sortValues()
	second.sortValues()

	for i := range len(first.values) {
		diff := first.values[i] - second.values[i]
		absoluteDiffArr[i] = int(math.Abs(float64(diff)))
	}

	var sum int

	for _, delta := range absoluteDiffArr {
		sum += delta
	}

	return sum
}

func main() {
	raw := internal.ReadFile("input.txt")

	locations := parse(raw)

	log.Printf("Result: %v", sumDiff(locations[0], locations[1]))
}
