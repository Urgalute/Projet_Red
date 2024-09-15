package Projet_Red

import "fmt"

func (p *Player) InitPlayer(nom string, classe string, pvmax int, pvactuel int) {
	item01 := Inventory{"Potion de santé", 1, "(+50pv)"}
	item02 := Inventory{"Potion de poison", 1, "(-10pv/s pour 3s)"}
	item03 := Inventory{"Potion de mana", 1, "(+25mana)"}
	item04 := Inventory{"Fourrure de loup", 1, "(C'est doux)"}
	item05 := Inventory{"Peau de troll", 1, "(C'est légérement poisseux, de la bave probablement)"}
	item06 := Inventory{"Cuir de sanglier", 1, "(Aussi robuste qu'un nain !)"}
	item07 := Inventory{"Plume de corbeau", 1, "(Le chatouilleur 2000 ultime)"}
	*p = Player{
		name:  nom,
		class: classe,
		level: 1,
		hpmax: pvmax,
		hp:    pvactuel,
		money: 100,
		inventory: []Inventory{
			item01, item02, item03, item04, item05, item06, item07,
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
	fmt.Println("Vos attaques :", p.skill)
	fmt.Println("Vous avez :", p.money, "pièces d'or.")
}
