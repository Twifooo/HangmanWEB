package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Structure de données pour passer les informations au template
type PageData struct {
	Title   string
	Message string
}

func homePage(w http.ResponseWriter, r *http.Request) {
	// Charger le template depuis le fichier
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Template non trouvé", http.StatusInternalServerError)
		return
	}

	// Remplir la structure avec des données à afficher
	data := PageData{
		Title:   "Bienvenue sur mon site",
		Message: "Ceci est un site dynamique généré avec Go !",
	}

	// Exécuter le template en y injectant les données
	tmpl.Execute(w, data)
}

func handleRequests() {
	// Associer la route principale à la fonction homePage
	http.HandleFunc("/", homePage)

	// Lancer le serveur
	fmt.Println("Le serveur est lancé sur le port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erreur lors du lancement du serveur:", err)
	}
}

func main() {
	handleRequests()
}
