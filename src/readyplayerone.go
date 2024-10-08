package Projet_Red

import (
	"fmt"
	"time"
)

// Menu combat du joueur
func (p *Player) CharTurn() {
	var input string
	fmt.Printf("%s: %d/%d\n", p.mname, p.mhp, p.mhpmax)
	fmt.Printf("%s: %d/%d\n", p.name, p.hp, p.hpmax)
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

// Menu combat du joueur avec critique
func (p *Player) CharTurnCrit() {
	var input string
	fmt.Printf("%s: %d/%d\n", p.mname, p.mhp, p.mhpmax)
	fmt.Printf("%s: %d/%d\n", p.name, p.hp, p.hpmax)
	fmt.Println(" ")
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
		p.CharTurnCrit()
	}
}

// Affichage inventaire de combat
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
			fmt.Println("\033[96m",i+1,"\033[96m", "- 	", objet.name, "x", objet.quantity, "	", objet.description)
		}
	}
	fmt.Println("--------------------------------------------------------")
	fmt.Println("       Tapez \033[96m0\033[0m pour retourner au combat !      ")
	fmt.Println("________________________________________________________")
	fmt.Println("--------------------------------------------------------")
	p.UseInventory()
}

// Affichage inventaire de combat mais pour garder le critique
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
			fmt.Println("\033[96m",i+1,"\033[96m", "- 	", objet.name, "x", objet.quantity, "	", objet.description)
		}
	}
	fmt.Println("--------------------------------------------------------")
	fmt.Println("       Tapez \033[96m0\033[0m pour retourner au combat !      ")
	fmt.Println("________________________________________________________")
	fmt.Println("--------------------------------------------------------")
	p.UseInventoryCrit()
}

// Menu d'utilisation d'un objet dans l'inventaire avec critique
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
		case "Potion de \033[32mpoison\033[0m":
			p.Poison()
			fmt.Println("Vous avez bu une potion de \033[32mpoison\033[0m")
			fmt.Println("                  ")
			p.FightInventoryCrit()
		case "Potion de \033[31msanté\033[0m":
			p.TakePot()
			fmt.Println("Vous avez bu une potion de \033[31msanté\033[0m")
			fmt.Println("                  ")
			p.FightInventoryCrit()
		case "Potion de \033[94mmana\033[0m":
			p.Mana()
			fmt.Println("Vous avez bu une potion de\033[34mmana\033[0m")
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

// Menu inventaire de combat
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
		case "Potion de \033[32mpoison\033[0m":
			p.Poison()
			fmt.Println("Vous avez bu une potion de \033[32mpoison\033[0m")
			fmt.Println("                  ")
		case "Potion de \033[31msanté\033[0m":
			p.TakePot()
			fmt.Println("Vous avez bu une potion de \033[31msanté\033[0m")
			fmt.Println("                  ")
		case "Potion de \033[34mmana\033[0m":
			p.Mana()
			fmt.Println("Vous avez bu une potion de \033[34mmana\033[0m")
			fmt.Println("                  ")
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

// Attaque simple
func (p *Player) AttackBasic() {
	damage := p.dammage
	p.mhp -= damage
	fmt.Println(" ")
	fmt.Println("Vous utilisez \033[93mAttaque basique\033[0m et infligez", damage, "dégâts à ", p.mname)
	fmt.Println(" ")

	fmt.Println(p.mname, p.mhp, "\033[92mPV\033[0m restants ")
}

