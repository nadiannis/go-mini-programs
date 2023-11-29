package main

import "fmt"

func main() {
	nums := []float64{1, 1, 2, 3, 3, 3, 4, 5, 5, 5, 6, 7}
	nums2 := []float64{10, 4, 2, 8, 24, 15, 9, 10, 0, 1, 2}
	nums3 := []float64{80, 23, 3, 1, 8, 7, 14, 55, 72, 98, 46, 41, 20}

	printResult(nums)
	printResult(nums2)
	printResult(nums3)
}

func printResult(nums []float64) {
	fmt.Println("Numbers:", nums)
	
	mean := mean(nums...)
	median := median(nums...)
	mode := mode(nums...)

	fmt.Println("\n=============== RESULT ===============")
	fmt.Println("Mean:", mean)
	fmt.Println("Median:", median)

	fmt.Print("Mode: ")
	for index, num := range mode {
		if (index == len(mode) - 1) {
			fmt.Println(num)
		} else {
			fmt.Print(num, ", ")
		}
	}
	fmt.Printf("\n\n")
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