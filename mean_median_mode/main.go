package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Println("\nEnter list of numbers separated by spaces (type q to quit)")
		fmt.Print(">>> ")
		scanner.Scan()
	
		input := scanner.Text()
		input = strings.TrimSpace(input)
	
		if input == "" {
			fmt.Println("Please type something")
			continue
		}
	
		if input == "q" {
			break
		}
	
		nums := strings.Split(input, " ")
		convertedNums, err := convertToFloats(nums)
		if err != nil {
			fmt.Println(err)
			continue
		}
	
		displayResult(convertedNums)
	}
}

func displayResult(nums []float64) {
	mean := mean(nums...)
	median := median(nums...)
	mode := mode(nums...)

	fmt.Println("* Mean:", mean)
	fmt.Println("* Median:", median)
	fmt.Print("* Mode: ")
	for index, num := range mode {
		if (index == len(mode) - 1) {
			fmt.Println(num)
		} else {
			fmt.Print(num, ", ")
		}
	}
}

func mean(nums ...float64) float64 {
	var sum float64

	for _, num := range nums {
		sum += float64(num)
	}
	
	return sum / float64(len(nums))
}

func median(nums ...float64) float64 {
	sort(nums)

	middle := int(len(nums) / 2)

	if len(nums) % 2 == 0 {
		return (nums[middle] + nums[middle - 1]) / 2
	} else {
		return nums[middle]
	}
}

func mode(nums ...float64) []float64 {
	counts := map[float64]int{}
	var maxCount int
	var result []float64

	for _, num := range nums {
		counts[float64(num)] += 1 
	}

	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}

	for num, count := range counts {
		if count == maxCount {
			result = append(result, num)
		}
	}

	return result
}

func sort(nums []float64) {
	unsortedUntilIndex := len(nums) - 1
	isSorted := false

	for !isSorted {
		isSorted = true

		for i := 0; i < unsortedUntilIndex; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				isSorted = false
			}
		}
		unsortedUntilIndex -= 1
	}
}

func convertToFloats(slice []string) ([]float64, error) {
	result := make([]float64, 0, len(slice))

	for _, item := range slice {
		convertedItem, err := strconv.ParseFloat(item, 64)
		if err != nil {
			return nil, fmt.Errorf("Unable to parse input, please type a number")
		}
		result = append(result, convertedItem)
	}

	return result, nil
}