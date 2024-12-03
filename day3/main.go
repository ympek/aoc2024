package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func findClosestIndex(indexes [][]int, boundary int) int {
	closestIndex := -1
	for i := 0; i < len(indexes); i++ {
		if indexes[i][1] > boundary {
			return closestIndex
		}
		closestIndex = indexes[i][1]
	}
	return closestIndex
}

func main() {
	inputStr, _ := os.ReadFile("input")
	mulRegex, _ := regexp.Compile(`mul\([\d]+,[\d]+\)`)
	doRegex, _ := regexp.Compile(`do\(\)`)
	dontRegex, _ := regexp.Compile(`don't\(\)`)

	muls := mulRegex.FindAllIndex(inputStr, -1)
	dos := doRegex.FindAllIndex(inputStr, -1)
	donts := dontRegex.FindAllIndex(inputStr, -1)

	sumWithoutDonts := 0
	sumWithDonts := 0
	for _, indexes := range muls {
		mul := string(inputStr[indexes[0]:indexes[1]])
		numbers := strings.TrimPrefix(mul, "mul(")
		numbers = strings.TrimSuffix(numbers, ")")
		operands := strings.Split(numbers, ",")
		a, _ := strconv.Atoi(operands[0])
		b, _ := strconv.Atoi(operands[1])
		sumWithoutDonts += a * b

		closestDoIndex := findClosestIndex(dos, indexes[0])
		closestDontIndex := findClosestIndex(donts, indexes[0])
		if closestDontIndex == -1 || closestDoIndex > closestDontIndex {
			sumWithDonts += a * b
		}
	}

	fmt.Println("Answer to part 1:", sumWithoutDonts)
	fmt.Println("Answer to part 2:", sumWithDonts)
}
