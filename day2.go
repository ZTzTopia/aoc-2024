package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func isReportSafe(levels []string) bool {
	var lastDiff int
	for i := 1; i < len(levels); i++ {
		last, err := strconv.Atoi(levels[i-1])
		if err != nil {
			log.Fatal(err)
		}

		curr, err := strconv.Atoi(levels[i])
		if err != nil {
			log.Fatal(err)
		}

		diff := curr - last
		if diff == 0 {
			return false
		}

		if math.Abs(float64(diff)) > 3 {
			return false
		}

		if (diff < 0 && lastDiff > 0) || (diff > 0 && lastDiff < 0) {
			return false
		}

		lastDiff = diff
	}

	return true
}

func isReportSafeReduced(levels []string) bool {
	for i := 0; i < len(levels); i++ {
		// reduced := append(levels[:i], levels[i+1:]...) // Why we can't just do this? why we need to copy the slice? I'm Go newbie xd
		reduced := append([]string(nil), levels[:i]...)
		reduced = append(reduced, levels[i+1:]...)
		if isReportSafe(reduced) {
			return true
		}
	}

	return false
}

func (*Method) Day2Part1(input string) string {
	reports := strings.Split(input, "\n")

	totalSafeReport := 0
	for _, report := range reports {
		levels := strings.Split(report, " ")
		if isReportSafe(levels) {
			totalSafeReport++
		}
	}

	return strconv.Itoa(totalSafeReport)
}

func (*Method) Day2Part2(input string) string {
	reports := strings.Split(input, "\n")

	var totalSafeReport int
	for _, report := range reports {
		levels := strings.Split(report, " ")
		if isReportSafeReduced(levels) {
			totalSafeReport++
		}
	}

	return strconv.Itoa(totalSafeReport)
}
