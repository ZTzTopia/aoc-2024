package main

import (
	"log"
	"strconv"
	"strings"
)

func day7ParseInput(input string) (target int, numbers []int) {
	parts := strings.Split(input, ": ")

	target, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}

	numberStrings := strings.Split(parts[1], " ")
	numbers = make([]int, len(numberStrings))

	for i, numberString := range numberStrings {
		numbers[i], err = strconv.Atoi(numberString)
		if err != nil {
			log.Fatal(err)
		}
	}

	return target, numbers
}

func day7IsValid(goal, temp int, numbers []int, allowConcat bool) bool {
	type State struct {
		temp  int
		index int
	}

	stack := []State{{temp, 0}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.index == len(numbers) {
			if current.temp == goal {
				return true
			}

			continue
		}

		currentNum := numbers[current.index]

		stack = append(stack, State{current.temp + currentNum, current.index + 1})
		stack = append(stack, State{current.temp * currentNum, current.index + 1})
		if allowConcat {
			concatTemp, _ := strconv.Atoi(strconv.Itoa(current.temp) + strconv.Itoa(currentNum))
			stack = append(stack, State{concatTemp, current.index + 1})
		}
	}

	return false
}

func day7Solve(input string, allowConcat bool) string {
	lines := strings.Split(input, "\n")
	totalSum := 0

	for _, line := range lines {
		target, numbers := day7ParseInput(line)
		if day7IsValid(target, numbers[0], numbers[1:], allowConcat) {
			totalSum += target
		}
	}

	return strconv.Itoa(totalSum)
}

func (*PuzzleSolver) Day7Part1(input string) string {
	return day7Solve(input, false)
}

func (*PuzzleSolver) Day7Part2(input string) string {
	return day7Solve(input, true)
}
