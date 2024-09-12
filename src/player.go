package Projet_Red

import "fmt"

func InitPlayer() {
	item01 := Inventory{"Potion", 3}

	player1 := Player{
		name:  "bob",
		class: "guerrier",
		level: 1,
		hpmax: 100,
		hp:    40,
		inventory: []Inventory{
			item01,
		},
	}
    player1.Display()
}

func (p *Player)Display(){
	fmt.Println("Votre nom :", p.name)
	fmt.Println("Votre classe :", p.class)
	fmt.Println("Votre niveau :", p.level)
	fmt.Println("Vos points de vie maximum :", p.hpmax)
	fmt.Println("Points de vie actuel :", p.hp)
	fmt.Println("Votre inventaire :", p.inventory)	
}
