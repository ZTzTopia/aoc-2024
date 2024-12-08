package main

import (
	"strconv"
	"strings"
)

var word = []rune{'X', 'M', 'A', 'S'}
var directions = []struct{ dx, dy int }{
	{0, 1},   // right
	{1, 0},   // down
	{1, 1},   // down-right
	{1, -1},  // down-left
	{0, -1},  // left
	{-1, 0},  // up
	{-1, -1}, // up-left
	{-1, 1},  // up-right
}
var diagonalDirections = []struct{ dx, dy int }{
	{1, 1},   // down-right
	{1, -1},  // down-left
	{-1, -1}, // up-left
	{-1, 1},  // up-right
}

func getGridElement(grids []string, x, y int) rune {
	if x < 0 || x >= len(grids) || y < 0 || y >= len(grids[x]) {
		return 0
	}
	return rune(grids[x][y])
}

func countWordOccurrences(grids []string, x, y int) int {
	count := 0
	for _, direction := range directions {
		matched := true
		for i := 0; i < len(word); i++ {
			nx, ny := x+i*direction.dx, y+i*direction.dy
			if getGridElement(grids, nx, ny) != rune(word[i]) {
				matched = false
				break
			}
		}

		if matched {
			count++
		}
	}

	return count
}

func getDiagonalElementFromA(grids []string, x, y, dir int) rune {
	nx, ny := x+1*diagonalDirections[dir].dx, y+1*diagonalDirections[dir].dy
	switch getGridElement(grids, nx, ny) {
	case rune('M'):
		return 1
	case rune('S'):
		return 2
	default:
		return -1
	}
}

func isSpecialXMas(grids []string, x, y int) bool {
	if getGridElement(grids, x, y) != rune('A') {
		return false
	}

	upLeft := getDiagonalElementFromA(grids, x, y, 0)
	downRight := getDiagonalElementFromA(grids, x, y, 2)
	if upLeft == -1 || downRight == -1 || upLeft == downRight {
		return false
	}

	upRight := getDiagonalElementFromA(grids, x, y, 1)
	downLeft := getDiagonalElementFromA(grids, x, y, 3)
	if upRight == -1 || downLeft == -1 || upRight == downLeft {
		return false
	}

	return true
}

func (*PuzzleSolver) Day4Part1(input string) string {
	grids := strings.Split(input, "\n")
	totalOccurrences := 0

	for x := 0; x < len(grids); x++ {
		for y := 0; y < len(grids[x]); y++ {
			totalOccurrences += countWordOccurrences(grids, x, y)
		}
	}

	return strconv.Itoa(totalOccurrences)
}

func (*PuzzleSolver) Day4Part2(input string) string {
	grids := strings.Split(input, "\n")

	totalOccurrences := 0
	for x := 0; x < len(grids); x++ {
		for y := 0; y < len(grids[x]); y++ {
			if isSpecialXMas(grids, x, y) {
				totalOccurrences++
			}
		}
	}

	return strconv.Itoa(totalOccurrences)
}
