package Projet_Red

import "fmt"

func InitPlayer(a string, b string) {
	item01 := Inventory{"Potion", 3}
	item02 := Inventory{"Potion1", 3}

	player1 := Player{
		name:  a,
		class: b,
		level: 1,
		hpmax: 100,
		hp:    40,
		inventory: []Inventory{
			item01, item02,
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
