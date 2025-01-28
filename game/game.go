package game

import (
	"strings"
	"unicode"
)

type Game struct {
	Word            string
	GuessedLetters  map[rune]bool
	Lives           int
	InitialLives    int
	Difficulty      string
	RevealedLetters int
	Pseudo          string
}

func NewGame(difficulty string, revealedLetters int, lives int) *Game {
	word := GetWord(difficulty)
	game := &Game{
		Word:            word,
		GuessedLetters:  make(map[rune]bool),
		Lives:           lives,
		InitialLives:    lives,
		Difficulty:      difficulty,
		RevealedLetters: revealedLetters,
	}

	for i := 0; i < revealedLetters; i++ {
		if i < len(word) {
			game.GuessedLetters[rune(word[i])] = true
		}
	}

	return game
}

func (g *Game) GuessLetter(letter rune) bool {
	letter = rune(strings.ToLower(string(letter))[0])
	if !unicode.IsLetter(letter) {
		return false
	}
	if g.GuessedLetters[letter] {
		return false
	}

	g.GuessedLetters[letter] = true

	if !strings.ContainsRune(g.Word, letter) {
		g.Lives--
		return false
	}

	return true
}

func (g *Game) GuessWord(word string) bool {
	guessRunes := []rune(strings.ToLower(word))
	wordRunes := []rune(strings.ToLower(g.Word))

	if len(guessRunes) == len(wordRunes) {
		match := true
		for i := range guessRunes {
			if guessRunes[i] != wordRunes[i] {
				match = false
				break
			}
		}

		if match {
			for _, letter := range g.Word {
				g.GuessedLetters[letter] = true
			}
			return true
		}
	}

	g.Lives -= 2
	return false
}

func (g *Game) DisplayWord() string {
	display := ""
	for _, letter := range g.Word {
		if g.GuessedLetters[letter] {
			display += string(letter)
		} else {
			display += "_"
		}
		display += " "
	}
	return strings.TrimSpace(display)
}

func (g *Game) IsGameOver() bool {
	return g.Lives <= 0 || g.HasWon()
}

func (g *Game) HasWon() bool {
	for _, letter := range g.Word {
		if !g.GuessedLetters[letter] {
			return false
		}
	}
	return true
}

func (g *Game) Score() int {
	return len(g.Word)*10 + g.Lives*5
}

func (g *Game) GuessedLettersString() []string {
	var guessedLetters []string
	for letter := range g.GuessedLetters {
		guessedLetters = append(guessedLetters, string(letter))
	}
	return guessedLetters
}
