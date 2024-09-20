package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) CombatGardes() {

	for tour := 1; p.g1hp > 0 && p.hp > 0; tour++ {
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[4m__________________________\033[0m")
		fmt.Println("\033[30m\033[107m\033[4m        Tour :", tour, "         \033[0m")
		time.Sleep(500 * time.Millisecond)
		g1init := p.InitiativeG1() + 5
		time.Sleep(200 * time.Millisecond)
		playerinit := p.InitiativeP() + 233

		fmt.Println(" ")
		fmt.Println("Initiative de", p.name, ":", playerinit, " !")
		fmt.Println("Initiative du Garde 1 :", g1init, " !")
		fmt.Println(" ")
		if playerinit > g1init {
			if playerinit >= 18 {
				p.CharTurnCritGarde()

			} else {
				p.CharTurnGarde()
			}

		} else if playerinit < g1init {
			fmt.Println(" ")
			fmt.Println("Le Garde 1 attaque !")
			fmt.Println(" ")
			time.Sleep(500 * time.Millisecond)
			if g1init >= 18 {
				if tour%3 == 0 {
					fmt.Println(" ")
					fmt.Println("Attaque lourde critique ! Vous subissez 120 points de dégats !")
					fmt.Println(" ")
					p.hp -= 4 * p.g1damage
					fmt.Println(p.name, p.hp, "\033[92mPV\033[0m restants ")
					fmt.Println(" ")
				} else {
					fmt.Println(" ")
					fmt.Println("Vous subissez un coup critique de 60 points de dégats !")
					fmt.Println(" ")
					p.hp -= 2 * p.g1damage
					fmt.Println(p.name, p.hp, "\033[92mPV\033[0m restants ")
					fmt.Println(" ")
				}
			} else {
				if tour%3 == 0 {
					fmt.Println(" ")
					fmt.Println("Attaque lourde ! Vous subissez 60 points de dégats !")
					fmt.Println(" ")
					p.hp -= 2 * p.g1damage
					fmt.Println(p.name, p.hp, "\033[92mPV\033[0m restants ")
					fmt.Println(" ")
				} else {
					fmt.Println(" ")
					fmt.Println("Vous subissez 30 points de dégats !")
					fmt.Println(" ")
					p.hp -= p.g1damage
					fmt.Println(p.name, p.hp, "\033[92mPV\033[0m restants ")
					fmt.Println(" ")
				}
			}

		}
	}
	if p.g1hp <= 0 {
		fmt.Println(" ")
		fmt.Println("Vous avez tué le garde !")
		p.MainMenu()

		return

	} else if p.hp <= 0 {
		fmt.Println(" ")
		fmt.Println("Le Garde vous a démoli ... ")
		p.Mort()
		return
	}
}

func (p *Player) CharTurnGarde() {
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
		p.AttackBasicGarde()
	case "2":
		fmt.Println(" ")
		fmt.Println("Sort")
		fmt.Println(" ")
		p.AttackSpellGarde()
	case "3":
		fmt.Println(" ")
		fmt.Println("Inventaire")
		fmt.Println(" ")
		p.FightInventoryGarde()
	default:
		fmt.Println(" ")
		fmt.Println("Choix invalide, veuillez réessayer.")
		fmt.Println(" ")
		p.CharTurnGarde()
	}
}

func (p *Player) CharTurnCritGarde() {
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
		p.AttackCritGarde()
	case "2":
		fmt.Println(" ")
		fmt.Println("Sort")
		fmt.Println(" ")
		p.AttackSpellCritGarde()
	case "3":
		fmt.Println(" ")
		fmt.Println("Inventaire")
		fmt.Println(" ")
		p.FightInventoryCritGarde()
	default:
		fmt.Println("Choix invalide, veuillez réessayer.")
		p.CharTurnCritGarde()
	}
}

