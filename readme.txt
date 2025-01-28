# Projet Hangman

## Description
Ce projet est une implémentation web du jeu classique du Pendu (Hangman) en utilisant Go pour le backend et HTML/CSS pour le frontend. Le jeu offre différents niveaux de difficulté, un système de vies personnalisable, et un tableau de classement.

## Fonctionnalités
- Choix de la difficulté (facile, moyen, difficile)
- Nombre de vies personnalisable
- Pseudo personnalisé
- Nombre de lettre affiché selectionnable (0 à 3 lettres)
- Système de score
- Tableau de classement (n'est pas disponible)
- Interface web interactive

## Installation
1. Assurez-vous d'avoir Go installé sur votre machine.
2. Clonez ce dépôt : git clone https://github.com/Twifooo/HangmanWEB
3. Naviguez dans le dossier du projet : cd hangman
4. Lancez le serveur : go run main.go

5. Ouvrez votre navigateur et allez à `http://localhost:8080`

## Structure du Projet
- main.go : Point d'entrée de l'application, gestion des routes
- game.go : Logique du jeu
- word.go : Gestion de la sélection des mots
- leaderboard.go : Gestion du tableau de classement
- templates/ : Fichiers HTML pour l'interface utilisateur
- static/ : Fichiers CSS et autres ressources statiques
- data/ : Fichiers texte contenant les mots pour chaque niveau de difficulté

## Répartition des Tâches et Organisation
- Le projet a été réalisé en bînome, l'esthetique, et la logique de jeu a été pensé et écrit par les deux étudiants.
Une verification par OpenAI a été faite à plusieurs reprises afin de comprendre, seulement lors de long moment d'incompréhension, nos erreurs.


## Gestion du Temps et Priorités
1. Mise en place de la structure de base et de l'esthetique.
2. Développement de l'interface utilisateur.
3. Création de la logique de jeu.
4. Implémentation des fonctionnalités avancées comme le classement.
5. Tests continus tout au long du développement


## Défis Rencontrés et Solutions
- **Défi** : Creation de la logique de jeu sans l'utilisation de JavaScript.
**Solution** : Utilisation du pattern HTTP POST/Redirect/GET (Post/Redirect/Get) avec des templates Go.
- **Défi** : Mise à jour en temps réel de l'interface utilisateur.
**Solution** : Utilisation de templates Go pour générer dynamiquement le HTML.

## Améliorations Futures
- Ajout d'un mode multijoueur
- Intégration d'une base de données pour le stockage permanent des scores
- Amélioration de l'interface utilisateur avec des animations
