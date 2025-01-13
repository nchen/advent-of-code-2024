package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, err := ReadLines("input-1.txt")
	if err != nil {
		fmt.Println("Read lines error: ", err)
		return
	}

	arr1, arr2, err := ParseTwoColumnNumbers(lines)
	if err != nil {
		fmt.Println("Parse numbers error: ", err)
		return
	}

	sort.Ints(arr1)
	sort.Ints(arr2)

	result := GetTotalDistance(arr1, arr2)
	fmt.Println("Total distance: ", result) // 1319616

	// Calculate similarity
	similarity := GetTotalSimilarity(arr1, arr2)
	fmt.Println("Total similarity: ", similarity) // 27267728
}

func ReadLines(fileName string) ([]string, error) {
	inputData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	inputStr := string(inputData)
	lines := strings.Split(inputStr, "\n")
	return lines, nil
}

func ParseTwoColumnNumbers(lines []string) ([]int, []int, error) {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)

	for _, line := range lines {
		if line == "" {
			continue
		}
		numbers := strings.Fields(line)
		num1, err := strconv.Atoi(numbers[0])
		if err != nil {
			fmt.Println(err)
			return nil, nil, err
		}
		arr1 = append(arr1, num1)

		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println(err)
			return nil, nil, err
		}
		arr2 = append(arr2, num2)
	}
	return arr1, arr2, nil
}

func GetTotalDistance(arr1 []int, arr2 []int) int {
	result := 0
	for i := 0; i < len(arr1); i++ {
		diff := math.Abs(float64(arr1[i] - arr2[i]))
		result += int(diff)
	}
	return result
}

func GetTotalSimilarity(arr1 []int, arr2 []int) int {
	// Calculate similarity
	arr2Counts := make(map[int]int)
	for _, num := range arr2 {
		arr2Counts[num] = arr2Counts[num] + 1
	}

	similarity := 0
	for _, num := range arr1 {
		if arr2Counts[num] > 0 {
			similarity += num * arr2Counts[num]
		}
	}
	return similarity
}
