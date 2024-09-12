package Projet_Red

import "fmt"

func (p *Player) InventoryDisplay() {
	fmt.Println(p.inventory)
	if  len(p.inventory) == 0 {
		fmt.Println("Votre inventaire est vide !")
	}else {
		for _, objet := range p.inventory {
			fmt.Println("-", objet)
		}
	}
}