package main

import (
	"fmt"
	"os"
	"strings"
)

func incrementIfXmasFound(counter *int, c1, c2, c3, c4 rune) {
	if c1 == 'X' && c2 == 'M' && c3 == 'A' && c4 == 'S' {
		*counter++
	} else if c1 == 'S' && c2 == 'A' && c3 == 'M' && c1 == 'X' {
		*counter++
	}
}

func incrementIfBigXFound(counter *int, center, tl, br, tr, bl rune) {
	if center != 'A' {
		return
	}
	if tl == 'S' && br == 'M' && tr == 'S' && bl == 'M' {
		*counter++
	}
	if tl == 'S' && br == 'M' && tr == 'M' && bl == 'S' {
		*counter++
	}
	if tl == 'M' && br == 'S' && tr == 'M' && bl == 'S' {
		*counter++
	}
	if tl == 'M' && br == 'S' && tr == 'S' && bl == 'M' {
		*counter++
	}
}

func main() {
	inputStr, _ := os.ReadFile("input")
	lines := strings.Split(string(inputStr), "\n")

	const PADDING = 6
	const INVALID = 'B'

	width := len(lines[0]) + PADDING
	height := len(lines) + PADDING

	grid := make([][]rune, height)
	for i := 0; i < height; i++ {
		grid[i] = make([]rune, width)
		for j := 0; j < width; j++ {
			grid[i][j] = INVALID
		}
	}

	for i, line := range lines {
		for j, ch := range line {
			grid[i+PADDING/2][j+PADDING/2] = ch
		}
	}

	puzzle1Count := 0
	puzzle2Count := 0
	for i := 3; i < height-3; i++ {
		for j := 3; j < width-3; j++ {
			// horizontal
			incrementIfXmasFound(&puzzle1Count, grid[i][j], grid[i][j+1], grid[i][j+2], grid[i][j+3])
			incrementIfXmasFound(&puzzle1Count, grid[i][j], grid[i][j-1], grid[i][j-2], grid[i][j-3])
			// vertical
			incrementIfXmasFound(&puzzle1Count, grid[i][j], grid[i+1][j], grid[i+2][j], grid[i+3][j])
			incrementIfXmasFound(&puzzle1Count, grid[i][j], grid[i-1][j], grid[i-2][j], grid[i-3][j])
			// diagonal
			incrementIfXmasFound(&puzzle1Count, grid[i][j], grid[i+1][j+1], grid[i+2][j+2], grid[i+3][j+3])
			incrementIfXmasFound(&puzzle1Count, grid[i][j], grid[i-1][j+1], grid[i-2][j+2], grid[i-3][j+3])
			incrementIfXmasFound(&puzzle1Count, grid[i][j], grid[i+1][j-1], grid[i+2][j-2], grid[i+3][j-3])
			incrementIfXmasFound(&puzzle1Count, grid[i][j], grid[i-1][j-1], grid[i-2][j-2], grid[i-3][j-3])

			// now for the second part, bit different algorithm
			incrementIfBigXFound(&puzzle2Count, grid[i][j], grid[i-1][j-1], grid[i+1][j+1], grid[i-1][j+1], grid[i+1][j-1])
		}
	}

	fmt.Println("Answer to part 1:", puzzle1Count)
	fmt.Println("Answer to part 2:", puzzle2Count)
}
