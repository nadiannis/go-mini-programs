package main

import (
	"fmt"
	"strings"
)

type parameter struct {
	min, max int
}

func main() {
	params := new(parameter)
	params.max = 10

	printNumbers(generateEvenNumbers(params), "even")
	printNumbers(generateOddNumbers(params), "odd")
	printNumbers(generateSquareNumbers(params), "square")
	printNumbers(generateCubeNumbers(params), "cube")
	printNumbers(generateFibonacciNumbers(params), "fibonacci")
	printNumbers(generatePrimeNumbers(params), "prime")
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

	if params.max < params.min {
		params.max = params.min
	}

	for num := params.min; num <= params.max; num++ {
		result = append(result, num*num)
	}

	return result
}

func generateCubeNumbers(params *parameter) []int {
	var result []int

	if params.max < params.min {
		params.max = params.min
	}

	for num := params.min; num <= params.max; num++ {
		result = append(result, num*num*num)
	}

	return result
}

func generateFibonacciNumbers(params *parameter) []int {
	var result []int
	num1, num2, nextTerm := 0, 1, 0

	if params.min < 0 {
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

func printNumbers(nums []int, sequenceType string) {
	fmt.Printf("%v:\n", strings.ToUpper(sequenceType))
	for index, num := range nums {
		if index == len(nums) - 1 {
			fmt.Println(num)
		} else {
			fmt.Print(num, " ")
		}
	}
}