package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type GameState struct {
	Mot             string   `json:"mot"`
	LettresDevinees []string `json:"lettres_devinees"`
	EssaisRestants  int      `json:"essais_restants"`
	Erreurs         int      `json:"erreurs"`
	EssaisMax       int      `json:"essais_max"`
	LettresUtilisees []string `json:"lettres_utilisees"`
}

type Score struct {
	Pseudo    string `json:"pseudo"`
	Categorie string `json:"categorie"`
	Resultat  string `json:"resultat"`
	Score     int    `json:"score"`
}

var (
	mots     []string
	gameData GameState
	scores   []Score
)

func init() {
	mots = loadWords("Data/verbes.txt")
	mots = append(mots, loadWords("Data/mots.txt")...)
	mots = append(mots, loadWords("Data/nomP.txt")...)
	loadScores()
}

func loadWords(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(strings.TrimSpace(string(data)), "\n")
}

func loadScores() {
	data, err := os.ReadFile("scores.json")
	if err == nil {
		json.Unmarshal(data, &scores)
	}
}

func saveScores() {
	data, err := json.MarshalIndent(scores, "", "  ")
	if err == nil {
		os.WriteFile("scores.json", data, 0644)
	}
}

func startGameHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	mot := mots[rand.Intn(len(mots))]
	lettresDevinees := make([]string, len(mot))
	for i := range lettresDevinees {
		lettresDevinees[i] = "_"
	}
	gameData = GameState{
		Mot:             mot,
		LettresDevinees: lettresDevinees,
		EssaisRestants:  6,
		Erreurs:         0,
		EssaisMax:       6,
		LettresUtilisees: []string{},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gameData)
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Lettre string `json:"lettre"`
	}
	json.NewDecoder(r.Body).Decode(&input)
	lettre := strings.ToLower(input.Lettre)

	if len(lettre) != 1 || !strings.ContainsAny(lettre, "abcdefghijklmnopqrstuvwxyz") {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if contains(gameData.LettresUtilisees, lettre) {
		http.Error(w, "Letter already used", http.StatusBadRequest)
		return
	}

	gameData.LettresUtilisees = append(gameData.LettresUtilisees, lettre)

	if strings.Contains(gameData.Mot, lettre) {
		for i, char := range gameData.Mot {
			if string(char) == lettre {
				gameData.LettresDevinees[i] = lettre
			}
		}
	} else {
		gameData.Erreurs++
		gameData.EssaisRestants--
	}

	if gameData.Erreurs >= gameData.EssaisMax || !contains(gameData.LettresDevinees, "_") {
		saveGameResult(r)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gameData)
}

func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

func saveGameResult(r *http.Request) {
	pseudo := r.Header.Get("Pseudo")
	categorie := r.Header.Get("Categorie")
	if pseudo == "" {
		pseudo = "Anonyme"
	}
	if categorie == "" {
		categorie = "Inconnue"
	}

	resultat := "win"
	if gameData.Erreurs >= gameData.EssaisMax {
		resultat = "lose"
	}
	score := Score{
		Pseudo:    pseudo,
		Categorie: categorie,
		Resultat:  resultat,
		Score:     gameData.EssaisRestants * 10,
	}
	scores = append(scores, score)
	saveScores()
}

func leaderboardHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(scores)
}

func main() {
	http.HandleFunc("/start", startGameHandler)
	http.HandleFunc("/guess", guessHandler)
	http.HandleFunc("/leaderboard", leaderboardHandler)

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img"))))
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))

	http.ListenAndServe(":8080", nil)
}
