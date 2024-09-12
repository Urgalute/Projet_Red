package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) InventoryDisplay() {
	fmt.Println("=========================================")
	fmt.Println("-----------------------------------------")
	fmt.Println("||        I N V E N T A I R E          ||")
	fmt.Println("-----------------------------------------")
	fmt.Println("=========================================")
	fmt.Println("                                         ")
	if len(p.inventory) == 0 {
		fmt.Println("Votre inventaire est vide !")
		time.Sleep(3 * time.Second)
		p.MainMenu()
	} else {
		for i, objet := range p.inventory {
			fmt.Println(i+1, "- 	", objet.name, "x", objet.quantity, "	")
		}
		p.MenuInventaire()
	}
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
		fmt.Println("glouglou la santé")
	}
}
