package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Must pass input file 'go run main.go input")
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file %v", err)
	}

	scanner := bufio.NewScanner(file)

	whole_text := ""
	for scanner.Scan() {
		text := scanner.Text()
		whole_text += text
	}

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// Second argument here dictates _how_ many matches we return. -1 is for All
	matches := re.FindAllStringSubmatch(whole_text, -1)

	result := 0

	for _, e := range matches {
		l, err := strconv.Atoi(e[1])
		if err != nil {
			log.Fatalf("Failed to convert to int: %v", err)
		}
		r, err := strconv.Atoi(e[2])
		if err != nil {
			log.Fatalf("Failed to convert to int: %v", err)
		}
		result += (l * r)
	}

	fmt.Printf("Answer: %v\n", result)

}
