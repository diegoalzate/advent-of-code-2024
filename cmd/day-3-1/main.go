package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/diegoalzate/advent-of-code-2024/internal"
)

type Operation struct {
	first  int
	second int
}

func (o *Operation) multiply() int {
	return o.first * o.second
}

// from "m" string onwards
func newOperation(reader *bufio.Reader) (Operation, error) {
	bytes, err := reader.ReadString(')')

	if err != nil {
		return Operation{}, err
	}

	if len(bytes) < 4 {
		return Operation{}, errors.New("not enough bytes to splice")
	}

	// expect mul(X,Y)
	mulWord := bytes[:4]

	if mulWord != "mul(" {
		return Operation{}, errors.New("wrong mul keyword")
	}

	closingBracket := bytes[len(bytes)-1]

	if closingBracket != ')' {
		return Operation{}, errors.New("missing closing bracket")
	}

	nums, err := readI(bytes)

	if err != nil {
		return Operation{}, err
	}

	return Operation{
		first:  nums[0],
		second: nums[1],
	}, nil
}

func readI(bytes string) ([2]int, error) {
	var out [2]int

	if len(bytes) < 4 {
		return out, errors.New("less than minimal amount of chars")
	}

	input := bytes[4 : len(bytes)-1]
	log.Print(input)

	numbers := strings.Split(input, ",")

	if len(numbers) != 2 {
		return out, errors.New("failed to get 2 numbers")
	}

	firstStr := numbers[0]
	secondStr := numbers[1]

	firstNum, err := strconv.Atoi(firstStr)

	if err != nil {
		return out, errors.New("failed to parse first number")
	}
	out[0] = firstNum

	secondNum, err := strconv.Atoi(secondStr)

	if err != nil {
		return out, errors.New("failed to parse second number")
	}
	out[1] = secondNum

	return out, nil
}

func parse(reader *bufio.Reader) []Operation {
	var out []Operation

	input, err := io.ReadAll(reader)

	if err != nil {
		log.Fatal(err)
	}

	for i, ch := range input {

		if ch != 'm' {
			continue
		}

		op, err := newOperation(
			bufio.NewReader(
				strings.NewReader(string(input[i:]))),
		)

		// stop clause
		if err == io.EOF {
			return out
		}

		if err != nil {
			// something else went wrong in the calc
			continue
		}

		out = append(out, op)

	}

	return out
}

func main() {
	reader, file := internal.Reader("input.txt")

	defer file.Close()

	operations := parse(reader)

	var sum int

	for _, op := range operations {
		sum += op.multiply()
	}

	log.Printf("Result: %v", sum)

}
