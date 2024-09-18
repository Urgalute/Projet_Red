package Projet_Red

import (
	"fmt"
)

func (p *Player) MenuAventure() {
	var input string
	fmt.Println("==================================")
	fmt.Println("============== Menu ==============")
	fmt.Println("1: 	Information personnage")
	fmt.Println("2: 	Inventaire")
	fmt.Println("3: 	Equipements")
	fmt.Println("4: 	Se reposer")
	fmt.Println("5: 	Reprendre la route")
	fmt.Println(" ")
	fmt.Println("0: 	Quitter")
	fmt.Println("==================================")
	fmt.Println("==================================")
	fmt.Println(" ")
	fmt.Print("Votre choix : ")
	fmt.Println(" ")
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("----------------")
		p.Display()
		p.MenuAventure()
		return
	case "2":
		fmt.Println("----------------")
		p.InventoryDisplayAventure()
		return
	case "3":
		fmt.Println("----------------")
		p.GearSlotAventure()
		return
	case "4":
		fmt.Println("\033[1mVous décidez de vous reposer avant de repartir.")
		fmt.Println("Vous récupérez 25% de vos points de vie maximum\033[0m")
		fmt.Println(" ")
		p.hp += (p.hpmax/4)
	case "5":
		fmt.Println("\033[1mVous vous remettez en route.\033[0m")
		fmt.Println(" ")
	case "0":
		var input string
		fmt.Println("----------------")
		fmt.Println(" ")
		fmt.Println("Êtes vous sûr de vouloir quitter. Il n'y a pas de sauvegarde !")
		fmt.Println("1.-	Oui")
		fmt.Println("2.-	Non")
		fmt.Println(" ")
		fmt.Scanln(&input)
		switch input {
		case "1":
			fmt.Println("Merci d'avoir joué !")
		case "2":
			fmt.Println("C'est reparti !")
			p.MenuAventure()
		default:
			fmt.Println("----------------")
			fmt.Println("Commande inconnue, réessayez.")
			p.MenuAventure()
			return
		}
	}
}
