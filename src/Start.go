package Projet_Red

import (
	"fmt"
	"log"
	"os"

	"github.com/eiannone/keyboard"
)

func (p *Player) StartMenu() {
	// Détection des touches du clavier
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	// Liste des options
	menu := []string{
		"	Nouvelle Partie			",
		"	Options				",
		"	Quitter				",
		"	Crédits				",
	}

	selectedIndex := 0 // L'index de l'élément sélectionné

	for {
		// Affichage du menu
		fmt.Println("\033[2J") // Clear screen
		fmt.Println("                           ")
		fmt.Println(`	     _       __     __   ____    _        _____ 
     /\     | |      \ \   / /  / __ \  | |      |_   _|
    /  \    | |       \ \_/ /  | |  | | | |        | |  
   / /\ \   | |        \   /   | |  | | | |        | |  
  / ____ \  | |____     | |    | |__| | | |____   _| |_ 
 /_/    \_\ |______|    |_|     \____/  |______| |_____|
                                                        `)
		fmt.Println(" ")
		fmt.Println(" ")
		for i, option := range menu {
			if i == selectedIndex {
				fmt.Printf("---> %s\n", option) // Afficher une flèche pour l'élément sélectionné
			} else {
				fmt.Printf("   %s\n", option)
			}
		}

		// Lire la touche pressée
		_, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		// Navigation avec les flèches haut et bas
		if key == keyboard.KeyArrowDown {
			selectedIndex++
			if selectedIndex >= len(menu) {
				selectedIndex = 0 // Retourner au début
			}
		} else if key == keyboard.KeyArrowUp {
			selectedIndex--
			if selectedIndex < 0 {
				selectedIndex = len(menu) - 1 // Retourner à la fin
			}
		}

		// Valider la sélection avec la touche "Entrée"
		if key == keyboard.KeyEnter {
			switch selectedIndex {
			case 0:
				p.MainMenu()
			case 1:

			case 2:
				fmt.Println(" ")
				fmt.Println("Au revoir !")
				fmt.Println(" ")
				os.Exit(0)
			case 3: 

			default :
				fmt.Println("Comment avez-vous fait pour lire ça ???")
			p.StartMenu()

			}
		}
	}
}
