package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Need the input file passed in, for example: 'go run main.go input'")
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	scanner := bufio.NewScanner(file)

	var matrix [][]rune

	for scanner.Scan() {
		text := scanner.Text()
		// Using runes here b/c they map to unicode so each character has a unique
		// number and this will work well to determine if we found our match
		matrix = append(matrix, []rune(text))
	}
	// TODO: Probably need to account for double counting. For example if we find
	// a vertical and horizontal match that intersects at the same letter
	// we would be double counting that letter
	// Maybe we need to serialize the coordinates of the matches we find
	// into a map and then count the number of unique coordinates at the end
	// to get the final match count
	// Or we could just use a 2d array of booleans that is the same size as the
	// input matrix and mark each coordinate we find as true. Then at the end
	// count the number of trues in that matrix
	matches := 0
	for y, yAxis := range matrix {
		for x, _ := range yAxis {
			// This would be a good to put in a goroutine
			north := findNorthPath(x, y)
			if len(north) > 0 && testPath(matrix, north) {
				// fmt.Printf("\nNorth Path match found %d is %v\n", y, north)
				matches++
			}

			east := findEastPath(matrix[y], x, y)
			if len(east) > 0 && testPath(matrix, east) {
				// fmt.Printf("\nEast Path match found %d is %v\n", y, east)
				matches++
			}

			south := findSouthPath(len(matrix)-1, x, y)
			if len(east) > 0 && testPath(matrix, south) {
				fmt.Printf("\nSouth Path match found %d is %v\n", y, south)
				matches++
			}
		}
	}
}

func testPath(matrix [][]rune, coords [][]int) bool {
	// runes for XMAS are 88 77 65 83
	// the total of these runes is 313
	// it also happens that there is no other combination of those
	// characters that can equal 313. Therefore if the total of the runes at
	// the coordinates is 313 we found a match (this is a directionless check)
	total := 0
	for _, coordinate := range coords {
		total += int(matrix[coordinate[0]][coordinate[1]] - 0)
	}
	if total == 313 { // magic number
		return true
	}
	return false
}

// Given an index I need to find all length of 4 that extends from that point
// So if we are given the index of X. This function should return a 2d array
// of indices that correspond to X and y
// . . . . . . . . . . .
// . . . y . . y . . y .
// . . . . y . y . y . .
// . . . . . y y y . . .
// . . . y y y X y y y .
// . . . . . y y y . . .
// . . . . y . y . y . .
// . . . y . . y . . y .
//
// If y should exceed the bounds of the input matrix ignore that path
func findPaths(matrix [][]int, x int, y int) [][]int {
	var result [][]int
	return result
}

// Given an index I need to find all length of 4 that extends from that point
// So if we are given the index of X. This function should return a 2d array
// of indices that correspond to X and y
// . . . . . . . . . . .
// . . . . . . y . . . .
// . . . . . . y . . . .
// . . . . . . y . . . .
// . . . . . . X . . . .
// . . . . . . . . . . .
// . . . . . . . . . . .
// . . . . . . . . . . .
//
// If y should exceed the bounds of the input matrix return an empty array
func findNorthPath(x int, y int) [][]int {
	var result [][]int
	end := y - 3
	// we found a valid north path
	if end >= 0 {
		return [][]int{{y, x}, {y - 1, x}, {y - 2, x}, {y - 3, x}}
	}
	return result
}

func findEastPath(row []rune, x int, y int) [][]int {
	var result [][]int
	end := x + 3
	if end <= (len(row) - 1) {
		return [][]int{{y, x}, {y, x + 1}, {y, x + 2}, {y, x + 3}}
	}
	return result
}

func findSouthPath(h int, x int, y int) [][]int {
	var result [][]int
	end := y + 3
	// we found a valid north path
	if end <= h {
		return [][]int{{y, x}, {y + 1, x}, {y + 2, x}, {y + 3, x}}
	}
	return result
}

func findWestPath(matrix [][]int, x int, y int) []int {
	var result []int
	return result
}

func findNorthEastPath(matrix [][]int, x int, y int) []int {
	var result []int
	return result
}

func findSouthEastPath(matrix [][]int, x int, y int) []int {
	var result []int
	return result
}

func findSouthWestPath(matrix [][]int, x int, y int) []int {
	var result []int
	return result
}

func findNorthWestPath(matrix [][]int, x int, y int) []int {
	var result []int
	return result
}
