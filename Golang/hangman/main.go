package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var WordList = []string{
	"DEVELOPER", "PROGRAM",
}

type Game struct {
	word           string
	guessedLetters map[rune]bool
	attemptsLeft   int
	maxAttempts    int
}

func NewGame(word string) *Game {
	return &Game{
		word:           strings.ToUpper(word),
		guessedLetters: make(map[rune]bool),
		attemptsLeft:   6,
		maxAttempts:    6,
	}
}

func getRandomWord() string {
	return strings.ToUpper(WordList[rand.Intn(len(WordList))])
}

func DrawHangman(attemptsLeft int) {
	stages := []string{
		`
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
		`
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`,
		`
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`,
		`
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`,
		`
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`,
		`
  +---+
  |   |
  O   |
      |
      |
      |
=========`,
		`
  +---+
  |   |
      |
      |
      |
      |
=========`,
	}

	index := attemptsLeft
	fmt.Println(stages[index])
}
func (g *Game) WelcomeUser() {
	g.maxAttempts = 6
	g.attemptsLeft = g.maxAttempts
	fmt.Println("====Welcome to the game of Hangman====")
	fmt.Println("You have attempts left: ", g.attemptsLeft, " out of total attempts: ", g.maxAttempts)
	DrawHangman(g.attemptsLeft)
}

func ConvertIntoRune(char string) rune {
	return rune(strings.ToUpper(char)[0])
}

func (g *Game) GuessWord(letter rune) (correct, alreadyGuessed bool) {
	if g.guessedLetters[letter] {
		return false, true
	}
	g.guessedLetters[letter] = true
	if strings.ContainsRune(g.word, letter) {
		return true, false
	}
	g.attemptsLeft--
	return false, false
}

func (g *Game) DisplayGuessedLetters() {
	fmt.Println("Guessed Letters: ")
	for char, _ := range g.guessedLetters {
		fmt.Printf("%c ", char)
	}
	fmt.Println()
}

func (g *Game) DisplayLetters(letter rune) string {
	fmt.Println("Guessing word: ")
	var guessedWord strings.Builder
	for _, ch := range g.word {
		if g.guessedLetters[ch] {
			guessedWord.WriteRune(ch)
			fmt.Printf("%c ", ch)
		} else {
			fmt.Printf("_")
		}
	}
	fmt.Println()
	return guessedWord.String()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	word := getRandomWord()
	fmt.Println("random word is ", word)

	game := NewGame(word)
	game.WelcomeUser()
	fmt.Println("Press Enter to start...")
	_, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	for {
		fmt.Println("Enter a letter...")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("input is invalid letter")
		}
		input = strings.TrimSpace(input)

		if len([]rune(input)) > 1 {
			fmt.Println("input should contain only 1 letter...")
			continue
		}
		letter := ConvertIntoRune(input)

		correct, alreadyGuessed := game.GuessWord(letter)
		game.DisplayGuessedLetters()

		fmt.Println("Attempts Left: ", game.attemptsLeft)
		DrawHangman(game.attemptsLeft)
		switch {
		case correct:
			fmt.Println("The letter is correctly guessed...")
		case alreadyGuessed:
			fmt.Println("The letter guessed already!!! please try again...")
		default:
			fmt.Println("Nope! That letter isn’t in the word.")
		}

		if game.attemptsLeft == 0 {
			fmt.Println("The attempts are exhausted. Better Luck Next Time...")
			break
		}
		guessedWord := game.DisplayLetters(letter)
		if guessedWord == game.word {
			fmt.Println("You have won THE MATCH")
			os.Exit(0)
		}
	}
	fmt.Println("The correct word is ", game.word)
}
