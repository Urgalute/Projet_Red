package Projet_Red

import (
	"fmt"
)

func (p *Player) Spell() {
	
	var input string
	fmt.Println("Voulez-vous apprendre le sort Boule de Feu ?")
	fmt.Println("1. Oui")
	fmt.Println("2. Non")
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("             ")
		fmt.Println("Vous avez appris le sort Boule de Feu !")
		fmt.Println("             ")
	case "2":
		fmt.Println("             ")
		fmt.Println("Vous avez décider de ne pas apprendre le sort Boule de feu !")
		fmt.Println("             ")
	default:
		fmt.Println("Cette fois-ci veuillez choisir le bon chiffre !")
		p.Spell()
	}
}

func (p *Player) LearnSpell() []skill{
	for _, skill := range []skill {
		if skill == "Livre de sort : Boule de feu"{
			fmt.Println("Vous avez déjà Boule de Feu")
		}
	}
}

