package main

import (
	"HangmanProject/game"
	"fmt"
	// "fmt"
	// "strings"
)

func main() {
	for {
		game.SelectingWord()
		game.Play()
		fmt.Println("Do you wanna Play again ? 'y' Or 'n'")
		var i string
		fmt.Scanf("%s", &i)
		if i != "y" {
			break
		}

	}
	//Derive a word we have to guess --------------> Done

	//Printing game states
	//	*Print Word you're guessing
	//	*Print Hangman state
	//Read User Input
	//	* Validate it
	//Determin if letter is a correct guess or not
	//	*if correct , update the guessed or not
	//	*if incorrect, update the hangman state
	//If word iks guessed -> game over, you win
	//If hangman is complete -> game over , you lose

}
