package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const mainMenuText = `
=========================== Generate Numbers ===========================
1 = Even | 2 = Odd | 3 = Square | 4 = Cube | 5 = Fibonacci | 6 = Prime
`

type parameter struct {
	min, max int
}

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	params := new(parameter)
	sequenceTypes := getSequenceTypeOptions()

	for true {
		input = getInput(scanner, fmt.Sprintf("%s\nChoose the type (type q to quit): ", mainMenuText))

		if input == "" {
			fmt.Println("\nPlease type something")
			continue
		}

		if input == "q" {
			break
		}

		switch sequenceTypes[input] {
		case "even":
			promptMinMax(scanner, params, sequenceTypes[input], generateEvenNumbers)
		case "odd":
			promptMinMax(scanner, params, sequenceTypes[input], generateOddNumbers)
		case "square":
			promptMinMax(scanner, params, sequenceTypes[input], generateSquareNumbers)
		case "cube":
			promptMinMax(scanner, params, sequenceTypes[input], generateCubeNumbers)
		case "fibonacci":
			promptMinMax(scanner, params, sequenceTypes[input], generateFibonacciNumbers)
		case "prime":
			promptMinMax(scanner, params, sequenceTypes[input], generatePrimeNumbers)
		default:
			fmt.Println("\nInvalid sequence type. You can only type 1, 2, 3, 4, 5, or 6.")
			continue
		}
	}
}

func generateEvenNumbers(params *parameter) []int {
	var result []int

	if params.max < params.min {
		params.max = params.min
	}

	for num := params.min; num <= params.max; num++ {
		if num%2 == 0 {
			result = append(result, num)
		}
	}

	return result
}

func generateOddNumbers(params *parameter) []int {
	var result []int

	if params.max < params.min {
		params.max = params.min
	}

	for num := params.min; num <= params.max; num++ {
		if num%2 != 0 {
			result = append(result, num)
		}
	}

	return result
}

func generateSquareNumbers(params *parameter) []int {
	var result []int

	if params.min < 0 {
		params.min = 0
	}

	if params.max < params.min {
		params.max = params.min
	}

	minSquareRoot := math.Sqrt(float64(params.min))
	maxSquareRoot := math.Sqrt(float64(params.max))
	min := int(math.Ceil(minSquareRoot))
	max := int(math.Floor(maxSquareRoot))

	for num := min; num <= max; num++ {
		result = append(result, num*num)
	}

	return result
}

func generateCubeNumbers(params *parameter) []int {
	var result []int

	if params.max < params.min {
		params.max = params.min
	}

	minCubeRoot := math.Cbrt(float64(params.min))
	maxCubeRoot := math.Cbrt(float64(params.max))
	min := int(math.Ceil(minCubeRoot))
	max := int(math.Floor(maxCubeRoot))

	for num := min; num <= max; num++ {
		result = append(result, num*num*num)
	}

	return result
}

func generateFibonacciNumbers(params *parameter) []int {
	var result []int
	num1, num2, nextTerm := 0, 1, 0

	if params.min != 0 {
		params.min = 0
	}

	for nextTerm <= params.max {
		result = append(result, nextTerm)
		num1 = num2
		num2 = nextTerm
		nextTerm = num1 + num2
	}

	return result
}

func generatePrimeNumbers(params *parameter) []int {
	var result []int

	if params.min < 2 {
		params.min = 2
	}

	if params.max < params.min {
		params.max = params.min
	}

	for num := params.min; num <= params.max; num++ {
		if !isPrime(num) {
			continue
		}
		result = append(result, num)
	}

	return result
}

func isPrime(num int) bool {
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return num > 1
}

func printNumbers(nums []int, sequenceType string, params *parameter) {
	fmt.Printf("\n%v (%v -> %v):\n", strings.ToUpper(sequenceType), params.min, params.max)
	for index, num := range nums {
		if index == len(nums)-1 {
			fmt.Println(num)
		} else {
			fmt.Print(num, " ")
		}
	}
}

func getInput(scanner *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func promptMinMax(scanner *bufio.Scanner, params *parameter, sequenceType string, generateNumber func(*parameter) []int) {
	var input string

	for true {
		params.min = 0

		if sequenceType == "even" || sequenceType == "odd" || sequenceType == "cube" {
			input = getInput(scanner, "\n>>> Minimum number (optional, default=0, type t to choose other type): ")

			if input == "" {
				params.min = 0
			} else {
				if input == "t" {
					break
				}

				min, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("\nInvalid input. You can only type a number or t.")
					continue
				}
				params.min = min
			}
		}

		input = getInput(scanner, "\n>>> Maximum number (type t to choose other type): ")

		if input == "" {
			fmt.Println("\nMaximum number is required")
			continue
		}

		if input == "t" {
			break
		}

		max, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("\nInvalid input. You can only type a number or t.")
			continue
		}
		params.max = max

		printNumbers(generateNumber(params), sequenceType, params)
	}
}

func getSequenceTypeOptions() map[string]string {
	return map[string]string{
		"1": "even",
		"2": "odd",
		"3": "square",
		"4": "cube",
		"5": "fibonacci",
		"6": "prime",
	}
}