func (p *Player) FightInventoryGarde() {
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
		p.CharTurnCritGarde()
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

func (p *Player) FightInventoryCritGarde() {
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
		p.CharTurnCritGarde()
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
func (p *Player) UseInventoryCritGarde() {
	var input int
	var selectedItem Inventory
	var items []Inventory = p.inventory
	fmt.Println("                  ")
	fmt.Scanln(&input)
	if input == 0 {
		p.CharTurnCritGarde()
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

func (p *Player) UseInventoryGarde() {
	var input int
	var selectedItem Inventory
	var items []Inventory = p.inventory
	fmt.Println("                  ")
	fmt.Scanln(&input)
	if input == 0 {
		p.CharTurnGarde()
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

func (p *Player) AttackBasicGarde() {
	damage := p.dammage
	p.g1hp -= damage
	fmt.Println(" ")
	fmt.Println("Vous utilisez \033[93mAttaque basique\033[0m et infligez", damage, "dégâts à ", p.g1name)
	fmt.Println(" ")

	fmt.Println(p.g1name, p.g1hp, "\033[92mPV\033[0m restants ")
}

func (p *Player) AttackSpellGarde() {
	fmt.Println("1. \033[33mCoup de poing\033[0m")
	if p.CheckSpell("Boule de feu") {
		fmt.Println("2. \033[91mBoule de feu\033[0m")
	}
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		damage := 30
		p.g1hp -= damage
		fmt.Println(" ")
		fmt.Println("Vous utilisez \033[33mCoup de poing\033[0m et infligez", damage, "dégâts à", p.g1name)
		fmt.Println(" ")
		fmt.Println(p.g1name, p.g1hp, "\033[92mPV\033[0m restants ")
	case "2":
		if p.CheckSpell("Boule de feu") {
			if p.mana >= 25 {
				p.mana -= 25
				damage := 18
				p.g1hp -= damage
				fmt.Println(" ")
				fmt.Println(p.name, p.mana, "\033[94mmana\033[0m restants ")
				fmt.Println(" ")
				fmt.Println("Vous utilisez \033[91mBoule de feu\033[0m et infligez", damage, "dégâts à ", p.g1name)
				fmt.Println(" ")
				fmt.Println(p.g1name, p.g1hp, "\033[92mPV\033[0m restants ")
			} else {
				fmt.Println(" ")
				fmt.Println("Vous n'avez pas assez de \033[94mmana\033[0m !")
				fmt.Println(" ")
				p.AttackSpellGarde()
			}
		} else {
			fmt.Println("Vous n'avez pas cette compétence veuillez en choisir celle que vous avez !")
			fmt.Println(" ")
			p.AttackSpellGarde()
		}
	case "0":
		p.CharTurnGarde()
	default:
		fmt.Println("Cette fois-ci veuillez choisir le bon chiffre !")
		p.CharTurnGarde()
	}
}

func (p *Player) AttackCritGarde() {
	damage := p.dammage
	p.g1hp -= damage * 2
	fmt.Println(" ")
	fmt.Println("Vous utilisez \033[93mAttaque critique\033[0m et infligez", damage*2, "dégâts à ", p.g1name)
	fmt.Println(" ")
	fmt.Println(p.g1name, p.g1hp, "\033[92mPV\033[0m restants ")
}

func (p *Player) AttackSpellCritGarde() {
	fmt.Println("1. \033[33mCoup de poing\033[0m")
	if p.CheckSpell("Boule de feu") {
		fmt.Println("2. Boule de feu")
	}
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("\033[33mCoup de poing\033[0m")
		damage := 10
		p.g1hp -= damage * 2
		fmt.Println(" ")
		fmt.Println("Vous utilisez \033[33mCoup de poing critique\033[0m et infligez", damage*2, "dégâts à", p.g1name)
		fmt.Println(" ")
		fmt.Println(p.g1name, p.g1hp, "\033[92mPV\033[0m restants ")
	case "2":
		if p.CheckSpell("Boule de feu") {
			if p.mana >= 25 {
				p.mana -= 25
				damage := 18
				p.g1hp -= damage * 2
				fmt.Println(" ")
				fmt.Println(p.name, p.mana, "\033[94mmana\033[0m restants ")
				fmt.Println(" ")
				fmt.Println("Vous utilisez \033[91mBoule de feu\033[0m critique et infligez", damage*2, "dégâts à ", p.g1name)
				fmt.Println(" ")
				fmt.Println(p.g1name, p.g1hp, "\033[92mPV\033[0m restants ")
			} else {
				fmt.Println(" ")
				fmt.Println("Vous n'avez pas assez de \033[94mmana\033[0m !")
				fmt.Println(" ")
				p.AttackSpellCritGarde()
			}
		} else {
			fmt.Println("Vous n'avez pas cette compétence veuillez en choisir une que vous possédez !")
			p.AttackSpellCritGarde()
		}
	case "0":
		p.CharTurnCritGarde()
	default:
		fmt.Println("Cette fois-ci veuillez choisir le bon chiffre !")
		p.CharTurnCritGarde()
	}

}
