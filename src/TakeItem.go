package Projet_Red

import (
	"fmt"
	"time"
)

//Potion de santé
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
							fmt.Println("Vous avez bu une potion de santé")
							p.RemoveItem("Potion de santé", 1)
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
	fmt.Println("Vous n'avez pas de potion de santé dans votre inventaire")
	fmt.Println("--------------------------------")
}

//Potion de poison
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
		ClearTerminal()
		p.RemoveItem("Potion de poison", 1)
		fmt.Println("Vous avez bu une potion de poison.")
		for i := 0; i < 3; i++ {
			p.hp -= 10
			fmt.Println("                  ")
			fmt.Println("PV : ", p.hp, "/", p.hpmax)
			time.Sleep(1 * time.Second)
			if p.hp <= 0 {
				fmt.Println("                  ")
				p.Mort()
				return
			}
		}
		fmt.Println("                  ")
		p.InventoryDisplay()
	} else if input == "2" {
		p.InventoryDisplay()
		return
	} else {
		fmt.Println("                  ")
		fmt.Println("Avez-vous bien lu vos choix possibles ??")
		fmt.Println("                  ")
		p.Poison()
		return
	}
}

//Potion de mana
func (p *Player) Mana(){
	var input string
	for _, object := range p.inventory {
		
		if object.name == "Potion de mana" {
			if p.mana > p.manamax-20 {
				fmt.Println("--------------------------------")
				fmt.Println("Il vous manque moins de 20 points de mana", p.mana, "/", p.manamax, ", êtes-vous sûr de vouloir utiliser une potion de mana ?")
				fmt.Println("1: 	Oui")
				fmt.Println("2: 	Retour")
				fmt.Print("Votre choix : ")
				fmt.Scanln(&input)
				switch input {
				case "1":
					fmt.Println("--------------------------------")
					p.mana = p.manamax
					fmt.Println("Mana : ", p.mana, "/", p.manamax)
					for _, item := range p.inventory {
						if item.name == "Potion de mana" && item.quantity > 0 {
							p.RemoveItem("Potion de mana", 1)
						}
					}
				case "2":
					return
				default:
					p.Mana()
				}
			} else {
				fmt.Println("--------------------------------")
				p.mana += 25
				fmt.Println("Mana : ", p.mana, "/", p.manamax)
						p.RemoveItem("Potion de mana", 1)
					}
			return
		}
	}
	fmt.Println("--------------------------------")
	fmt.Println("Vous n'avez pas de potion de mana dans votre inventaire")
	fmt.Println("--------------------------------")


}

//Equipement d'équipement
func (p *Player) EquipGear(name string, quantity int, description string) bool {
	for i := 0; i < len(p.equipement); i++ {
        if p.equipement[i].name == name {
            return true
        }
    }
    p.equipement = append(p.equipement, Equipement{name, quantity, description})
    switch name {
    case "Casque en acier":
        p.hpmax += 15
    case "Robe magique":
        p.hpmax += 25
    case "Bottes en cuir":
        p.hpmax += 10
    default:
        fmt.Println("Aucun effet spécifique pour cet équipement.")
    }
    return false
}