package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <filename>")
	}
	result := 0

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var left_side []int
	var right_side []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		l_part, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalf("Error converting part: %v", err)
		}
		left_side = append(left_side, l_part)

		r_part, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalf("Error converting file: %v", err)
		}
		right_side = append(right_side, r_part)
	}

	// result = one_one(left_side, right_side)

	result = one_two(left_side, right_side)

	fmt.Printf("Answer %d\n", result)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
}

func one_two(left_side []int, right_side []int) int {
	result := make(chan int)
	var wg sync.WaitGroup

	for _, e := range left_side {
		wg.Add(1)
		go func(e int) {
			defer wg.Done()
			occurrences := countOccurrences(right_side, e)
			result <- (e * occurrences)
		}(e)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	total := 0
	for value := range result {
		total += value
	}

	return total
}

// Body of logic used to find the 1.1 exercise
func one_one(left_side []int, right_side []int) int {
	result := 0
	sort.Ints(left_side)
	sort.Ints(right_side)

	for i := range left_side {
		diff := math.Abs(float64(left_side[i] - right_side[i]))
		result += int(diff)
	}

	return result
}

// Helpers
func countOccurrences[T comparable](slice []T, target T) int {
	result := 0

	for _, e := range slice {
		if e == target {
			result++
		}
	}

	return result
}
