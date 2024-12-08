package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	puzzles, lastDay := loadPuzzles()

	var puzzle *Puzzle
	if os.Getenv("ASK_DAY_AND_PART") == "1" {
		var day int
		var part int

		fmt.Print("Day: ")
		if _, err := fmt.Scan(&day); err != nil {
			log.Fatal(err)
		}

		fmt.Print("Part: ")
		if _, err := fmt.Scan(&part); err != nil {
			log.Fatal(err)
		}

		puzzle = findPuzzlePart(puzzles, day, part)
		if puzzle == nil {
			fmt.Printf("Puzzle for day %d part %d not found\n", day, part)
			fmt.Println("Available puzzles: ")
			for _, p := range puzzles {
				fmt.Printf("\t- %d.%d\n", p.Day, p.Part)
			}

			fmt.Println("Using the last puzzle instead")
			puzzle = findLastPuzzlePart(puzzles, lastDay)
		}

		fmt.Println()
	} else {
		puzzle = findLastPuzzlePart(puzzles, lastDay)
	}

	fmt.Printf("Day %d, Part %d\n", puzzle.Day, puzzle.Part)

	input, err := os.ReadFile(fmt.Sprintf("input/day%d.txt", puzzle.Day))
	if err != nil {
		log.Fatal(err)
	}

	t := time.Now()
	fmt.Println("Result:", puzzle.Solve(string(input)))
	fmt.Println("Elapsed time:", time.Since(t))
}
