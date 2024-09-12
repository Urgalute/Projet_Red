package Projet_Red

import "fmt"

func (p *Player) TakePot() {
	var input string
	for _, object := range p.inventory {
		if object.name == "Potion de santé" {
		if p.hp > p.hpmax-50 {
			fmt.Println("--------------------------------")
			fmt.Println("Il vous manque moins de 50 points de vie", p.hp, "/", p.hpmax, ", êtes-vous sûr de vouloir utiliser une potion de vie ?")
			fmt.Println("1: 	Oui")
			fmt.Println("2: 	Retour")
			fmt.Print("Votre choix : ")
			fmt.Scanln(&input)
			switch input {
			case "1":
				fmt.Println("--------------------------------")
				p.hp = p.hpmax
				fmt.Println("PV : ", p.hp, "/", p.hpmax)
				for i, item := range p.inventory {
					if item.name == "Potion de santé" && item.quantity > 0 {
						p.inventory[i].quantity--
						p.RemoveItem("Potion de santé", 1)
					}
				}
			case "2":
				return
			default:
				p.TakePot()
			}
		} else {
			fmt.Println("--------------------------------")
			p.hp += 50
			fmt.Println("PV : ", p.hp, "/", p.hpmax)
			for i, item := range p.inventory {
				if item.name == "Potion de santé" && item.quantity > 0 {
					p.inventory[i].quantity--
				}
			}
		}
		return
	}
} 
	fmt.Println("--------------------------------")
	fmt.Println("Vous n'avez pasa de potion de santé dans votre inventaire")
	fmt.Println("--------------------------------")
}
