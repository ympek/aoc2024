package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getMiddleNumber(arr []int) int {
	return arr[(len(arr)-1)/2]
}

func isUpdateCorrect(mustComeBefore map[int][]int, mustComeAfter map[int][]int, update []int) bool {
	for iu := 0; iu < len(update); iu++ {
		page := update[iu]
		for k := iu + 1; k < len(update); k++ {
			if slices.Contains(mustComeBefore[page], update[k]) {
				return false
			}
		}
		for k := 0; k < iu; k++ {
			if slices.Contains(mustComeAfter[page], update[k]) {
				return false
			}
		}
	}
	return true
}

func correctUpdate(mustComeBefore map[int][]int, mustComeAfter map[int][]int, update []int) []int {
	for iu := 0; iu < len(update); iu++ {
		page := update[iu]
		for k := iu + 1; k < len(update); k++ {
			if slices.Contains(mustComeBefore[page], update[k]) {
				// incorrect! update[k] must come before page, move it
				result := append([]int{}, update[0:iu]...)
				result = append(result, update[k])
				result = append(result, update[iu:k]...)
				result = append(result, update[k+1:]...)
				return result
			}
		}

		for k := 0; k < iu; k++ {
			if slices.Contains(mustComeAfter[page], update[k]) {
				// incorrect! update[k] must come after page, move it
				result := append([]int{}, update[0:k]...)
				result = append(result, update[k+1:iu+1]...)
				result = append(result, update[iu+1:]...)
				return result
			}
		}
	}
	return update
}

func main() {
	inputStr, _ := os.ReadFile("input")
	lines := strings.Split(string(inputStr), "\n")

	mustComeBefore := make(map[int][]int)
	mustComeAfter := make(map[int][]int)

	updates := make(map[int][]int)
	// var correctUpdates []int
	iu := 0
	for _, line := range lines {
		if strings.Contains(line, "|") {
			// page ordering rule
			parts := strings.Split(line, "|")
			predecessor, _ := strconv.Atoi(parts[0])
			successor, _ := strconv.Atoi(parts[1])
			mustComeBefore[successor] = append(mustComeBefore[successor], predecessor)
			mustComeAfter[predecessor] = append(mustComeAfter[predecessor], successor)
		}

		if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			for _, part := range parts {
				number, _ := strconv.Atoi(part)
				updates[iu] = append(updates[iu], number)
			}
			iu++
		}
	}

	correctSum := 0
	correctedSum := 0
	for _, update := range updates {
		isCorrect := isUpdateCorrect(mustComeBefore, mustComeAfter, update)
		if isCorrect {
			correctSum += getMiddleNumber(update)
		} else {
			correctedUpdate := correctUpdate(mustComeBefore, mustComeAfter, update)
			i := 1
			for {
				if isUpdateCorrect(mustComeBefore, mustComeAfter, correctedUpdate) {
					break
				}
				correctedUpdate = correctUpdate(mustComeBefore, mustComeAfter, correctedUpdate)
				i++
			}
			correctedSum += getMiddleNumber(correctedUpdate)
		}
	}

	fmt.Println("Answer to part 1:", correctSum)
	fmt.Println("Answer to part 2:", correctedSum)
}
