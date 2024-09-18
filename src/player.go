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
		mana:		  75,
		manamax:      100,
		dammage:      5,
		xp:           0,
		xpmax:        10,
		money:        100,
		inventorymax: 10,
		inventory:    []Inventory{gear01, gear02, gear03},
		skill:        []string{"Coup de poing"},
		equipement:   []Equipement{},
		mname:        "Gobelin vicieux",
		mhpmax:       40,
		mhp:          40,
		mdammage:     5,
		mxp:          5,
		mlevel:       1,
	}
}

func (p *Player) Display() {
	fmt.Println("Votre nom :", p.name)
	fmt.Println("Votre classe :", p.class)
	fmt.Println("Votre niveau :", p.level, "(", p.xp)")
	fmt.Println("Vos points de vie: ", p.hp, "/", p.hpmax)
	fmt.Println("Vos attaques :", p.skill)
}

func (p *Player) Experience() {
	if p.level == 6 {
		fmt.Println("Vous avez atteint le niveau maximum")
	} else {
		p.xp += p.mxp
		fmt.Println("Vous gagnez", p.mxp, "points d'expériences")
		switch p.xp {
		case 10:
			p.level = 2
			p.xpmax = 30
			fmt.Println("Félicitations !! Vous passez au niveau", p.level, ". Vos statistique augmentent, mais celles de vos ennemies aussi !")
		case 30:
			p.level = 3
			p.xpmax = 60
			fmt.Println("Félicitations !! Vous passez au niveau", p.level, ". Vos statistique augmentent, mais celles de vos ennemies aussi !")
		case 60:
			p.level = 4
			p.xpmax = 100
			fmt.Println("Félicitations !! Vous passez au niveau", p.level, ". Vos statistique augmentent, mais celles de vos ennemies aussi !")
		case 100:
			p.level = 5
			p.xpmax = 150
			fmt.Println("Félicitations !! Vous passez au niveau", p.level, ". Vos statistique augmentent, mais celles de vos ennemies aussi !")
		case 150:
			p.level = 6
			fmt.Println("Félicitations !! Vous passez au niveau", p.level, "(Niveau maximum). Vos statistique augmentent, mais celles de vos ennemies aussi !")
		}
		switch p.level {
		case 1:
			p.dammage = 5
			p.mxp = 5
			p.mhpmax = 40
			p.mdammage = 5
			p.mlevel = 1
		case 2:
			p.dammage = 8
			p.mxp = 10
			p.mhpmax = 50
			p.mdammage = 8
			p.hpmax += 5
			p.manamax += 10
			p.mlevel = 2
		case 3:
			p.dammage = 11
			p.mxp = 15
			p.mhpmax = 60
			p.mdammage = 11
			p.hpmax += 5
			p.manamax += 10
			p.mlevel = 3
		case 4:
			p.dammage = 14
			p.mxp = 20
			p.mhpmax = 70
			p.mdammage = 14
			p.hpmax += 5
			p.manamax += 10
			p.mlevel = 4
		case 5:
			p.dammage = 17
			p.mxp = 25
			p.mhpmax = 80
			p.mdammage = 17
			p.hpmax += 5
			p.manamax += 10
			p.mlevel = 5
		case 6:
			p.dammage = 20
			p.mhpmax = 90
			p.mdammage = 20
			p.hpmax += 5
			p.manamax += 10
			p.mlevel = 6
		}
	}
}
