package Projet_Red

import "fmt"

//Retirer un item de l'inventaire, rentrer un nom et une qty
func (p *Player) RemoveItem(name string, quantity int) {
	for i := 0; i < len(p.inventory); i++ {
		if p.inventory[i].name == name {
			p.inventory[i].quantity -= quantity
			if p.inventory[i].quantity <= 0 {
				fmt.Println(i)
				p.inventory = append(p.inventory[:i], p.inventory[i+1:]...)
				break
			}
		}
	}
}

//Ajouter un item à l'inventaire, rentrer un nom, une qty et une description
func (p *Player) AddItem(name string, quantity int, description string) {
	if p.CheckItem(name) {
		for i := 0; i < len(p.inventory); i++ {
			if p.inventory[i].name == name {
                p.inventory[i].quantity += quantity
                return
            }
		}
	} else if !p.CheckItem(name) {
		p.inventory = append(p.inventory, Inventory{name, quantity, description})
	}
}

//Check si un objet est présent dans l'inventaire du joueur
func (p *Player) CheckItem(name string) bool {
	for _, item := range p.inventory {
		if item.name == name {
			return true
		}
	}
	return false
}

func (p *Player) CheckQtyItem(name string) int {
	count := 0
	for _, item := range p.inventory {
        if item.name == name {
			count += item.quantity
        }
    }
    return count
}