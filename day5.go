package main

import (
	"sort"
	"strconv"
	"strings"
)

func processUpdates(input string, checkSorted bool) int {
	parts := strings.Split(input, "\n\n")

	rulesSections := strings.Split(parts[0], "\n")
	updateSections := strings.Split(parts[1], "\n")

	rules := make(map[string]bool)
	for _, rule := range rulesSections {
		rules[rule] = true
	}

	sum := 0
	for _, updateSection := range updateSections {
		updateParts := strings.Split(updateSection, ",")
		sort.Slice(updateParts, func(i, j int) bool {
			l := updateParts[i]
			r := updateParts[j]

			if _, exists := rules[l+"|"+r]; exists {
				return true
			}

			if _, exists := rules[r+"|"+l]; exists {
				return false
			}

			return false
		})

		sortedUpdateSection := strings.Join(updateParts, ",")
		if checkSorted && updateSection != sortedUpdateSection {
			continue
		}

		if !checkSorted && updateSection == sortedUpdateSection {
			continue
		}

		middleIndex := len(updateParts) / 2
		middleValue, err := strconv.Atoi(updateParts[middleIndex])
		if err == nil {
			sum += middleValue
		}
	}

	return sum
}

func (*PuzzleSolver) Day5Part1(input string) string {
	return strconv.Itoa(processUpdates(input, true))
}

func (*PuzzleSolver) Day5Part2(input string) string {
	return strconv.Itoa(processUpdates(input, false))
}
