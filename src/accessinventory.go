package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) InventoryDisplay() {
	if len(p.inventory) == 0 {
		fmt.Println("Votre inventaire est vide !")
		time.Sleep(3 * time.Second)
		p.MainMenu()
	} else {
		for _, objet := range p.inventory {
			fmt.Println("-", objet)
		}
	}
}
