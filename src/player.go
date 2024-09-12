package Projet_Red

import "fmt"

func (p *Player) InitPlayer(a string, b string) {
	item01 := Inventory{"Potion de santé", 2, "(+50pv)"}
	item02 := Inventory{"Potion de poison", 1, "(-10pv/s pour 3s)"}

	*p = Player{
		name:  a,
		class: b,
		level: 1,
		hpmax: 100,
		hp:    40,
		inventory: []Inventory{
			item02, item01,
		},
	}

}

func (p *Player) Display() {
	fmt.Println("Votre nom :", p.name)
	fmt.Println("Votre classe :", p.class)
	fmt.Println("Votre niveau :", p.level)
	fmt.Println("Vos points de vie maximum :", p.hpmax)
	fmt.Println("Points de vie actuel :", p.hp)
}
