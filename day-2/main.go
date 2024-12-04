package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Report []int

func stringToIntegerSlice(line []string) (Report, error) {
	var numbers []int
	for _, elm := range line {
		number, err := strconv.Atoi(elm)
		if err != nil {
			return numbers, err
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func parseInput(filepath string) ([]Report, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var reports []Report

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		report, err := stringToIntegerSlice(line)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}
	return reports, nil
}

func isMonotonic(report Report) bool {
	reversed := slices.Clone(report)
	slices.Reverse(reversed)
	return slices.IsSorted(report) || slices.IsSorted(reversed)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func hasGradualEvolution(report Report) bool {
	for i := 1; i < len(report); i++ {
		levelDiff := report[i] - report[i-1]
		if abs(levelDiff) < 1 || abs(levelDiff) > 3 {
			return false
		}
	}

	return true
}

func isSafeReport(report Report) bool {
	return isMonotonic(report) && hasGradualEvolution(report)
}

func countSafeReport(reports []Report, tolerant bool) int {
	safeReports := 0
	for _, report := range reports {
		if isSafeReport(report) {
			safeReports++
			continue
		}

		if !tolerant {
			continue
		}

		for i := range report {
			cuttedReport := slices.Clone(report)
			cuttedReport = append(cuttedReport[:i], cuttedReport[i+1:]...)

			if isSafeReport(cuttedReport) {
				safeReports++
				break
			}
		}
	}

	return safeReports
}

func main() {
	reports, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}

	safeReports := countSafeReport(reports, false)

	fmt.Println("Safe reports:", safeReports)

	toleratedReports := countSafeReport(reports, true)

	fmt.Println("Tolerated reports:", toleratedReports)
}
