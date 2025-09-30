package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Must pass in input file 'go run main input'")
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var reports [][]int

	for scanner.Scan() {
		text := scanner.Text()
		levels_s := strings.Fields(text)
		var levels []int
		for _, r := range levels_s {
			i, err := strconv.Atoi(r)
			if err != nil {
				log.Fatalf("Could not convert string to int: %v", err)
			}
			levels = append(levels, i)
		}
		reports = append(reports, levels)
	}

	safe_count := 0

	for _, report := range reports {
		var last int
		hasLast := false
		safe := true
		var direction int
		for _, level := range report {
			if hasLast == false {
				// we are on the first one set it and move on
				last = level
				hasLast = true
				continue
			}
			diff := level - last
			// if diff is 0 they are the same and unsafe
			if diff == 0 {
				safe = false
				break
			}
			// if the absolute diff is ever more than 2 we increased or decreased by too much
			if Abs(diff) > 3 {
				safe = false
				break
			}
			// fist comparison we need to initialize direction
			if direction == 0 {
				if diff > 0 {
					direction = 1
				} else {
					direction = -1
				}
			}
			if diff > 0 && direction == -1 {
				safe = false
				break
			}
			if diff < 0 && direction == 1 {
				safe = false
				break
			}
			last = level
		}
		if safe {
			safe_count++
		}
	}

	fmt.Printf("Safe reports: %v", safe_count)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
