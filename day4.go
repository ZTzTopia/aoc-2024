package main

import (
	"strconv"
	"strings"
)

const Word = "XMAS"
const WordLength = len(Word)

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

func getGridElement(grids []string, x, y int) (rune, bool) {
	if x < 0 || x >= len(grids) || y < 0 || y >= len(grids[x]) {
		return 0, false
	}
	return rune(grids[x][y]), true
}

func countWordOccurrences(grids []string, x, y int) int {
	count := 0
	for _, direction := range directions {
		matched := true
		for i := 0; i < WordLength; i++ {
			nx, ny := x+i*direction.dx, y+i*direction.dy
			element, ok := getGridElement(grids, nx, ny)
			if !ok {
				matched = false
				break
			}

			if element != rune(Word[i]) {
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

func (*Method) Day4Part1(input string) string {
	grids := strings.Split(input, "\n")
	totalOccurrences := 0

	for x := 0; x < len(grids); x++ {
		for y := 0; y < len(grids[x]); y++ {
			totalOccurrences += countWordOccurrences(grids, x, y)
		}
	}

	return strconv.Itoa(totalOccurrences)
}

// Best code ever right?
func foo(grids []string, x, y, dir int) (bool, string) {
	nx, ny := x+1*directions[dir].dx, y+1*directions[dir].dy
	if nx < 0 || nx >= len(grids) {
		return false, ""
	}

	if ny < 0 || ny >= len(grids[nx]) {
		return false, ""
	}

	if grids[nx][ny] != 'S' && grids[nx][ny] != 'M' {
		return false, ""
	}

	return true, string(grids[nx][ny])
}

func isSpecialXMas(grids []string, x, y int) bool {
	if grids[x][y] != 'A' {
		return false
	}

	upLeft, upLeftS := foo(grids, x, y, 6)
	downRight, downRightS := foo(grids, x, y, 2)
	if upLeft && downRight && upLeftS != downRightS {
		upRight, upRightS := foo(grids, x, y, 7)
		downLeft, downLeftS := foo(grids, x, y, 3)
		if upRight && downLeft && upRightS != downLeftS {
			return true
		}
	}

	return false
}

func (*Method) Day4Part2(input string) string {
	grids := strings.Split(input, "\n")

	var totalOccurrences int
	for x := 0; x < len(grids); x++ {
		for y := 0; y < len(grids[x]); y++ {
			if isSpecialXMas(grids, x, y) {
				totalOccurrences++
			}
		}
	}

	return strconv.Itoa(totalOccurrences)
}
