package main

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/diegoalzate/advent-of-code-2024/internal"
)

// rules
type Rule struct {
	before []int
	after  []int
}

type Rules = map[int]Rule

func NewRules(lines []string) Rules {
	out := make(Rules)

	for _, line := range lines {
		// for each line add number to before and after
		nums := strings.Split(line, "|")

		if len(nums) < 2 {
			log.Fatal("should have a number on either side of the rule")
		}

		before, err := strconv.Atoi(nums[0])

		if err != nil {
			log.Fatal(err)
		}

		after, err := strconv.Atoi(nums[1])

		if err != nil {
			log.Fatal(err)
		}

		out[before] = Rule{
			before: out[before].before,
			after:  append(out[before].after, after),
		}

		out[after] = Rule{
			before: append(out[after].before, before),
			after:  out[after].after,
		}
	}

	return out
}

// a lot of rules on what should be before or after a number
// 79|18
// 18|95
// number -> [before: []int, after[]int]

// pages
// a map of the number we are looking for and the position

// 72 -> 2
// 29 -> 3
// 34 -> 4

type PagesMap map[int]int

func NewPagesMap(pages Pages) PagesMap {
	out := make(PagesMap)

	for i, num := range pages {
		out[num] = i
	}

	return out
}

type Pages []int

func (p Pages) valid(pagesMap PagesMap, rules Rules) bool {
	valid := true

	for i, num := range p {
		numRules := rules[num]

		// before
		for _, beforeNum := range numRules.before {
			beforeIdx, ok := pagesMap[beforeNum]

			if !ok {
				continue
			}

			if beforeIdx > i {
				log.Printf("failed: expected num %v to be before %v, but indexes look like %v", num, beforeNum, pagesMap)
				return false
			}
		}

		// after
		for _, afterNum := range numRules.after {
			afterIdx, ok := pagesMap[afterNum]

			if !ok {
				continue
			}

			if afterIdx < i {
				log.Printf("failed: expected num %v to be after %v, but indexes look like %v", num, afterNum, pagesMap)
				return false
			}
		}
	}

	return valid
}

func (p Pages) middle() int {
	probableMiddleIdx := math.Floor(float64(len(p)) / float64(2))

	return p[int(probableMiddleIdx)]
}

func NewPages(line string) Pages {
	var out Pages

	chars := strings.Split(line, ",")

	for _, ch := range chars {
		val, err := strconv.Atoi(ch)

		if err != nil {
			log.Fatal(err)
		}

		out = append(out, val)
	}

	return out
}

func parse(inputLoc string) (ruleStr []string, pagesStr []string) {
	var blankLineIdx int

	lines := internal.ReadFile(inputLoc)

	// find blank line
	for i, line := range lines {
		if line == "" {
			blankLineIdx = i
			break
		}
	}

	return lines[:blankLineIdx], lines[blankLineIdx+1:]
}

// get the sum of the middle
// parse
// go through each number and check if they are in front or behind

func main() {
	rulesStr, pagesStr := parse("input.txt")

	rules := NewRules(rulesStr)

	log.Print(rules)

	var result int

	for _, pageStr := range pagesStr {
		pages := NewPages(pageStr)
		log.Print(pages)
		pagesMap := NewPagesMap(pages)

		if pages.valid(pagesMap, rules) {
			log.Print("valid")
			result += pages.middle()
		}
	}

	log.Printf("Result: %v", result)
}
