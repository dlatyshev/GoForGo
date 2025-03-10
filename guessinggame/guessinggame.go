package guessinggame

import (
	"fmt"
	"math/rand"
)

func Game() {
	fmt.Println("Welcome to the Guessing Game!")
	secretNumber := rand.Intn(100) + 1

	max_guesses := 7

	for i := 1; i <= max_guesses; {
		var guess int
		fmt.Printf("Guess #%d: ", i)
		_, err := fmt.Scanln(&guess)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
		} else {
			if guess == secretNumber {
				fmt.Println("Congratulations! You guessed the number.")
				break
			} else if guess < secretNumber {
				fmt.Println("Too low. Try again.")
			} else {
				fmt.Println("Too high. Try again.")
			}
			i++
			if i > max_guesses {
				fmt.Printf("Sorry, you've run out of guesses. The secret number was %d.\n", secretNumber)
			}
		}
	}
}
