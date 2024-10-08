package Projet_Red

import (
	"fmt"
	"time"
)
var killgob int
var deathgob int
func (p *Player) Entrainement() {
	fmt.Println("\033[96m1.\033[0m- Informations sur le gobelin d'entraînement ")
	fmt.Println("\033[96m2.\033[0m- Commencer le combat ")
	fmt.Println("  ")
	fmt.Println("  ")
	fmt.Println("\033[96m0.\033[0m- Pour revenir au menu principal")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		p.InitGobelin()
		fmt.Println(" ")
		p.Entrainement()
		return
	case "2":
		for tour := 1; p.mhp > 0 && p.hp > 0; tour++ {
			fmt.Println(" ")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[4m__________________________\033[0m")
			fmt.Println("\033[30m\033[107m\033[4m        Tour :", tour,"         \033[0m")
			time.Sleep(500 * time.Millisecond)
			playerinit := p.InitiativeP() + 3
			time.Sleep(200 * time.Millisecond)
			gobinit := p.InitiativeM() + 1
			fmt.Println(" ")
			fmt.Println("Initiative de", p.name, ":", playerinit, " !")
			fmt.Println("Initiative du gobelin :", gobinit, " !")
			fmt.Println(" ")
			if playerinit > gobinit {
				if playerinit >= 18 {
					p.CharTurnCrit()

				} else {
					p.CharTurn()
				}

			} else if playerinit < gobinit {
				fmt.Println(" ")
				fmt.Println("Le gobelin attaque !")
				fmt.Println(" ")
				time.Sleep(500 * time.Millisecond)
				if gobinit >= 18 {
					if tour%3 == 0 {
						fmt.Println(" ")
						fmt.Println("Attaque lourde critique ! Vous subissez 20 points de dégats !")
						fmt.Println(" ")
						p.hp -= 4 * p.mdammage
						fmt.Println(p.name, p.hp, "\033[92mPV\033[0m restants ")
						fmt.Println(" ")
					} else {
						fmt.Println(" ")
						fmt.Println("Vous subissez un coup critique de 10 points de dégats !")
						fmt.Println(" ")
						p.hp -= 2 * p.mdammage
						fmt.Println(p.name, p.hp, "\033[92mPV\033[0m restants ")
						fmt.Println(" ")
					}
				} else {
					if tour%3 == 0 {
						fmt.Println(" ")
						fmt.Println("Attaque lourde ! Vous subissez 10 points de dégats !")
						fmt.Println(" ")
						p.hp -= 2 * p.mdammage
						fmt.Println(p.name, p.hp, "\033[92mPV\033[0m restants ")
						fmt.Println(" ")
					} else {
						fmt.Println(" ")
						fmt.Println("Vous subissez 5 points de dégats !")
						fmt.Println(" ")
						p.hp -= p.mdammage
						fmt.Println(p.name, p.hp, "\033[92mPV\033[0m restants ")
						fmt.Println(" ")
					}
				}

			}
		}
		if p.mhp <= 0 {
			fmt.Println(" ")
			fmt.Println("Vous avez tué le gobelin, on peut désormais vous décerner le titre de Goblin Slayer !")
			p.Experience()
			p.mhp = p.mhpmax
			killgob ++
			p.Titre()
			p.MainMenu()

			return

		} else if p.hp <= 0 {
			fmt.Println(" ")
			fmt.Println("Le gobelin vous a démoli ... Vous avez gagné le titre : le Faible")
			deathgob ++
			p.Titre()
			p.Mort()
			return
		}
		return
	case "0":
		p.MainMenu()
		return
	default:
		p.Entrainement()
		return
	}
}
