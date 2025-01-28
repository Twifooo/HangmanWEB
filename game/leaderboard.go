package game

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type LeaderboardEntry struct {
	Pseudo string
	Score  int
}

func AddToLeaderboard(entry LeaderboardEntry) {
	leaderboard := GetLeaderboard()
	leaderboard = append(leaderboard, entry)
	sort.Slice(leaderboard, func(i, j int) bool {
		return leaderboard[i].Score > leaderboard[j].Score
	})
	if len(leaderboard) > 10 {
		leaderboard = leaderboard[:10]
	}
	saveLeaderboard(leaderboard)
}

func GetLeaderboard() []LeaderboardEntry {
	file, err := os.Open("leaderboard.json")
	if err != nil {
		return []LeaderboardEntry{}
	}
	defer file.Close()

	var leaderboard []LeaderboardEntry
	json.NewDecoder(file).Decode(&leaderboard)
	return leaderboard
}

func saveLeaderboard(leaderboard []LeaderboardEntry) {
	file, err := os.Create("leaderboard.json")
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde du classement:", err)
		return
	}
	defer file.Close()

	json.NewEncoder(file).Encode(leaderboard)
}
