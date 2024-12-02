package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/diegoalzate/advent-of-code-2024/internal"
)

type Location struct {
	values []int
}

type CountMap struct {
	values map[int]int
}

func (l *Location) CountMap() CountMap {
	simMap := CountMap{
		values: make(map[int]int),
	}

	for _, val := range l.values {
		simMap.values[val] += 1
	}

	return simMap
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

func similaritySum(first Location, second Location) int {
	var sum int

	countMap := second.CountMap()

	for _, val := range first.values {
		sum += val * countMap.values[val]
	}

	return sum
}

func main() {
	raw := internal.ReadFile("input.txt")

	locations := parse(raw)

	log.Printf("Result: %v", similaritySum(locations[0], locations[1]))
}
