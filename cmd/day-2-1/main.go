package main

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/diegoalzate/advent-of-code-2024/internal"
)

type Sequence struct {
	values []int
}

func parse(lines []string) []Sequence {
	var out []Sequence

	for _, line := range lines {
		fields := strings.Fields(line)

		var values []int
		for _, str := range fields {
			num, err := strconv.Atoi(str)

			if err != nil {
				log.Fatal(err)
			}

			values = append(values, num)
		}

		out = append(out, Sequence{
			values: values,
		})
	}

	return out
}

func (s *Sequence) safe() bool {
	var expectedMultiplier int
	maxDiff := 3
	minDiff := 1

	for i := 0; i < len(s.values)-1; i++ {
		curr := s.values[i]
		next := s.values[i+1]

		diff := curr - next

		if math.Abs(float64(diff)) < float64(minDiff) || math.Abs(float64(diff)) > float64(maxDiff) {
			return false
		}

		var actualMultiplier int

		if diff < 0 {
			actualMultiplier = -1
		} else {
			actualMultiplier = 1
		}

		if expectedMultiplier != 0 && expectedMultiplier != actualMultiplier {
			return false
		}

		expectedMultiplier = actualMultiplier
	}

	return true
}

func main() {
	raw := internal.ReadFile("input.txt")
	sequences := parse(raw)

	var safeSequences int

	for _, seq := range sequences {
		if seq.safe() {
			safeSequences += 1
		}
	}

	log.Printf("Result: %v", safeSequences)
}