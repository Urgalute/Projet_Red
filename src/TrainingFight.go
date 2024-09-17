package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) Entrainement() {
	fmt.Println("1.- Informations sur le gobelin d'entraînement ")
	fmt.Println("2.- Commencer le combat ")
	fmt.Println("  ")
	fmt.Println("  ")
	fmt.Println("0.- Pour revenir au menu principal")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		p.InitGobelin()
		fmt.Println(" ")
		p.Entrainement()
		return
	case "2":
		for tour := 1; (p.mhp != 0 && p.hp > 0) || (p.mhp > 0 && p.hp != 0); tour++ {
			playerinit := p.InitiativeP()+3
			time.Sleep(1 * time.Second)
			gobinit := p.InitiativeM()+1
			fmt.Println("Initiative de", p.name, ":", playerinit, " !")
			fmt.Println("Initiative du gobelin :", gobinit, " !")
			if playerinit > gobinit {
				fmt.Println(p.name, " attaque !")
				time.Sleep(500 * time.Millisecond)
				if playerinit >= 18 {
					fmt.Println("Attaque critique ! Vous infligez 20 dégats au gobelin")
					p.mhp -= 20
				}else {
					fmt.Println("Vous infligez 10 dégats au gobelin")
					p.mhp -= 10	
				}
			} else if playerinit < gobinit {
				fmt.Println("Le gobelin attaque !")
				time.Sleep(500 * time.Millisecond)
				if gobinit >= 18 {
					if tour%3 == 0 {
						fmt.Println("Attaque lourde critique ! Vous subissez 20 points de dégats !")
						p.hp -= 4*p.mdamage
					}else {
						fmt.Println("Vous subissez un coup critique de 10 points de dégats !")
						p.hp -= 2*p.mdamage
					}		
				}else {
					if tour%3 == 0 {
						fmt.Println("Attaque lourde ! Vous subissez 10 points de dégats !")
						p.hp -= 2*p.mdamage
					}else {
						fmt.Println("Vous subissez 5 points de dégats !")
						p.hp -= p.mdamage
					}		
				}
				
			}
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Tour :", tour)
			time.Sleep(500 * time.Millisecond)
		}
		return
	case "0":
		p.MainMenu()
		return
	}
}
