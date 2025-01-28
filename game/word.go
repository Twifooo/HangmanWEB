package game

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func GetWord(difficulty string) string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Erreur lors de l'obtention du répertoire courant : %v\n", err)
		return "default"
	}
	dataDir := filepath.Join(dir, "data")
	fmt.Printf("Répertoire de données : %s\n", dataDir)

	files, err := ioutil.ReadDir(dataDir)
	if err != nil {
		fmt.Printf("Erreur lors de la lecture du répertoire data : %v\n", err)
	} else {
		fmt.Println("Contenu du répertoire data :")
		for _, file := range files {
			fmt.Println(file.Name())
		}
	}
	filename := filepath.Join(dataDir, difficulty+".txt")
	fmt.Printf("Tentative de lecture du fichier : %s\n", filename)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("Le fichier %s n'existe pas\n", filename)
		return "default"
	}
	words, err := readWordsFromFile(filename)
	if err != nil {
		fmt.Printf("ERREUR lecture fichier %s : %v\n", filename, err)
		return "default"
	}
	if len(words) == 0 {
		fmt.Printf("AUCUN MOT dans %s\n", filename)
		return "default"
	}
	rand.Seed(time.Now().UnixNano())
	selectedWord := words[rand.Intn(len(words))]
	fmt.Printf("Mot sélectionné : %s\n", selectedWord)
	return selectedWord
}
func readWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir le fichier %s : %v", filename, err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("erreur lors de la lecture du fichier %s : %v", filename, err)
	}
	fmt.Printf("Nombre de lignes lues dans %s : %d\n", filename, lineCount)
	fmt.Printf("Nombre de mots valides lus dans %s : %d\n", filename, len(words))
	if len(words) > 0 {
		fmt.Printf("Premier mot : %s, Dernier mot : %s\n", words[0], words[len(words)-1])
	}
	return words, nil
}
