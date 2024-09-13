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
	for i, item := range p.inventory {
		fmt.Println("Name: ", name, "item.Name: ", item.name, " quantity: ", quantity, "Item: ", item, "p.inventory[0].name", p.inventory[i].name)
		if name == p.inventory[i].name {
			fmt.Println("if")
			p.inventory[i].quantity += quantity
			break
		} else {
			p.inventory = append(p.inventory, Inventory{name, quantity, description})
			fmt.Println("append")
		}
	}
}
