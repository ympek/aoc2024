package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	inputStr, _ := os.ReadFile("input")
	lines := strings.Split(string(inputStr), "\n")

	var leftList, rightList []int
	for _, line := range lines {
		numbers := strings.Split(line, "   ")
		if len(numbers) != 2 {
			continue
		}
		left, _ := strconv.Atoi(numbers[0])
		leftList = append(leftList, left)
		right, _ := strconv.Atoi(numbers[1])
		rightList = append(rightList, right)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	var sumOfDistances float64
	var similarityScore int
	for i := 0; i < len(leftList); i++ {
		sumOfDistances += math.Abs(float64(leftList[i] - rightList[i]))

		// could be made much more efficient given we have sorted lists
		occurences := 0
		for j := 0; j < len(rightList); j++ {
			if leftList[i] == rightList[j] {
				occurences++
			}
		}
		similarityScore += leftList[i] * occurences
	}

	fmt.Println("Answer to part 1:", int(sumOfDistances))
	fmt.Println("Answer to part 2:", similarityScore)
}
