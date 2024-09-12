package Projet_Red

import (
	"fmt"
)

func (p *Player) Market() {
	var input string
	fmt.Println("1: 	Potion de vie")
	fmt.Println("2: 	Potion de poison")
	fmt.Println("3: 	Potion de vie GRATUITE !! (1 fois seulement)")
	fmt.Println("0: 	Retour")
	fmt.Print("Votre choix : ")
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("Vous avez achetez une potion de vie !")
		p.AddItem("Potion de santé" ,1 , "(+50pv)")
		p.Market()
	case "2":
		fmt.Println("Vous avez achetez une potion de poison !")
		p.AddItem("Potion de poison" ,1 ,"(-10pv/s pour 3s)")
		p.Market()
	case "3":
		var potgratos bool
	if potgratos {
		fmt.Println("Vous avez déjà eu votre potion gratuite !")
	}else {
		fmt.Println("Vous obtenez votre potion de vie gratuite !")
	p.AddItem("Potion de santé" ,1 , "(+50pv)")
	potgratos = true
	p.Market()
		}
	case "0":
		p.MainMenu()
	default:
		fmt.Println("Cette fois-ci veuillez choisir le bon chiffre !")
		p.Market()
	}

}
