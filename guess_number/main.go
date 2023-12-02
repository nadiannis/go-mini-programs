package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const min, max = 1, 100

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Println("\n============= GUESS THE NUMBER =============")
		input = getInput(scanner, ">>> Choose level (easy|medium|hard): ")
		level := strings.ToLower(input)
		var chance int

		switch level {
		case "easy":
			chance = 10
		case "medium":
			chance = 7
		case "hard":
			chance = 4
		default:
			fmt.Println("Invalid level")
			continue
		}

		randomNum := rand.Intn(max-min) + min
		guessesTaken := 0
		var guess int

		fmt.Println("\nI'm thinking of a number between 1 and 100.")
		fmt.Printf("You only have %v chances to guess.\n", chance)

		for guessesTaken < chance {
			input = getInput(scanner, "\n>>> Take a guess: ")

			convertedInput, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input. You can only type an integer.")
				continue
			}
			guess = convertedInput
			guessesTaken += 1

			if guess > randomNum {
				fmt.Println("Too high!")
			} else if guess < randomNum {
				fmt.Println("Your guess is too low.")
			} else {
				break
			}
		}

		displayResult(guess, randomNum, guessesTaken)
		promptPlayAgain(scanner, &input)
	}
}

func getInput(scanner *bufio.Scanner, prompt string) string {
	fmt.Printf(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func promptPlayAgain(scanner *bufio.Scanner, input *string) {
	*input = getInput(scanner, "\n>>> Play again? (y|n): ")
	*input = strings.ToLower(*input)

	if *input == "y" {
		return
	} else if *input == "n" {
		os.Exit(0)
	} else {
		fmt.Println("Invalid input")
		promptPlayAgain(scanner, input)
	}
}

func displayResult(guess, randomNum, guessesTaken int) {
	if guess == randomNum {
		fmt.Printf("\nCongratulations! You guessed my number in %v guesses!\n", guessesTaken)
	} else {
		fmt.Printf("\nSorry, you're wrong. The number is %v.\n", randomNum)
	}
}