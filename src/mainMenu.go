package Projet_Red

import (
	"fmt"
	"os"
)

// Menu Principal
func (p *Player) MainMenu() {
	var input string
	fmt.Println("---------------------------------")
	fmt.Println("--- Menu Principal ---")
	fmt.Println("1: 	Information personnage")
	fmt.Println("2: 	Inventaire")
	fmt.Println("3: 	Marchand")
	fmt.Println("4: 	Forgeron")
	fmt.Println("5: 	Equipements")
	fmt.Println("6: 	Combat")
	fmt.Println("0: 	Quitter")
	fmt.Print("Votre choix : ")
	fmt.Scanln(&input)
	switch input {
	case "1":
		ClearTerminal()
		fmt.Println("----------------")
		p.Display()
		p.MainMenu()
		return
	case "2":
		ClearTerminal()
		fmt.Println("----------------")
		p.InventoryDisplay()
		return
	case "3":
		ClearTerminal()
		fmt.Println("----------------")
		fmt.Println("Bienvenue chez le Marchand ! ")
		p.Market()
		return
	case "4":
		ClearTerminal()
		fmt.Println("----------------")
		p.BlackSmithMenu()
		p.MainMenu()
		return
	case "5":
		ClearTerminal()
		fmt.Println("----------------")
		p.GearSlot()
		return
	case "6":
		ClearTerminal()
		fmt.Println("----------------")
		p.Entrainement()
		return
	case "0":
		ClearTerminal()
		fmt.Println("----------------")
		fmt.Println("Merci d'avoir utilisé notre jeu!")
		os.Exit(0)
	default:
		ClearTerminal()
		fmt.Println("----------------")
		fmt.Println("Commande inconnue, réessayez.")
		p.MainMenu()
		return
	}
}
