package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) InventoryDisplay() {
	fmt.Println("=========================================================")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("||                 I N V E N T A I R E                 ||")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("=========================================================")
	fmt.Println("                                                         ")
	fmt.Println(" Tapez le chiffre correspondant pour utiliser votre objet")
	fmt.Println("                                                         ")
	if len(p.inventory) == 0 {
		fmt.Println("Votre inventaire est vide !")
		time.Sleep(2 * time.Second)
		fmt.Println("Retour au menu principal")
		time.Sleep(2 * time.Second)
		p.MainMenu()
	} else {
		for i, objet := range p.inventory {
			fmt.Println(i+1, "- 	", objet.name, "x", objet.quantity, "	", objet.description)
		}
	}
	fmt.Println("--------------------------------------------------------")
	fmt.Println("       Tapez 0 pour revenir au menu principal        ")
	fmt.Println("________________________________________________________")
	fmt.Println("--------------------------------------------------------")
	p.MenuInventaire()
}
func (p *Player) MenuInventaire() {
	var input int
	var selectedItem Inventory
	var items []Inventory = p.inventory
	fmt.Println("                  ")
	fmt.Scanln(&input)
	if input == 0 {
		p.MainMenu()
		fmt.Println("                  ")
	} else {
		fmt.Println("                  ")
		selectedItem = items[input-1]
		switch selectedItem.name {
		case "Potion de poison":
			p.Poison()
			fmt.Println("Vous avez bu une potion de poison.")
			fmt.Println("                  ")
			p.InventoryDisplay()
		case "Potion de santé":
			p.TakePot()
			fmt.Println("Vous avez bu une potion de santé")
			fmt.Println("                  ")
			p.InventoryDisplay()
		case "Livre de sort : Boule de feu":
			p.Spell()
			p.InventoryDisplay()
		default:
			p.InventoryDisplay()
		}
	}
}
