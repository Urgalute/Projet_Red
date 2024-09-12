package Projet_Red

import "fmt"

func (p *Player) TakePot() {
	var input string
	if p.hp > p.hpmax-50 {
		fmt.Println("--------------------------------")
		fmt.Println("Il vous manque moins de 50 points de vie, êtes-vous sur de vouloir utiliser une potion de vie ?")
		fmt.Println("1: 	Oui")
		fmt.Println("2: 	Retour")
		fmt.Print("Votre choix : ")
		fmt.Scanln(&input)
		switch input {
		case "1":
			p.hp = p.hpmax
			for i, item := range p.inventory {
				if item.name == "Potion de santé" && item.quantity > 0 {
					p.inventory[i].quantity--
				}
			}
		case "2":
			return
		default:
			p.TakePot()
		}
	} else {
		p.hp += 50
		for i, item := range p.inventory {
			if item.name == "Potion de santé" && item.quantity > 0 {
				p.inventory[i].quantity--
			}
		}
	}
}
