package Projet_Red

import "fmt"

//Retirer un item de l'inventaire, rentrer un nom et une qty
func (p *Player) RemoveItem(name string, quantity int) {
	for i, item := range p.inventory {
		if item.name == name {
			p.inventory[i].quantity -= quantity
			if p.inventory[i].quantity <= 0 {
				p.inventory = append(p.inventory[:i], p.inventory[i+1:]...)
			}
		}
	}
}

func (p *Player) AddItem(name string, quantity int, description string) {
	for _, item := range p.inventory {
		fmt.Println("var" ,item.name, name)
		if item.name == name {
			item.quantity += quantity
			fmt.Println("qty" ,item.quantity, quantity)
		} else {
			p.inventory = append(p.inventory, Inventory{name, quantity, description})
		}
	}
}
