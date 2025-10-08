package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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
	matches := make(map[string][][]int)

	for y, yAxis := range matrix {
		for x, _ := range yAxis {
			// This would be a good to put in a goroutine
			north := findNorthPath(x, y)
			if len(north) > 0 && testPath(matrix, north) {
				matches = mergeMap(matches, north)
				// fmt.Printf("\nmerged map: %v\n", matches)
				// fmt.Printf("\nMap length after north %d\n", len(matches))
			}

			east := findEastPath(len(matrix[y]), x, y)
			if len(east) > 0 && testPath(matrix, east) {
				matches = mergeMap(matches, east)
				// fmt.Printf("\nmerged map: %v\n", matches)
				// fmt.Printf("\nMap length after east %d\n", len(matches))
			}

			south := findSouthPath(len(matrix)-1, x, y)
			if len(south) > 0 && testPath(matrix, south) {
				matches = mergeMap(matches, south)
			}

			west := findWestPath(x, y)
			if len(west) > 0 && testPath(matrix, west) {
				matches = mergeMap(matches, west)
			}

			northEast := findNorthEastPath(matrix, x, y)
			if len(northEast) > 0 && testPath(matrix, northEast) {
				matches = mergeMap(matches, northEast)
			}

			southEast := findSouthEastPath(matrix, x, y)
			if len(southEast) > 0 && testPath(matrix, southEast) {
				matches = mergeMap(matches, southEast)
			}

			southWest := findSouthWestPath(matrix, x, y)
			if len(southWest) > 0 && testPath(matrix, southWest) {
				matches = mergeMap(matches, southWest)
			}

			northWest := findNorthWestPath(matrix, x, y)
			if len(northWest) > 0 && testPath(matrix, northWest) {
				matches = mergeMap(matches, northWest)
			}
		}
	}
	fmt.Printf("\nFound matches: %v\n", len(matches))
}

func mergeMap(destination map[string][][]int, source map[string][][]int) map[string][][]int {
	for k, v := range source {
		if len(destination[k]) == 0 {
			destination[k] = v
		}
	}
	return destination
}

