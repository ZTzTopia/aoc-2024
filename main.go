package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
    puzzles, lastDay := loadPuzzles()

	if os.Getenv("ASK_DAY_AND_PART") == "1" {
		var day int
		fmt.Print("Enter day: ")
		_, err := fmt.Scanf("%d", &day)
		if err != nil {
			log.Fatal(err)
		}

		var part int
		fmt.Print("Enter part: ")
		_, err = fmt.Scanf("%d", &part)
		if err != nil {
			log.Fatal(err)
		}
	}

	input, err := os.ReadFile(fmt.Sprintf("input/day%d.txt", lastDay))
    if err != nil {
        log.Fatal(err)
    }

	t := time.Now()
	p := findLastPuzzlePart(puzzles, lastDay)
	fmt.Println("Result:", p.Solve(string(input)))
	fmt.Println("Elapsed time:", time.Since(t))
}
