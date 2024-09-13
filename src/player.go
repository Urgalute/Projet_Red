package Projet_Red

import "fmt"

func (p *Player) InitPlayer(a string, b string) {
	item01 := Inventory{"Potion de santé", 1, "(+50pv)"}
	item02 := Inventory{"Potion de poison", 1, "(-10pv/s pour 3s)"}
	item03 := Inventory{"Potion de mana", 1, "(+35 mana)"}
	*p = Player{
		name:  a,
		class: b,
		level: 1,
		hpmax: 100,
		hp:    40,
		inventory: []Inventory{
			item01, item02,
		},
		skill: []string{"Coup de poing"},
	}
}

func (p *Player) Display() {
	fmt.Println("Votre nom :", p.name)
	fmt.Println("Votre classe :", p.class)
	fmt.Println("Votre niveau :", p.level)
	fmt.Println("Vos points de vie maximum :", p.hpmax)
	fmt.Println("Points de vie actuel :", p.hp)
}
