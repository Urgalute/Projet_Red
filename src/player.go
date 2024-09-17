package Projet_Red

import "fmt"

func (p *Player) InitPlayer(nom string, classe string, pvmax int, pvactuel int) {
	gear01 := Inventory{"Casque en acier", 1, "+15pvmax"}
	gear02 := Inventory{"Robe magique", 1, "+25pvmax"}
	gear03 := Inventory{"Bottes en cuir", 1, "+10pvmax"}

	*p = Player{
		name:         nom,
		class:        classe,
		level:        1,
		hpmax:        pvmax,
		hp:           pvactuel,
		money:        100,
		inventorymax: 10,
		inventory:    []Inventory{gear01, gear02, gear03},
		skill:        []string{"Coup de poing"},
		equipement:   []Equipement{},
		mname:        "Gobelin vicieux",
		mhpmax:       25,
		mhp:          25,
		mdamage:      5,
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
	fmt.Println("Vos équipements :", p.equipement)
}

func (p *Player) Experience() {
	
}