func testPath(matrix [][]rune, coords map[string][][]int) bool {

	key := []rune{'X', 'M', 'A', 'S'}
	keyCount := 0
	for _, coordinates := range coords {
		for _, coordinate := range coordinates {
			if matrix[coordinate[0]][coordinate[1]] != key[keyCount] {
				return false
			}
			keyCount++
		}
	}
	return true
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

func createMapKey(coords [][]int) string {
	// Make a copy of coords so this function remains 'pure'
	// otherwise the sort.Slice would modify the input unknowing to the
	// caller. Really they just want a key not their input to be quietly sorted
	cp := make([][]int, len(coords))
	for i := range coords {
		cp[i] = []int{coords[i][0], coords[i][1]}
	}

	// Using a custom sorting function to sort by pairs
	// This function asks the question "should 'i' come before 'j'"
	// if it returns 'true' then the answer is yes.
	sort.Slice(cp, func(i, j int) bool {
		// First we compare the 'x' coordinates
		if cp[i][0] != cp[j][0] {
			return cp[i][0] < cp[j][0]
		}
		// if the x coordinates are the same we compare the y
		return cp[i][1] < cp[j][1]
	})

	// build a compact, unambiguous key: "x,y;x,y;..."
	var b bytes.Buffer
	for _, c := range cp {
		b.WriteString(strconv.Itoa(c[0]))
		b.WriteByte(',')
		b.WriteString(strconv.Itoa(c[1]))
		b.WriteByte(';')
	}
	return b.String()
}

// Given an index I need to find the coordinates of the y's
// . . . . . . . . . . .
// . . . . . . y . . . .
// . . . . . . y . . . .
// . . . . . . y . . . .
// . . . . . . X . . . .
// . . . . . . . . . . .
// . . . . . . . . . . .
// . . . . . . . . . . .
//
// # If y should exceed the bounds of the input matrix return an empty array
//
// returns a map keyed on the sorted
func findNorthPath(x int, y int) map[string][][]int {
	var coordinates [][]int
	end := y - 3
	// we found a valid north path
	result := make(map[string][][]int)
	if end >= 0 {
		coordinates = [][]int{{y, x}, {y - 1, x}, {y - 2, x}, {y - 3, x}}
		result[createMapKey(coordinates)] = coordinates
	}
	return result
}

func findEastPath(length int, x int, y int) map[string][][]int {
	var coordinates [][]int
	end := x + 3
	result := make(map[string][][]int)
	if end <= length-1 {
		coordinates = [][]int{{y, x}, {y, x + 1}, {y, x + 2}, {y, x + 3}}
		result[createMapKey(coordinates)] = coordinates
	}
	return result
}

func findSouthPath(height int, x int, y int) map[string][][]int {
	var coordinates [][]int
	end := y + 3
	result := make(map[string][][]int)
	// we found a valid north path
	if end <= height {
		coordinates = [][]int{{y, x}, {y + 1, x}, {y + 2, x}, {y + 3, x}}
		result[createMapKey(coordinates)] = coordinates
	}
	return result
}

// Given an index I need to find the coordinates of the y's
// . . . . . . . . . . .
// . . . . . . . . . . .
// . . . . . . . . . . .
// . . . . . . . . . . .
// . . . y y y X . . . .
// . . . . . . . . . . .
// . . . . . . . . . . .
// . . . . . . . . . . .
//
// # If y should exceed the bounds of the input matrix return an empty array
//
// returns a map keyed on the sorted
func findWestPath(x int, y int) map[string][][]int {
	var coordinates [][]int
	end := x - 3
	result := make(map[string][][]int)
	if end >= 0 {
		coordinates = [][]int{{y, x}, {y, x - 1}, {y, x - 2}, {y, x - 3}}
		result[createMapKey(coordinates)] = coordinates
	}
	return result
}

// Given an index I need to find the coordinates of the y's
// . . . . . . . . . . .
// . . . . . . . . . y .
// . . . . . . . . y . .
// . . . . . . . y . . .
// . . . . . . X . . . .
// . . . . . . . . . . .
// . . . . . . . . . . .
// . . . . . . . . . . .
//
// # If y should exceed the bounds of the input matrix return an empty array
//
// returns a map keyed on the sorted coordinates
func findNorthEastPath(matrix [][]rune, x int, y int) map[string][][]int {
	var coordinates [][]int
	yEnd := y - 3
	xEnd := x + 3

	result := make(map[string][][]int)
	if yEnd >= 0 && xEnd <= len(matrix[y])-1 {
		// Found a path
		coordinates = [][]int{{y, x}, {y - 1, x + 1}, {y - 2, x + 2}, {y - 3, x + 3}}
		result[createMapKey(coordinates)] = coordinates
	}
	return result
}

// Given an index I need to find the coordinates of the y's
// . . . . . . . . . . .
// . . . . . . . . . . .
// . . . . . . . . . . .
// . . . . . . . . . . .
// . . . . . . X . . . .
// . . . . . . . y . . .
// . . . . . . . . y . .
// . . . . . . . . . y .
//
// # If y should exceed the bounds of the input matrix return an empty array
//
// returns a map keyed on the sorted coordinates
func findSouthEastPath(matrix [][]rune, x int, y int) map[string][][]int {
	var coordinates [][]int
	yEnd := y + 3
	xEnd := x + 3
	result := make(map[string][][]int)

	if yEnd <= len(matrix)-1 && xEnd <= len(matrix[y])-1 {
		coordinates = [][]int{{y, x}, {y + 1, x + 1}, {y + 2, x + 2}, {y + 3, x + 3}}
		result[createMapKey(coordinates)] = coordinates
	}
	return result
}

// Given an index I need to find the coordinates of the y's
// . . . . . . . . . . .
// . . . . . . . . . . .
// . . . . . . . . . . .
// . . . . . . . . . . .
// . . . . . . X . . . .
// . . . . . y . . . . .
// . . . . y . . . . . .
// . . . y . . . . . . .
//
// # If y should exceed the bounds of the input matrix return an empty array
//
// returns a map keyed on the sorted coordinates
func findSouthWestPath(matrix [][]rune, x int, y int) map[string][][]int {
	var coordinates [][]int
	yEnd := y + 3
	xEnd := x - 3
	result := make(map[string][][]int)
	if yEnd <= len(matrix)-1 && xEnd >= 0 {
		coordinates = [][]int{{y, x}, {y + 1, x - 1}, {y + 2, x - 2}, {y + 3, x - 3}}
		result[createMapKey(coordinates)] = coordinates
	}
	return result
}

// Given an index I need to find the coordinates of the y's
// . . . . . . . . . . .
// . . . y . . . . . . .
// . . . . y . . . . . .
// . . . . . y . . . . .
// . . . . . . X . . . .
// . . . . . . . . . . .
// . . . . . . . . . . .
// . . . . . . . . . . .
//
// # If y should exceed the bounds of the input matrix return an empty array
//
// returns a map keyed on the sorted coordinates
func findNorthWestPath(matrix [][]rune, x int, y int) map[string][][]int {
	var coordinates [][]int
	yEnd := y - 3
	xEnd := x - 3
	result := make(map[string][][]int)

	if yEnd >= 0 && xEnd >= 0 {
		coordinates = [][]int{{y, x}, {y - 1, x - 1}, {y - 2, x - 2}, {y - 3, x - 3}}
		result[createMapKey(coordinates)] = coordinates
	}
	return result
}
