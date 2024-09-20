package Projet_Red

import (
	"fmt"
)

func (p *Player) MenuAventure() {
	var input string
	fmt.Println("==================================")
	fmt.Println("============== \033[96mMenu\033[0m ==============")
	fmt.Println("\033[96m1\033[0m: 	Information personnage")
	fmt.Println("\033[96m2\033[0m: 	Inventaire")
	fmt.Println("\033[96m3\033[0m: 	Equipements")
	fmt.Println("\033[96m4\033[0m: 	Se reposer")
	fmt.Println("\033[96m5\033[0m: 	Reprendre la route")
	fmt.Println(" ")
	fmt.Println("\033[96m0\033[0m: 	Quitter")
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
		fmt.Println(" ")
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
