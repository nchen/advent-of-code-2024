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
	inputData, err := os.ReadFile("input-1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	inputStr := string(inputData)
	lines := strings.Split(inputStr, "\n")

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
			return
		}
		arr1 = append(arr1, num1)

		num2, err := strconv.Atoi(numbers[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		arr2 = append(arr2, num2)
	}
	sort.Ints(arr1)
	sort.Ints(arr2)

	result := 0
	for i := 0; i < len(arr1); i++ {
		diff := math.Abs(float64(arr1[i] - arr2[i]))
		result += int(diff)
	}

	fmt.Println("Total distance: ", result)

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
	fmt.Println("Total similarity: ", similarity)
}
