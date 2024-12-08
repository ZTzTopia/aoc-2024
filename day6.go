package main

import (
	"strconv"
	"strings"
)

type Direction uint8

const (
	U Direction = iota
	R
	D
	L
)

type Point struct{ x, y int }

func (p Point) Next(direction Direction) Point {
	switch direction {
	case U:
		return Point{p.x - 1, p.y}
	case R:
		return Point{p.x, p.y + 1}
	case D:
		return Point{p.x + 1, p.y}
	case L:
		return Point{p.x, p.y - 1}
	default:
		return p
	}
}

func getCurrentPosition(grids []string) Point {
	for x := 0; x < len(grids); x++ {
		for y := 0; y < len(grids[x]); y++ {
			if getGridElement(grids, x, y) == rune('^') {
				return Point{x, y}
			}
		}
	}

	return Point{}
}

func (*PuzzleSolver) Day6Part1(input string) string {
	grids := strings.Split(input, "\n")

	point := getCurrentPosition(grids)
	direction := U
	visited := map[Point]struct{}{
		point: {},
	}

	for {
		nextPoint := point.Next(direction)
		if nextPoint.x < 0 || nextPoint.x >= len(grids) || nextPoint.y < 0 || nextPoint.y >= len(grids[nextPoint.x]) {
			break
		}

		if getGridElement(grids, nextPoint.x, nextPoint.y) == rune('#') {
			direction = (direction + 1) % 4
		} else {
			point = nextPoint
			visited[point] = struct{}{}
		}
	}

	return strconv.Itoa(len(visited))
}
