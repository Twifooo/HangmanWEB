package main

import (
	"hangman/game"
	"html/template"
	"net/http"
	"strconv"
)

type GameData struct {
	Game    *game.Game
	Message string
}

var currentGame GameData

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/game", gameHandler)
	http.HandleFunc("/play", playHandler)
	http.HandleFunc("/leaderboard", leaderboardHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	currentGame.Game = nil
	tmpl := template.Must(template.ParseFiles("templates/landing.html"))
	tmpl.Execute(w, nil)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		pseudo := r.FormValue("pseudo")
		difficulty := r.FormValue("difficulty")
		revealedLetters, _ := strconv.Atoi(r.FormValue("revealedLetters"))
		lives, _ := strconv.Atoi(r.FormValue("lives"))

		if lives < 1 {
			lives = 6
		}

		currentGame.Game = game.NewGame(difficulty, revealedLetters, lives)
		currentGame.Game.Pseudo = pseudo
		currentGame.Message = "Nouvelle partie commencée"

		http.Redirect(w, r, "/play", http.StatusSeeOther)
	} else {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		tmpl.Execute(w, nil)
	}
}

func playHandler(w http.ResponseWriter, r *http.Request) {
	if currentGame.Game == nil {
		http.Redirect(w, r, "/game", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		guess := r.FormValue("guess")
		action := r.FormValue("action")

		if action == "quit" {
			currentGame.Message = "Partie abandonnée. Le mot était : " + currentGame.Game.Word
			currentGame.Game = nil
		} else if len(guess) == 1 {
			correct := currentGame.Game.GuessLetter(rune(guess[0]))
			if correct {
				currentGame.Message = "Bonne devinette !"
			} else {
				currentGame.Message = "Mauvaise devinette. Vous perdez une vie."
			}
		} else if len(guess) > 1 {
			correct := currentGame.Game.GuessWord(guess)
			if correct {
				currentGame.Message = "Bravo, vous avez deviné le mot !"
			} else {
				currentGame.Message = "Mauvaise devinette. Vous perdez deux vies."
			}
		}

		if currentGame.Game.IsGameOver() {
			if currentGame.Game.HasWon() {
				currentGame.Message = "Félicitations, vous avez gagné ! Le mot était : " + currentGame.Game.Word
			} else {
				currentGame.Message = "Désolé, vous avez perdu. Le mot était : " + currentGame.Game.Word
			}
			game.AddToLeaderboard(game.LeaderboardEntry{
				Pseudo: currentGame.Game.Pseudo,
				Score:  currentGame.Game.Score(),
			})
			currentGame.Game = nil
		}
	}

	tmpl := template.Must(template.ParseFiles("templates/game.html"))
	tmpl.Execute(w, currentGame)
}

func leaderboardHandler(w http.ResponseWriter, r *http.Request) {
	leaderboard := game.GetLeaderboard()
	tmpl := template.Must(template.ParseFiles("templates/leaderboard.html"))
	tmpl.Execute(w, leaderboard)
}
