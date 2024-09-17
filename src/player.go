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
		manamax:      100,
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
	if p.level == 6 {
		fmt.Println("Vous avez atteint le niveau maximum")
	} else {
		p.xp += p.mxp 
		fmt.Println("Vous gagnez", p.mxp,"points d'expériences")
	switch p.xp{
		case 20: p.level = 2 ; p.xpmax = 40 
		fmt.Println("Félicitations !! Vous passez au niveau", p.level,". Vos statistique augmentent, mais celles de vos ennemies aussi !")
		case 40: p.level = 3 ; p.xpmax = 60
		fmt.Println("Félicitations !! Vous passez au niveau", p.level,". Vos statistique augmentent, mais celles de vos ennemies aussi !")
		case 60: p.level = 4 ; p.xpmax = 80
		fmt.Println("Félicitations !! Vous passez au niveau", p.level,". Vos statistique augmentent, mais celles de vos ennemies aussi !")
		case 80: p.level = 5 ; p.xpmax = 100
		fmt.Println("Félicitations !! Vous passez au niveau", p.level,". Vos statistique augmentent, mais celles de vos ennemies aussi !")
		case 100: p.level = 6
		fmt.Println("Félicitations !! Vous passez au niveau", p.level,"(Niveau maximum). Vos statistique augmentent, mais celles de vos ennemies aussi !")
		}	
	switch p.level{
		case 1: p.dammage = 5 ; Gobelin.xp = 5 ; Gobelin.maxlife = 40 ; Gobelin.dammage = 3 ; Gobelin.level = 1
		case 2: p.dammage = 8 ; Gobelin.xp = 10 ; Gobelin.maxlife = 50 ; Gobelin.dammage = 5 ; p.maxlife += 5 ; p.maxmana += 10 ; Gobelin.level = 2
		case 3: p.dammage = 11 ; Gobelin.xp = 15 ; Gobelin.maxlife = 60 ; Gobelin.dammage = 7 ; p.maxlife += 5 ; p.maxmana += 10  ; Gobelin.level = 3; Gobelin.loot["Pièces d'or"] = 13
		case 4: p.dammage = 14 ; Gobelin.xp = 20 ; Gobelin.maxlife = 70 ; Gobelin.dammage = 9 ; p.maxlife += 5 ; p.maxmana +=10  ; Gobelin.level = 4
		case 5: p.dammage = 17 ; Gobelin.xp = 25 ; Gobelin.maxlife = 80 ; Gobelin.dammage = 11 ; p.maxlife += 5 ; p.maxmana += 10 ; Gobelin.level = 5
		case 6: p.dammage = 20 ; Gobelin.maxlife = 90 ; Gobelin.dammage = 13 ; p.maxlife += 5 ; p.maxmana += 10 ; Gobelin.level = 6 ; Gobelin.loot["Pièces d'or"] = 20
		}
	}
}
