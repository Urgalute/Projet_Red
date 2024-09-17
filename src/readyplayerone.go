package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) CharTurn() {
	var input string
	fmt.Println("Menu:")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Sort")
	fmt.Println("3. Inventaire")
	fmt.Print("Choisissez une option: ")
	fmt.Scanln(&input)

	switch input {
	case "1":
		fmt.Println("Attaquer")
		p.AttackBasic()
	case "2":
		fmt.Println("Sort")
		p.AttackSpell()
	case "3":
		fmt.Println("Inventaire")
		p.FightInventory()
	default:
		fmt.Println("Choix invalide, veuillez réessayer.")
	}
}

func (p *Player) FightInventory() {
	fmt.Println("=========================================================")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("||                 I N V E N T A I R E                 ||")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("=========================================================")
	fmt.Println("                                                         ")
	fmt.Println(" Tapez le chiffre correspondant pour utiliser votre objet")
	fmt.Println("                                                         ")
	if len(p.inventory) == 0 {
		fmt.Println("Votre inventaire est vide !")
		time.Sleep(2 * time.Second)
		fmt.Println("Retour au menu principal")
		time.Sleep(2 * time.Second)
		p.CharTurn()
		return
	} else {

		for i, objet := range p.inventory {
			fmt.Println(i+1, "- 	", objet.name, "x", objet.quantity, "	", objet.description)
		}
	}
	fmt.Println("--------------------------------------------------------")
	fmt.Println("       Tapez 0 pour retourner au combat !      ")
	fmt.Println("________________________________________________________")
	fmt.Println("--------------------------------------------------------")
	p.UseInventory()
}
func (p *Player) UseInventory() {
	var input int
	var selectedItem Inventory
	var items []Inventory = p.inventory
	fmt.Println("                  ")
	fmt.Scanln(&input)
	if input == 0 {
		p.CharTurn()
		fmt.Println("                  ")
	} else if input < len(p.inventory)+1 {
		fmt.Println("                  ")
		selectedItem = items[input-1]
		switch selectedItem.name {
		case "Potion de poison":
			p.Poison()
			fmt.Println("Vous avez bu une potion de poison.")
			fmt.Println("                  ")
			p.FightInventory()
		case "Potion de santé":
			p.TakePot()
			fmt.Println("Vous avez bu une potion de santé")
			fmt.Println("                  ")
			p.FightInventory()
		case "Potion de mana":
			p.Mana()
			fmt.Println("Vous avez bu une potion de mana")
			fmt.Println("                  ")
			p.FightInventory()
		default:
			fmt.Println("Vous ne pouvez pas effectuer cette action en plein combat ! ")
			fmt.Println("                  ")
			p.FightInventory()
			return
		}
	} else {
		fmt.Println("Cette fois-ci veuillez choisir le bon chiffre !")
		p.FightInventory()
		return
	}
}

func (p *Player) AttackBasic() {
	damage := 5
	p.mhp -= damage
	fmt.Println("Vous utilisez Attaque basique et inflige", damage, "dégâts à ", p.mname,)
	fmt.Println(p.mname, p.mhp, "restants :",)
}

func (p *Player) AttackSpell() {
	fmt.Println("1. Coup de poing")
	if p.CheckSpell("Boule de feu") {
		fmt.Println("2. Boule de feu")
	}
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("Coup de poing")
		damage := 8
		p.mhp -= damage
		fmt.Println("Vous utilisez Coup de poing et inflige", damage ,"dégâts à", p.mname,)
		fmt.Println(p.mname, p.mhp, "restants :",)
	case "2":
		if p.CheckSpell("Boule de feu") {
		fmt.Println("Boule de feu")
		if p.mana >= 25 {
			p.mana -= 25
		}
		damage := 18
		p.mhp -= damage
		fmt.Println("Vous utilisez Boule de feu et inflige", damage, "dégâts à ", p.mname,)
		fmt.Println(p.mname, p.mhp, "restants :", )
	} else { 
		fmt.Println("Vous n'avez pas cette compétence veuillez en choisir celle que vous avez !")
		p.AttackSpell()
	}
	case "0":
		p.CharTurn()
	default:
		fmt.Println("Cette fois-ci veuillez choisir le bon chiffre !")
		p.CharTurn()
	}


}