// Menu des sorts
func (p *Player) AttackSpell() {
	fmt.Println("1. \033[33mCoup de poing\033[0m", "\033[94m(15 PM)\033[0m")
	if p.CheckSpell("Boule de feu") {
		fmt.Println("2. \033[91mBoule de feu\033[0m")
	}
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		if p.mana >= 15 {
			p.mana -= 15
			damage := 10
			p.mhp -= damage
			fmt.Println(p.name, p.mana, "\033[94mmana\033[0m restants ")
			fmt.Println(" ")
			fmt.Println("Vous utilisez \033[33mCoup de poing\033[0m et infligez", damage, "dégâts à", p.mname)
			fmt.Println(" ")
			fmt.Println(p.mname, p.mhp, "\033[92mPV\033[0m restants ")
		} else {
			fmt.Println(" ")
			fmt.Println("Vous n'avez pas assez de \033[94mmana\033[0m !")
			fmt.Println(" ")
			p.AttackSpell()
		}
	case "2":
		if p.CheckSpell("Boule de feu") {
			if p.mana >= 25 {
				p.mana -= 25
				damage := p.dammage + 5
				p.mhp -= damage
				fmt.Println(" ")
				fmt.Println(p.name, p.mana, "\033[94mmana\033[0m restants ")
				fmt.Println(" ")
				fmt.Println("Vous utilisez \033[91mBoule de feu\033[0m et infligez", damage, "dégâts à ", p.mname)
				fmt.Println(" ")
				fmt.Println(p.mname, p.mhp, "\033[92mPV\033[0m restants ")
			} else {
				fmt.Println(" ")
				fmt.Println("Vous n'avez pas assez de \033[94mmana\033[0m !")
				fmt.Println(" ")
				p.AttackSpell()
			}
		} else {
			fmt.Println("Vous n'avez pas cette compétence veuillez en choisir celle que vous avez !")
			fmt.Println(" ")
			p.AttackSpell()
		}
	case "0":
		p.CharTurn()
	default:
		fmt.Println("Cette fois-ci veuillez choisir le bon chiffre !")
		p.CharTurn()
	}
}

func (p *Player) AttackCrit() {
	damage := p.dammage
	p.mhp -= damage * 2
	fmt.Println(" ")
	fmt.Println("Vous utilisez \033[93mAttaque critique\033[0m et infligez", damage*2, "dégâts à ", p.mname)
	fmt.Println(" ")
	fmt.Println(p.mname, p.mhp, "\033[92mPV\033[0m restants ")
}

func (p *Player) AttackSpellCrit() {
	fmt.Println("\033[96m1.\033[0m \033[33mCoup de poing\033[0m", "\033[94m(15 PM)\033[0m")
	if p.CheckSpell("Boule de feu") {
		fmt.Println("\033[96m2.\033[0m Boule de feu", "\033[94m(15 PM)\033[0m")
	}
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("\033[33mCoup de poing\033[0m")
		if p.mana >= 15 {
			p.mana -= 15
			damage := 10
			p.mhp -= damage * 2
			fmt.Println(p.name, p.mana, "\033[94mmana\033[0m restants ")
			fmt.Println(" ")
			fmt.Println("Vous utilisez \033[33mCoup de poing critique\033[0m et infligez", damage*2, "dégâts à", p.mname)
			fmt.Println(" ")
			fmt.Println(p.mname, p.mhp, "\033[92mPV\033[0m restants ")
		} else {
			fmt.Println(" ")
			fmt.Println("Vous n'avez pas assez de \033[94mmana\033[0m !")
			fmt.Println(" ")
			p.AttackSpellCrit()
		}
	case "2":
		if p.CheckSpell("Boule de feu") {
			if p.mana >= 25 {
				p.mana -= 25
				damage := 18
				p.mhp -= damage * 2
				fmt.Println(" ")
				fmt.Println(p.name, p.mana, "\033[94mmana\033[0m restants ")
				fmt.Println(" ")
				fmt.Println("Vous utilisez \033[91mBoule de feu\033[0m critique et infligez", damage*2, "dégâts à ", p.mname)
				fmt.Println(" ")
				fmt.Println(p.mname, p.mhp, "\033[92mPV\033[0m restants ")
			} else {
				fmt.Println(" ")
				fmt.Println("Vous n'avez pas assez de \033[94mmana\033[0m !")
				fmt.Println(" ")
				p.AttackSpellCrit()
			}
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
