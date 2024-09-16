package Projet_Red

import (
	"fmt"
	"time"
)

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
					for _, item := range p.inventory {
						if item.name == "Potion de santé" && item.quantity > 0 {
							//p.inventory[i].quantity--
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
				for _, item := range p.inventory {
					if item.name == "Potion de santé" && item.quantity > 0 {
						p.RemoveItem("Potion de santé", 1)
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

func (p *Player) Poison() {
	var input string
	fmt.Println("                  ")
	fmt.Println("Il s'agit d'une potion de poison, êtes-vous sûr de vouloir la boire ??")
	fmt.Println("                  ")
	fmt.Println("1 - 		Oui !")
	fmt.Println("2 - 	Non, ça ira.")
	fmt.Println("                  ")
	fmt.Scanln(&input)
	if input == "1" {
		p.RemoveItem("Potion de poison", 1)
		for i := 0; i < 3; i++ {
			p.hp -= 10
			fmt.Println("                  ")
			fmt.Println("PV : ", p.hp, "/", p.hpmax)
			time.Sleep(1 * time.Second)
			if p.hp <= 0 {
				fmt.Println("                  ")
				p.Mort()
			}
		}
		fmt.Println("                  ")
		p.InventoryDisplay()
	} else if input == "2" {
		p.InventoryDisplay()
	} else {
		fmt.Println("                  ")
		fmt.Println("Avez-vous bien lu vos choix possibles ??")
		fmt.Println("                  ")
		p.Poison()
	}
}

func (p *Player) EquipGear(name string, quantity int, description string) bool {
	if p.CheckGear(name) {
		for i := 0; i < len(p.equipement); i++ {
			if p.equipement[i].name == name {
				fmt.Println("Vous avez deja un item équipé !")
				return true
			}
		}
	} else if !p.CheckGear(name) {
		p.equipement = append(p.equipement, Equipement{name, quantity, description})
		if name == "Casque en acier" {
			p.hpmax += 15
		}else if name == "Robe magique" {
			p.hpmax += 25
		}else if name == "Bottes en cuir" {
			p.hpmax += 10
		}
	}
	return false
}
