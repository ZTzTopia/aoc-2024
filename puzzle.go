package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

type Method struct {
}

type Puzzle struct {
	Day   int
	Part  int
	Solve func(input string) string
}

func findPuzzles(puzzles []Puzzle, day int) (p1, p2 *Puzzle) {
	for _, p := range puzzles {
		if p.Day == day {
			if p.Part == 1 {
				p1 = &p
			} else if p.Part == 2 {
				p2 = &p
			}
		}
	}
	return
}

func findLastPuzzlePart(puzzles []Puzzle, day int) *Puzzle {
	p1, p2 := findPuzzles(puzzles, day)
	if p2 != nil {
		return p2
	}
	return p1
}

func extractPuzzleNumbers(input string) (int, int, error) {
	re := regexp.MustCompile(`Day(\d+)Part(\d+)`)

	matches := re.FindStringSubmatch(input)
	if matches == nil {
		return 0, 0, fmt.Errorf("no matches found")
	}

	d, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0, 0, err
	}

	p, err := strconv.Atoi(matches[2])
	if err != nil {
		return 0, 0, err
	}

	return d, p, nil
}

func loadPuzzles() ([]Puzzle, int) {
	var puzzles []Puzzle
	var lastDay int

	pst := &Method{}
	pt := reflect.TypeOf(pst)

	for i := 0; i < pt.NumMethod(); i++ {
		m := pt.Method(i)
		d, p, err := extractPuzzleNumbers(m.Name)
		if err != nil {
			continue
		}

		puzzles = append(puzzles, Puzzle{
			Day:  d,
			Part: p,
			Solve: func(s string) string {
				args := []reflect.Value{
					reflect.ValueOf(pst),
					reflect.ValueOf(s),
				}
				res := m.Func.Call(args)[0]
				return res.String()
			},
		})

		if d > lastDay {
			lastDay = d
		}
	}

	return puzzles, lastDay
}
