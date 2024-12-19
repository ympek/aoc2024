package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Up = iota
	Right
	Down
	Left
)

type pos struct {
	x int
	y int
}

func changeDirection(dir int) int {
	return (dir + 1) % 4
}

func printMap(positions map[pos]bool) {
	result := ""
	for k, _ := range positions {
		result += fmt.Sprintf("(%d,%d);", k.x, k.y)
	}
	fmt.Println(result)
}

// returns if guard exited and after visiting how many squares
func solve(rows []string) (bool, int) {
	gridH := len(rows) - 1
	gridW := len(rows[0])

	direction := Up
	guardX := 0
	guardY := 0

	for i := 0; i < gridH; i++ {
		for j := 0; j < gridW; j++ {
			// first find the guard position, ^
			if rows[i][j] == '^' {
				guardX = j
				guardY = i
			}
		}
	}

	visitedPositions := make(map[pos]bool)
	nextX := 0
	nextY := 0

	numIterations := 0
	iterThreshold := 10000 // assume a loop after this threshold

	for {
		numIterations++
		visitedPositions[pos{x: guardX, y: guardY}] = true
		if numIterations > iterThreshold {
			return false, len(visitedPositions)
		}
		if direction == Up && guardY == 0 {
			break
		}
		if direction == Down && guardY == gridH-1 {
			break
		}
		if direction == Left && guardX == 0 {
			break
		}
		if direction == Right && guardX == gridW-1 {
			break
		}
		// move the guard
		if direction == Up {
			nextX = guardX
			nextY = guardY - 1
		}
		if direction == Down {
			nextX = guardX
			nextY = guardY + 1
		}
		if direction == Left {
			nextX = guardX - 1
			nextY = guardY
		}
		if direction == Right {
			nextX = guardX + 1
			nextY = guardY
		}

		cell := rows[nextY][nextX]

		if cell == '#' {
			direction = changeDirection(direction)
		} else {
			guardX = nextX
			guardY = nextY
		}
	}

	return true, len(visitedPositions)
}

func main() {
	contents, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(contents), "\n")

	_, ans1 := solve(rows)
	// simply brute force the answer to part 2 because it seemed good enough (and it is)
	loopsCreated := 0
	for i := 0; i < len(rows)-1; i++ {
		for j := 0; j < len(rows[0]); j++ {
			if rows[i][j] == '.' {
				bytes := []byte(rows[i])
				bytes[j] = '#'
				var sb strings.Builder
				_, _ = sb.Write(bytes)
				rows[i] = sb.String()
				exited, _ := solve(rows)
				if !exited {
					loopsCreated++
				}
				sb.Reset()
				bytes[j] = '.'
				_, _ = sb.Write(bytes)
				rows[i] = sb.String()
			}
		}
	}

	fmt.Println("Answer to part 1:", ans1)
	fmt.Println("Answer to part 2:", loopsCreated)
}
