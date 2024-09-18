package Projet_Red

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

//Création de personnage
func (p *Player) PersoInit() {
	fmt.Println("  ")
	fmt.Println("Bonjour ... Joueur ?? Ou peut être préférez vous que l'on vous appelle autrement ?")
	reader := bufio.NewReader(os.Stdin)
	inputnom, _ := reader.ReadString('\n')
	inputnom = strings.ReplaceAll(inputnom, " ", "")
	inputnom = strings.TrimSpace(inputnom)
	if inputnom == "" {
		fmt.Println("Le nom ne peut pas être vide.")
		p.PersoInit()
		return
	}
	fmt.Println("  ")
	var resultnom string
	prenom := []rune(inputnom)
	for i, r := range prenom {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			if r <= 'Z' && r >= 'A' {
				prenom[i] = r - ('A' - 'a')
			}
			resultnom += string(r)
		} else {
			fmt.Println("L'un de vos caractères n'est pas pris en charge.")
			p.PersoInit()
			return
		}
	}
	prenom[0] -= 32
	fmt.Println("  ")
	fmt.Println("Donc, vous vous appelez", string(prenom), ", sympa comme prénom")
	fmt.Println("  ")
	time.Sleep(1 * time.Second)
	fmt.Println("Et quelle classe vous inspire le plus,", string(prenom), "?")
	fmt.Println("  ")
	time.Sleep(1 * time.Second)
	fmt.Println("1 -  Guerrier : PVmax = 120")
	fmt.Println("2 -  Voleur : PVmax = 100")
	fmt.Println("3 -  Mage : PVmax = 80")
	fmt.Println("  ")
	var inputclasse string
	var hpmax int
	fmt.Scanln(&inputclasse)
	fmt.Println("  ")
	switch inputclasse {
	case "1":
		fmt.Println("La violence à l'état pur.")
		fmt.Println("  ")
		inputclasse = "Guerrier"
		hpmax = 120
	case "2":
		fmt.Println("Être sournois est une seconde nature pour vous.")
		fmt.Println("  ")
		inputclasse = "Voleur"
		hpmax = 100
	case "3":
		fmt.Println("Fini les tours de passe passe avec des cartes.")
		fmt.Println("  ")
		inputclasse = "Mage"
		hpmax = 80
	default:
		fmt.Println("Commande inconnue, réessayez.")
		fmt.Println("  ")
		p.PersoInit()
		return
	}
	time.Sleep(1 * time.Second)
	fmt.Println("  ")
	fmt.Println("Je pense avoir tout les renseignements qu'il me faut, bon jeu !")
	fmt.Println("  ")
	time.Sleep(1 * time.Second)
	p.InitPlayer(string(prenom), inputclasse, hpmax)
	p.MainMenu()
}
