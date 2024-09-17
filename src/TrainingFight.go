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
		for tour := 1; p.mhp != 0; tour++ {
			playerinit := p.InitiativeP()+3
			time.Sleep(1 * time.Second)
			gobinit := p.InitiativeM()+1
			fmt.Println("Initiative de", p.name, ":", playerinit, " !")
			fmt.Println("Initiative du gobelin :", gobinit, " !")
			if playerinit > gobinit {
				fmt.Println(p.name, " commence !")

			} else if playerinit < gobinit {
				fmt.Println("Le gobelin commence !")
			}
			fmt.Println("Tour :", tour)
		}
	case "0":
		p.MainMenu()
	}
}
