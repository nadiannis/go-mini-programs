package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	options := getOptions()
	var playerChoice string
	
	for {
		getInput(&playerChoice, "\nRock, Paper, Scissors?\n>>> Choose one (q to quit): ")
		playerChoice = strings.ToLower(playerChoice)
		computerChoice := options[rand.Intn(len(options))]

		if playerChoice == "q" {
			break
		}

		result, err := decide(playerChoice, computerChoice)
		if err != nil {
			fmt.Println(err)
			continue
		}

		displayResult(result, playerChoice, computerChoice)
	}
}

func getInput(input *string, prompt string) {
	fmt.Print(prompt)
	fmt.Scan(input)
	*input = strings.TrimSpace(*input)
}

func decide(playerChoice, computerChoice string) (string, error) {
	var result string
	options := getOptions()

	if !containsString(options, playerChoice) {
		return "", errors.New("\nInvalid input. You can only choose rock, paper, or scissors.")
	}

	if playerChoice == computerChoice {
		result = "tie"
	} else if playerChoice == "rock" {
		if computerChoice == "paper" {
			result = "lose"
		} else {
			result = "win"
		}
	} else if playerChoice == "paper" {
		if computerChoice == "rock" {
			result = "win"
		} else {
			result = "lose"
		}
	} else if playerChoice == "scissors" {
		if computerChoice == "rock" {
			result = "lose"
		} else {
			result = "win"
		}
	}

	return result, nil
}

func displayResult(result, playerChoice, computerChoice string) {
	actions := getActions()

	fmt.Println("\n* Your choice:", playerChoice)
	fmt.Println("* Computer's choice:", computerChoice)

	if result == "win" {
		fmt.Printf("\nYou win!\n%s %s %s\n", capitalize(playerChoice), actions[playerChoice], computerChoice)
		return
	}

	if result == "lose" {
		fmt.Printf("\nYou lose!\n%s %s %s\n", capitalize(computerChoice), actions[computerChoice], playerChoice)
		return
	}

	if result == "tie" {
		fmt.Println("\nTie!")
		return
	}
}

func containsString(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

func capitalize(str string) string {
	return strings.ToUpper(strings.Split(str, "")[0]) + str[1:]
}

func getOptions() []string {
	return []string{"rock", "paper", "scissors"}
}

func getActions() map[string]string {
	return map[string]string{
		"rock":     "smashes",
		"paper":    "covers",
		"scissors": "cut",
	}
}
