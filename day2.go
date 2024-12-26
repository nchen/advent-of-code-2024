package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputData, err := os.ReadFile("input-2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	inputStr := string(inputData)
	lines := strings.Split(inputStr, "\n")

	safeCount := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		numbers := strings.Fields(line)

		sign := 0
		isFirst := true
		isSafe := true

		for i := 1; i < len(numbers); i++ {
			num, err := strconv.Atoi(numbers[i])
			if err != nil {
				fmt.Println(err)
				return
			}
			prev, err := strconv.Atoi(numbers[i-1])
			if err != nil {
				fmt.Println(err)
				return
			}
			diff := num - prev
			absDiff := int(math.Abs(float64(diff)))
			if absDiff < 1 || absDiff > 3 {
				isSafe = false
				break
			}
			if isFirst {
				if diff > 0 {
					sign = 1
				} else {
					sign = -1
				}
				isFirst = false
			} else {
				if sign*diff < 0 {
					isSafe = false
					break
				}
			}
		}
		if isSafe {
			safeCount += 1
			// fmt.Println(line)
		}
	}
	fmt.Println("Safe count:", safeCount)
}
