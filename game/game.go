package game

import (
	"HangmanProject/ui"
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"os"
	"strings"
	"time"
)

var guessedLetters = map[rune]bool{}
var targetWord string

func SelectingWord() {

	// Open the dictionary file
	file, err := os.Open("game/dictionary.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the whole file content into a string
	var fileContent string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		line = strings.TrimSuffix(line, ",")
		fileContent += line + ","
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// remove last character , if last one is ','
	if fileContent[len(fileContent)-1] == ',' {

		fileContent = fileContent[:len(fileContent)-1]
	}

	// Split the file content by commas

	words := strings.Split(fileContent, ",")
	cleanWords := []string{}

	for _, word := range words {
		word = strings.TrimSpace(word)
		if word != "" {
			cleanWords = append(cleanWords, strings.Trim(word, "\""))
		}
	}

	// Clean quotes from words
	for i := range words {
		if i == (len(words) - 1) {
			break
		}

		// Remove the surrounding quotes if any
		words[i] = strings.Trim(cleanWords[i], "\"")
	}
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Number of words:", len(cleanWords))

	// Select a random word
	randomIndex := rand.Intn(len(cleanWords))
	randomWord := cleanWords[randomIndex]
	targetWord = randomWord
	// Print the random word

	for _, ch := range targetWord {

		guessedLetters[rune(ch)] = false
	}
	//	Send Word to ui for representing
	ui.WordStates(0, targetWord, guessedLetters)
}

func Play() {
	var stateNum int = 0
	incorrect := true
	reader := bufio.NewReader(os.Stdin)
	var input string
	//You should check the code located here. 1/3
	var wrongGuesses []rune
	//That's enough 1/3
	for stateNum < 9 {
		fmt.Print("Enter your Guess : ")

		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		for _, ch := range targetWord {

			if len(input) > 1 {
				if strings.EqualFold(input, targetWord) {
					ClearScreen()
					color.Green("You've survived ")
					ui.WordStates(10, targetWord, guessedLetters)
					return
				} else {
					ClearScreen()
					fmt.Print("The Target Word Was ")
					color.Green(targetWord)
					ui.WordStates(9, targetWord, guessedLetters)
					color.Red("YOU DEAD")
					return
				}
			} else if strings.EqualFold(string(ch), input) {
				ClearScreen()
				color.Blue("Good guess")
				incorrect = false
				guessedLetters[rune(ch)] = true
				ui.WordStates(stateNum, targetWord, guessedLetters)

			}
		}
		for _, guessed := range guessedLetters {
			if !guessed {

			}
		}
		if allGuessed() {
			ClearScreen()
			color.Green(targetWord)
			ui.WordStates(10, targetWord, guessedLetters)
			color.Green("You've survived")
			return
		}
		if incorrect {
			ClearScreen()
			fmt.Println("that wasn't good guess")
			stateNum++
			ui.WordStates(stateNum, targetWord, guessedLetters)
			//You should check the code located here. 2/3
			wrongGuesses = append(wrongGuesses, rune(input[0]))
			//That's enough 2/3
		}
		incorrect = true

		//You should check the code located here. 3/3
		fmt.Print("Wrong guesses: ")
		for _, ch := range wrongGuesses {
			fmt.Printf("%c ", ch)
		}
		fmt.Printf("\n")
		//That's enough 3/3
	}
	color.Cyan("Target word is ")
	color.Yellow(targetWord)
	color.Red("You DEAD")
}

func allGuessed() bool {
	for _, guessed := range guessedLetters {
		if !guessed {
			return false
		}
	}
	return true
}
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
