package Projet_Red

import (
	"fmt"
	"strings"
	"time"
)

func (p *Player) PersoInit() {
	fmt.Println("  ")
	fmt.Println("Bonjour ... Joueur ?? Ou peut être préférez vous que l'on vous appelle autrement ?")
	var inputnom string
	fmt.Scanln(&inputnom)
	inputnom = strings.TrimSpace(inputnom)
	if strings.Contains(inputnom, " ") {
		fmt.Println("Les espaces ne sont pas autorisés dans le prénom ! Réécrivez le sans espaces !")
		p.PersoInit()
	}
	fmt.Println("  ")
	prenom := []rune(inputnom)
	for i, r := range prenom {
		if r < 'A' || (r > 'Z' && r < 'a') || r > 'z' {
			fmt.Println("Attention, un caractère interdit s'est glissé dans votre prénom ! Réécrivez le sans !")
			p.PersoInit()

		}
		if r <= 'Z' && r >= 'A' {
			prenom[i] = r - ('A' - 'a')
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
	var hp int
	var hpmax int
	fmt.Scanln(&inputclasse)
	fmt.Println("  ")
	switch inputclasse {
	case "1":
		fmt.Println("La violence à l'état pur.")
		fmt.Println("  ")
		inputclasse = "Guerrier"
		hpmax = 120
		hp = hpmax / 2
	case "2":
		fmt.Println("Être sournois est une seconde nature pour vous.")
		fmt.Println("  ")
		inputclasse = "Voleur"
		hpmax = 100
		hp = hpmax / 2
	case "3":
		fmt.Println("Fini les tours de passe passe avec des cartes.")
		fmt.Println("  ")
		inputclasse = "Mage"
		hpmax = 80
		hp = hpmax / 2
	default:
		fmt.Println("Commande inconnue, réessayez.")
		fmt.Println("  ")
		p.PersoInit()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("  ")
	fmt.Println("Je pense avoir tout les renseignements qu'il me faut, bon jeu !")
	fmt.Println("  ")
	time.Sleep(1 * time.Second)
	p.InitPlayer(string(prenom), inputclasse, hpmax, hp)
	p.MainMenu()
}
