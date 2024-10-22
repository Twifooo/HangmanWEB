package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// Nombre maximum d'essais
const essaisMax = 6

// Fonction pour afficher un titre
func afficherTitre() {
	couleurTitre := color.New(color.FgCyan, color.Bold).SprintFunc()
	titre := `
  _    _       _                      
 | |  | |     | |                     
 | |__| | ___ | |__   __ _ _ __ ___  
 |  __  |/ _ \| '_ \ / _` + "`" + ` | '_ ` + "`" + ` _ \ 
 | |  | | (_) | | | | (_| | | | | | |
 |_|  |_|\___/|_| |_|\__,_|_| |_| |_|
`
	fmt.Println(couleurTitre(titre))
}

// Fonction pour afficher un message avec une animation
func afficherMessageAvecAnimation(message string) {
	for _, char := range message {
		fmt.Print(string(char))
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Println()
}

// Fonction pour afficher l'état du pendu avec couleur
func afficherPendu(erreurs int) {
	couleurPendu := color.New(color.FgGreen).SprintFunc()
	pendu := []string{
		`
         _______
        |/      |
    	|
        |
        |
        |
       _|___`,
		`
         _______
        |/      |
        |      (_)
        |
        |
        |
       _|___`,
		`
         _______
        |/      |
        |      (_)
        |       |
        |
        |
       _|___`,
		`
         _______
        |/      |
        |      (_)
        |      \|
        |
        |
       _|___`,
		`
         _______
        |/      |
        |      (_)
        |      \|/
        |
        |
       _|___`,
		`
         _______
        |/      |
        |      (_)
        |      \|/
        |       |
        |
       _|___`,
		`
         _______
        |/      |
        |      (_)
        |      \|/
        |       |
        |      / \
       _|___`,
	}

	fmt.Println(couleurPendu(pendu[erreurs]))
}

// Fonction pour lire les mots depuis un fichier
func lireMotsDepuisFichier(nomFichier string) ([]string, error) {
	file, err := os.Open(nomFichier)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var mots []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		mots = append(mots, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return mots, nil
}

// Fonction pour choisir un mot aléatoire dans la liste
func choisirMotAleatoire(mots []string) string {
	rand.Seed(time.Now().UnixNano())
	return strings.ToLower(mots[rand.Intn(len(mots))])
}

// Fonction pour révéler un certain nombre de lettres au début
func revelerLettres(mot string, nbLettres int) []rune {
	lettresRevelees := make([]rune, len(mot))
	for i := range lettresRevelees {
		lettresRevelees[i] = '_'
	}

	rand.Seed(time.Now().UnixNano())
	indicesDejaReveles := make(map[int]bool)

	for i := 0; i < nbLettres; i++ {
		for {
			index := rand.Intn(len(mot))
			if !indicesDejaReveles[index] {
				indicesDejaReveles[index] = true
				lettresRevelees[index] = rune(mot[index])
				break
			}
		}
	}

	return lettresRevelees
}

// Fonction pour gérer le jeu
func lancerJeu(mot string, nbLettresRevelees int) {
	erreurs := 0
	essaisRestants := essaisMax
	lettresDevinees := revelerLettres(mot, nbLettresRevelees)

	var lettreOuMot string
	lettresUtilisees := make(map[rune]bool)

	for erreurs < essaisMax && strings.Contains(string(lettresDevinees), "_") {
		afficherTitre()
		afficherPendu(erreurs)
		fmt.Println("Mot :", string(lettresDevinees))
		fmt.Printf("Essais restants : %d\n", essaisRestants)

		// Afficher les lettres utilisées
		fmt.Print("Lettres utilisées : ")
		for lettre := range lettresUtilisees {
			fmt.Printf("%c ", lettre)
		}
		fmt.Println()

		fmt.Print("Choisissez une lettre ou un mot : ")
		fmt.Scanln(&lettreOuMot)
		lettreOuMot = strings.ToLower(lettreOuMot)

		if len(lettreOuMot) == 1 {
			lettre := rune(lettreOuMot[0])
			if !lettresUtilisees[lettre] {
				lettresUtilisees[lettre] = true
				if strings.ContainsRune(mot, lettre) {
					for i, l := range mot {
						if l == lettre {
							lettresDevinees[i] = lettre
						}
					}
				} else {
					erreurs++
					essaisRestants--
					afficherMessageAvecAnimation("Mauvaise lettre !")
				}
			} else {
				afficherMessageAvecAnimation("Vous avez déjà deviné cette lettre.")
			}
		} else {
			if lettreOuMot == mot {
				lettresDevinees = []rune(mot)
			} else {
				erreurs += 2
				essaisRestants -= 2
				afficherMessageAvecAnimation("Mauvais mot ! Vous avez perdu deux essais.")
			}
		}
	}
	afficherPendu(erreurs)
	if erreurs >= essaisMax {
		afficherMessageAvecAnimation("Vous avez perdu ! Le mot était : " + mot)
	} else {
		afficherMessageAvecAnimation("Bravo ! Vous avez trouvé le mot : " + mot)
	}
}

// Fonction pour afficher le menu et choisir un fichier
func choisirFichier() string {
	fmt.Println("Choisissez une catégorie de mots :")
	fmt.Println("1: Noms propres (nomP.txt)")
	fmt.Println("2: Mots courants (mots.txt)")
	fmt.Println("3: Verbes (verbes.txt)")
	var choix int
	for {
		fmt.Print("Entrez le numéro correspondant à votre choix : ")
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			return "nomP.txt"
		case 2:
			return "mots.txt"
		case 3:
			return "verbes.txt"
		default:
			fmt.Println("Choix invalide. Veuillez entrer 1, 2 ou 3.")
		}
	}
}

// Fonction principale
func main() {
	fichierMots := choisirFichier()
	var nbLettresRevelees int
	fmt.Print("Combien de lettres voulez-vous révéler au début ? (1 recommandé et 3 maximum) : ")
	fmt.Scanln(&nbLettresRevelees)
	if nbLettresRevelees < 0 {
		fmt.Println("Le nombre de lettres à révéler doit être un entier positif.")
		return
	}
	if nbLettresRevelees > 3 {
		fmt.Println("Vous ne pouvez pas révéler plus de 3 lettres.")
		return
	}
	mots, err := lireMotsDepuisFichier(fichierMots)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}
	if len(mots) == 0 {
		fmt.Println("Le fichier est vide ou ne contient aucun mot valide.")
		return
	}
	mot := choisirMotAleatoire(mots)
	lancerJeu(mot, nbLettresRevelees)
}
