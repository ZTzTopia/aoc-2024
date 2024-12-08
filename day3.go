package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func getAllMultiplications(input string) int {
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	matches := re.FindAllStringSubmatch(input, -1)
	if matches == nil {
		log.Fatal("No matches found")
	}

	var sum int
	for _, match := range matches {
		x, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}

		y, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}

		sum += x * y
	}

	return sum
}

func (*PuzzleSolver) Day3Part1(input string) string {
	sum := getAllMultiplications(input)
	return strconv.Itoa(sum)
}

func (*PuzzleSolver) Day3Part2(input string) string {
	parts := strings.Split(input, "don't()")
	var doMultiplications string

	for i, part := range parts {
		if i == 0 {
			doMultiplications += part
			continue
		}

		parts := strings.Split(part, "do()")
		if len(parts) > 1 {
			doMultiplications += strings.Join(parts[1:], "")
		}
	}

	sum := getAllMultiplications(doMultiplications)
	return strconv.Itoa(sum)
}
