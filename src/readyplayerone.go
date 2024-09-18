package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) CharTurn() {
	var input string
	fmt.Println(" ")
	fmt.Println("A vous d'attaquer !")
	fmt.Println(" ")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Sort")
	fmt.Println("3. Inventaire")
	fmt.Print("Choisissez une option: ")
	fmt.Scanln(&input)

	switch input {
	case "1":
		fmt.Println(" ")
		fmt.Println("Attaquer")
		fmt.Println(" ")
		p.AttackBasic()
	case "2":
		fmt.Println(" ")
		fmt.Println("Sort")
		fmt.Println(" ")
		p.AttackSpell()
	case "3":
		fmt.Println(" ")
		fmt.Println("Inventaire")
		fmt.Println(" ")
		p.FightInventory()
	default:
		fmt.Println(" ")
		fmt.Println("Choix invalide, veuillez réessayer.")
		fmt.Println(" ")
		p.CharTurn()
	}
}

func (p *Player) CharTurnCrit() {
	var input string
	fmt.Println("A vous d'attaquer !")
	fmt.Println(" ")
	fmt.Println("1. Attaque Critique")
	fmt.Println("2. Sort Critique")
	fmt.Println("3. Inventaire")
	fmt.Print("Choisissez une option: ")
	fmt.Scanln(&input)

	switch input {
	case "1":
		fmt.Println(" ")
		fmt.Println("Attaque critique")
		fmt.Println(" ")
		p.AttackCrit()
	case "2":
		fmt.Println(" ")
		fmt.Println("Sort")
		fmt.Println(" ")
		p.AttackSpellCrit()
	case "3":
		fmt.Println(" ")
		fmt.Println("Inventaire")
		fmt.Println(" ")
		p.FightInventoryCrit()
	default:
		fmt.Println("Choix invalide, veuillez réessayer.")
		p.CharTurn()
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

func (p *Player) FightInventoryCrit() {
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
		p.CharTurnCrit()
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
	p.UseInventoryCrit()
}
func (p *Player) UseInventoryCrit() {
	var input int
	var selectedItem Inventory
	var items []Inventory = p.inventory
	fmt.Println("                  ")
	fmt.Scanln(&input)
	if input == 0 {
		p.CharTurnCrit()
		fmt.Println("                  ")
	} else if input < len(p.inventory)+1 {
		fmt.Println("                  ")
		selectedItem = items[input-1]
		switch selectedItem.name {
		case "Potion de poison":
			p.Poison()
			fmt.Println("Vous avez bu une potion de poison.")
			fmt.Println("                  ")
			p.FightInventoryCrit()
		case "Potion de santé":
			p.TakePot()
			fmt.Println("Vous avez bu une potion de santé")
			fmt.Println("                  ")
			p.FightInventoryCrit()
		case "Potion de mana":
			p.Mana()
			fmt.Println("Vous avez bu une potion de mana")
			fmt.Println("                  ")
			p.FightInventoryCrit()
		default:
			fmt.Println("Vous ne pouvez pas effectuer cette action en plein combat ! ")
			fmt.Println("                  ")
			p.FightInventoryCrit()
			return
		}
	} else {
		fmt.Println("Cette fois-ci veuillez choisir le bon chiffre !")
		p.FightInventoryCrit()
		return
	}
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
	damage := p.dammage
	p.mhp -= damage
	fmt.Println(" ")
	fmt.Println("Vous utilisez Attaque basique et infligez", damage, "dégâts à ", p.mname)
	fmt.Println(" ")

	fmt.Println(p.mname, p.mhp, " PV restants ")
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
		damage := 10
		p.mhp -= damage
		fmt.Println(" ")
		fmt.Println("Vous utilisez Coup de poing et infligez", damage, "dégâts à", p.mname)
		fmt.Println(" ")
		fmt.Println(p.mname, p.mhp, "PV restants ")
	case "2":
		if p.CheckSpell("Boule de feu") {
			fmt.Println("Boule de feu")
			if p.mana >= 25 {
				p.mana -= 25

				damage := 18
				p.mhp -= damage
				fmt.Println(" ")
				fmt.Println(p.name, p.mana, "Mana restants ")
				fmt.Println(" ")
				fmt.Println("Vous utilisez Boule de feu et infligez", damage, "dégâts à ", p.mname)
				fmt.Println(" ")
				fmt.Println(p.mname, p.mhp, "PV restants ")
			} else {
				fmt.Println(" ")
				fmt.Println("Vous n'avez pas assez de mana !")
				fmt.Println(" ")
				p.AttackSpellCrit()
			}
<<<<<<< HEAD
=======
			damage := p.dammage + 5
			p.mhp -= damage
			fmt.Println(" ")
			fmt.Println("Vous utilisez Boule de feu et infligez", damage, "dégâts à ", p.mname)
			fmt.Println(" ")
			fmt.Println(p.mname, p.mhp, "PV restants ")
>>>>>>> 2b704d3817a9bd4f7cecfdf8bc7414eeee6d7682
		} else {
			fmt.Println("Vous n'avez pas cette compétence veuillez en choisir celle que vous avez !")
			fmt.Println(" ")
			p.AttackSpellCrit()
		}
	case "0":
		p.CharTurn()
	default:
		fmt.Println("Cette fois-ci veuillez choisir le bon chiffre !")
		p.CharTurn()
	}
}

func (p *Player) AttackCrit() {
<<<<<<< HEAD
	damage := 5
	p.mhp -= damage * 2
=======
	damage := p.dammage
	p.mhp -= damage*2
>>>>>>> 2b704d3817a9bd4f7cecfdf8bc7414eeee6d7682
	fmt.Println(" ")
	fmt.Println("Vous utilisez Attaque critique et infligez", damage*2, "dégâts à ", p.mname)
	fmt.Println(" ")
	fmt.Println(p.mname, p.mhp, "PV restants ")
}

func (p *Player) AttackSpellCrit() {
	fmt.Println("1. Coup de poing")
	if p.CheckSpell("Boule de feu") {
		fmt.Println("2. Boule de feu")
	}
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("Coup de poing")
<<<<<<< HEAD
		damage := 8
		p.mhp -= damage * 2
=======
		damage := 10
		p.mhp -= damage*2
>>>>>>> 2b704d3817a9bd4f7cecfdf8bc7414eeee6d7682
		fmt.Println(" ")
		fmt.Println("Vous utilisez Coup de poing critique et infligez", damage*2, "dégâts à", p.mname)
		fmt.Println(" ")
		fmt.Println(p.mname, p.mhp, "PV restants ")
	case "2":
		if p.CheckSpell("Boule de feu") {
			fmt.Println("Boule de feu")
			if p.mana >= 25 {
				p.mana -= 25

				damage := 18
				p.mhp -= damage * 2
				fmt.Println(" ")
				fmt.Println(p.name, p.mana, "Mana restants ")
				fmt.Println(" ")
				fmt.Println("Vous utilisez Boule de feu critique et infligez", damage*2, "dégâts à ", p.mname)
				fmt.Println(" ")
				fmt.Println(p.mname, p.mhp, "PV restants ")
			} else {
				fmt.Println(" ")
				fmt.Println("Vous n'avez pas assez de mana !")
				fmt.Println(" ")
				p.AttackSpellCrit()
			}
<<<<<<< HEAD
=======
			damage := p.dammage+5
			p.mhp -= damage*2
			fmt.Println(" ")
			fmt.Println("Vous utilisez Boule de feu critique et infligez", damage*2, "dégâts à ", p.mname)
			fmt.Println(" ")
			fmt.Println(p.mname, p.mhp, "PV restants ")
>>>>>>> 2b704d3817a9bd4f7cecfdf8bc7414eeee6d7682
		} else {
			fmt.Println("Vous n'avez pas cette compétence veuillez en choisir une que vous possédez !")
			p.AttackSpellCrit()
		}
	case "0":
		p.CharTurnCrit()
	default:
		fmt.Println("Cette fois-ci veuillez choisir le bon chiffre !")
		p.CharTurnCrit()
	}

}
