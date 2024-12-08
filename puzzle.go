package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

type PuzzleSolver struct {
}

type Puzzle struct {
	Day   int
	Part  int
	Solve func(input string) string
}

func findPuzzles(puzzles []Puzzle, day int) (part1, part2 *Puzzle) {
	for _, puzzle := range puzzles {
		if puzzle.Day == day {
			if puzzle.Part == 1 {
				part1 = &puzzle
			} else if puzzle.Part == 2 {
				part2 = &puzzle
			}
		}
	}
	return
}

func findPuzzlePart(puzzles []Puzzle, day, part int) *Puzzle {
	part1, part2 := findPuzzles(puzzles, day)
	if part == 1 {
		return part1
	} else if part == 2 {
		return part2
	}

	return nil
}

func findLastPuzzlePart(puzzles []Puzzle, day int) *Puzzle {
	part1, part2 := findPuzzles(puzzles, day)
	if part2 != nil {
		return part2
	}

	return part1
}

func extractPuzzleNumbers(input string) (int, int, error) {
	re := regexp.MustCompile(`Day(\d+)Part(\d+)`)
	matches := re.FindStringSubmatch(input)

	if matches == nil {
		return 0, 0, fmt.Errorf("no matches found in input: %s", input)
	}

	day, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, 0, err
	}

	part, err := strconv.Atoi(matches[2])
	if err != nil {
		return 0, 0, err
	}

	return day, part, nil
}

func loadPuzzles() ([]Puzzle, int) {
	var puzzles []Puzzle
	var lastDay int

	solver := &PuzzleSolver{}
	solverType := reflect.TypeOf(solver)

	for i := 0; i < solverType.NumMethod(); i++ {
		method := solverType.Method(i)

		day, part, err := extractPuzzleNumbers(method.Name)
		if err != nil {
			continue
		}

		puzzles = append(puzzles, Puzzle{
			Day:  day,
			Part: part,
			Solve: func(input string) string {
				args := []reflect.Value{
					reflect.ValueOf(solver),
					reflect.ValueOf(input),
				}
				results := method.Func.Call(args)
				return results[0].String()
			},
		})

		if day > lastDay {
			lastDay = day
		}
	}

	return puzzles, lastDay
}
