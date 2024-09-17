package Projet_Red

import "fmt"

func (p *Player) Entrainement() {
	fmt.Println("1.- Informations sur le gobelin d'entraînement ")
	fmt.Println("2.- Commencer le combat ")
	fmt.Println("  ")
	fmt.Println("  ")
	fmt.Println("0.- Pour revenir au menu principal")
	var input string
	fmt.Scanln(input)
	switch input {
	case "1":
		p.InitGobelin()
		p.Entrainement()
		return
	case "2":
		for tour := 1; p.mhp != 0; tour++ {
			fmt.Println("Initiative de", p.name, ":", p.InitiativeP(), "!")
			fmt.Println("Initiative du gobelin :", p.InitiativeM(), "!")
			if p.InitiativeP() > p.InitiativeM() {
				fmt.Println(p.name, " commence !")
			} else if p.InitiativeP() < p.InitiativeM() {
				fmt.Println("Le gobelin commence !")
			}
		}
	case "0":
		p.MainMenu()
	}
}
