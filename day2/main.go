package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isReportSafe(report []int) bool {
	if report[0] == report[1] {
		return false
	}
	decreasing := report[0] > report[1]

	for i := 0; i < len(report)-1; i++ {
		if decreasing {
			if report[i] < report[i+1] || report[i]-report[i+1] < 1 || report[i]-report[i+1] > 3 {
				return false
			}
		} else {
			if report[i] > report[i+1] || report[i+1]-report[i] < 1 || report[i+1]-report[i] > 3 {
				return false
			}
		}
	}
	return true
}

func isReportSafeAfterProblemDampener(report []int) bool {
	for i := 0; i < len(report); i++ {
		dampenedReport := append([]int{}, report[0:i]...)
		dampenedReport = append(dampenedReport, report[i+1:len(report)]...)
		if isReportSafe(dampenedReport) {
			return true
		}
	}
	return false
}

func main() {
	inputStr, _ := os.ReadFile("input")
	lines := strings.Split(string(inputStr), "\n")

	safeReports := 0
	safeReportsWithProblemDampener := 0
	for _, line := range lines {
		numbers := strings.Split(line, " ")
		if len(numbers) < 2 {
			continue
		}
		var report []int
		for _, x := range numbers {
			level, _ := strconv.Atoi(x)
			report = append(report, level)
		}

		if isReportSafe(report) {
			safeReports++
			safeReportsWithProblemDampener++
		} else {
			if isReportSafeAfterProblemDampener(report) {
				safeReportsWithProblemDampener++
			}
		}
	}

	fmt.Println("Answer to part 1:", safeReports)
	fmt.Println("Answer to part 2:", safeReportsWithProblemDampener)
}
