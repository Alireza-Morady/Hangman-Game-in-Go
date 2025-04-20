package ui

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var Mainaddrress = "Resources/Hangman"

func WordStates(State int, targetWord string, guessedLetters map[rune]bool) {
	if State < 9 {
		for _, ch := range targetWord {

			if ch == ' ' {
				fmt.Print(" ")

			} else if guessedLetters[ch] {
				fmt.Printf("%c", ch)

			} else {
				fmt.Print("_")

			}
			fmt.Print(" ")
		}
		fmt.Printf("\t\t")
		for i := 0; i < (9 - State); i++ {
			fmt.Print("❤️ ")
		}
		fmt.Println()
	}

	address := Mainaddrress + strconv.Itoa(State)

	file, err := os.Open(address)
	if err != nil {
		fmt.Println("can not open the hangman file !!!")
		panic(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("can not Read  file ")
		panic(err)
	}
	fmt.Println(string(content))

}
