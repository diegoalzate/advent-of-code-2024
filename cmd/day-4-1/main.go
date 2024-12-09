package main

import (
	"log"
	"strings"

	"github.com/diegoalzate/advent-of-code-2024/internal"
)

// looks like a recursive problem
// there is a matrix
// we would want to walk in all directions but not walk in a direction we have already checked
// we know when we found a path when it says XMAS
// we know we did not find a path when it does not start with X
// we want to know if we have already seen this path
// we want a counter of all of the paths
// does this have to be recursive? not really you are going to go in one direction 4 times,
// it is not a new decision every time

type Runner struct {
	matrix Matrix
}

func NewRunner(fileLoc string) Runner {
	m := newMatrix(fileLoc)

	return Runner{
		matrix: m,
	}
}

func (r Runner) walk() int {
	// go through all positions
	// check if in any dir it can get to XMAS
	// add to a counter if it does
	var count int
	for y, row := range r.matrix.chars {
		for x := range row {
			// this is the position you are evaluating
			pos := Position{
				x: x,
				y: y,
			}

			val := r.matrix.chars[pos.y][pos.x]

			if val != "X" {
				continue
			}

			for _, dir := range DIRECTIONS {
				// go places in that direction
				if !r.matrix.safe(pos, dir) {
					continue
				}

				secondPos := pos.next(dir)

				if !r.matrix.safe(secondPos, dir) {
					continue
				}

				thirdPos := secondPos.next(dir)

				if !r.matrix.safe(thirdPos, dir) {
					continue
				}

				fourthPos := thirdPos.next(dir)

				secondCh := r.matrix.chars[secondPos.y][secondPos.x]
				thirdCh := r.matrix.chars[thirdPos.y][thirdPos.x]
				fourthCh := r.matrix.chars[fourthPos.y][fourthPos.x]

				scanned := secondCh + thirdCh + fourthCh

				if scanned == "MAS" {
					count += 1
				}
			}

		}
	}

	return count
}

type Matrix struct {
	maxW  int
	maxY  int
	chars [][]string
}

func newMatrix(fileLoc string) Matrix {
	chars := [][]string{}

	lines := internal.ReadFile(fileLoc)

	for _, line := range lines {
		row := strings.Split(line, "")
		chars = append(chars, row)
	}

	return Matrix{
		maxW:  len(chars[0]) - 1,
		maxY:  len(chars) - 1,
		chars: chars,
	}
}

func (m Matrix) safe(p Position, dir Direction) bool {
	newPos := p.next(dir)

	// check if out of bounds
	if newPos.x > m.maxW || newPos.x < 0 {
		return false
	}

	if newPos.y > m.maxY || newPos.y < 0 {
		return false
	}

	return true
}

type Direction struct {
	x int
	y int
}

var (
	LEFT       = Direction{-1, 0}
	RIGHT      = Direction{1, 0}
	UP         = Direction{0, 1}
	DOWN       = Direction{0, -1}
	UP_RIGHT   = Direction{1, 1}
	UP_LEFT    = Direction{-1, 1}
	DOWN_RIGHT = Direction{1, -1}
	DOWN_LEFT  = Direction{-1, -1}
)

var DIRECTIONS = []Direction{
	LEFT,
	RIGHT,
	UP,
	DOWN,
	UP_RIGHT,
	DOWN_RIGHT,
	DOWN_LEFT,
	UP_LEFT,
}

type Position struct {
	x int
	y int
}

func (p Position) next(dir Direction) Position {
	newX := p.x + dir.x
	newY := p.y + dir.y

	return Position{
		x: newX,
		y: newY,
	}
}

func main() {
	runner := NewRunner("input.txt")

	count := runner.walk()

	log.Printf("Result:  %v", count)
}

// for pt 2
// wip, you could probably check if consecutive corners are the same and look for an A
