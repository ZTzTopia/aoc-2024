package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func readLeftRight(input string) ([]int, []int) {
	var l []int
	var r []int

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var x, y int
		_, err := fmt.Sscanf(line, "%d %d", &x, &y)
		if err != nil {
			log.Fatal(err)
		}

		l = append(l, x)
		r = append(r, y)
	}

	return l, r
}

func (*Method) Day1Part1(input string) string {
	l, r := readLeftRight(input)

	sort.Ints(l)
	sort.Ints(r)

	sum := 0
	for i := 0; i < len(l); i++ {
		if r[i]-l[i] > 0 {
			sum += r[i] - l[i]
		} else {
			sum += l[i] - r[i]
		}
	}

	return strconv.Itoa(sum)
}

func (*Method) Day1Part2(input string) string {
	l, r := readLeftRight(input)

	sum := 0
	for i := 0; i < len(l); i++ {
		count := 0
		for j := 0; j < len(r); j++ {
			if l[i] == r[j] {
				count += 1
			}
		}

		sum += l[i] * count
	}

	return strconv.Itoa(sum)
}
