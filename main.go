package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"math/rand"
	"strconv"
	"time"
)

func generateRandomNumber() int {
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := randomGenerator.Intn(100)
	return randomNumber
}

func showMenu() (string, error) {

	options := []string{
		"1. Easy",
		"2. Medium",
		"3. Hard",
		"4. Exit menu",
	}

	prompt := promptui.Select{
		Label: "Select difficulty",
		Items: options,
	}

	index, _, err := prompt.Run()
	if err != nil {
		fmt.Println("Prompt err", err)
	}

	return options[index], err

}

func playGame(gameType string, chances int) {
	fmt.Printf("Great, you have selected %s level, you have %v chances\n", gameType, chances)

	guessPrompt := promptui.Prompt{
		Label: "Enter your guess",
	}

	randomNumber := generateRandomNumber()
	start := time.Now()

	for {

		guessNumber, err := guessPrompt.Run()
		if err != nil {
			fmt.Printf("Error reding guess number %s\n", err)
			continue
		}

		guessNumberInt, err := strconv.Atoi(guessNumber)
		if err != nil {
			fmt.Printf("Error converting input to number: %s\n", err)
			continue
		}

		if guessNumberInt > randomNumber {
			fmt.Printf("Number is less than %v\n", guessNumberInt)
			chances--
		} else if guessNumberInt < randomNumber {
			fmt.Printf("Number is more than %v\n", guessNumberInt)
			chances--
		} else {
			fmt.Printf("Correct guess!\n")

			end := time.Now()
			diff := end.Sub(start)

			fmt.Printf("It took you %v seconds to win, with %v chances remaining.\n", int(diff.Seconds()), chances)
			break
		}

		if chances == 0 {
			fmt.Println("You lost")
			break

		}

	}

}

func main() {
	fmt.Printf("Welcome to the Number Guessing Game!\n " +
		"I'm thinking of a number between 1 and 100.\n" +
		"You have 5 chances to guess the correct number. \n")
	fmt.Println("\nPlease choose level of difficulty:")

	for {

		selectedLevel, err := showMenu()
		if err != nil {
			fmt.Println(err)
		}

		if selectedLevel == "4. Exit menu" {
			fmt.Println("Goodbye")
			break
		}

		switch selectedLevel {
		case "1. Easy":
			playGame("Easy", 10)
		case "2. Medium":
			playGame("Medium", 5)
		case "3. Hard":
			playGame("Hard", 3)
		}

	}

}
