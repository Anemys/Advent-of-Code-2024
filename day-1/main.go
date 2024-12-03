package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseInput(filepath string) ([]int, []int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var leftList, rightList []int

	for scanner.Scan() {
		line := scanner.Text()
		ids := strings.Fields(line)

		leftId, err := strconv.Atoi(ids[0])
		if err != nil {
			return nil, nil, err
		}
		rightId, err := strconv.Atoi(ids[1])
		if err != nil {
			return nil, nil, err
		}

		leftList = append(leftList, leftId)
		rightList = append(rightList, rightId)
	}

	return leftList, rightList, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func computeTotalDistance(leftList, rightList []int) int {
	slices.Sort(leftList)
	slices.Sort(rightList)

	total_distance := 0
	for i := 0; i < min(len(leftList), len(rightList)); i++ {
		total_distance += abs(leftList[i] - rightList[i])
	}

	return total_distance
}

func count(number int, list []int) int {
	count := 0
	for _, elm := range list {
		if elm == number {
			count++
		}
	}
	return count
}

func computeSimilarityScore(leftList, rightList []int) int {
	similarityScore := 0
	for _, id := range leftList {
		similarityScore += id * count(id, rightList)
	}

	return similarityScore
}

func main() {
	leftList, rightList, err := parseInput("input.txt")
	if err != nil {
		panic(err)
	}

	totalDistance := computeTotalDistance(leftList, rightList)

	fmt.Println("Total distance:", totalDistance)

	similarityScore := computeSimilarityScore(leftList, rightList)

	fmt.Println("Similarity score:", similarityScore)
}
