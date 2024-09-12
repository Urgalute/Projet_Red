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
		p.MenuInventaire()
	}
	fmt.Println("                                                        ")
	fmt.Println("________________________________________________________")
	fmt.Println("--------------------------------------------------------")
}
func (p *Player) MenuInventaire() {
	var input int
	var selectedItem Inventory
	var items []Inventory = p.inventory
	fmt.Scanln(&input)
	selectedItem = items[input-1]
	switch selectedItem.name {
	case "Potion de poison":
		fmt.Println("glouglou le poison")
	case "Potion de santé":
		p.TakePot()
		fmt.Println("glouglou la santé")
	}
}
