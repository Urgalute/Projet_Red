package Projet_Red

import "fmt"

func (p *Player) InitPlayer(nom string, classe string, pvmax int, pvactuel int) {
	gear01 := Inventory{"Casque en acier", 1, "+15pvmax || 5po"}
	gear02 := Inventory{"Robe magique", 1, "+25pvmax || 5po"}
	gear03 := Inventory{"Bottes en cuir", 1, "+10pvmax || 5po"}
	*p = Player{
		name:  nom,
		class: classe,
		level: 1,
		hpmax: pvmax,
		hp:    pvactuel,
		money: 100,
		inventory: []Inventory{gear01, gear02, gear03},
		skill: []string{"Coup de poing"},
		equipement: []Equipement{},
	}
}

func (p *Player) Display() {
	fmt.Println("Votre nom :", p.name)
	fmt.Println("Votre classe :", p.class)
	fmt.Println("Votre niveau :", p.level)
	fmt.Println("Vos points de vie maximum :", p.hpmax)
	fmt.Println("Points de vie actuel :", p.hp)
	fmt.Println("Vos attaques :", p.skill)
	fmt.Println("Vous avez :", p.money, "pièces d'or.")
}
