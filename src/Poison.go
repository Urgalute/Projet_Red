package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) Poison() {
	var input string
	fmt.Println("                  ")
	fmt.Println("Il s'agit d'une potion de poison, êtes-vous sûr de vouloir la boire ??")
	fmt.Println("                  ")
	fmt.Println("1 - 		Oui !")
	fmt.Println("2 - 	Non, ça ira.")
	fmt.Println("                  ")
	fmt.Scanln(&input)
	if input == "1"{
	for i := 0 ; i < 3 ; i++{
		p.hp -= 10
		fmt.Println("PV : " , p.hp,"/", p.hpmax)
		time.Sleep(1 * time.Second)
		if p.hp <= 0 {
			fmt.Println("                  ")
			fmt.Println("Vous venez de mourir des dégâts de poison")
			// Fonction vous êtes mort
		}
	}
	fmt.Println("                  ")
	p.InventoryDisplay()
	}else if input == "2" {
		p.InventoryDisplay()
	}else {
		fmt.Println("                  ")
		fmt.Println("Avez-vous bien lu vos choix possibles ??")
		fmt.Println("                  ")
		p.Poison()
	}
}