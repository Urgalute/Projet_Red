package Projet_Red

import "fmt"

func (p *Player) Mort() {
	fmt.Println("		Vous êtes mort, souhaitez vous réssuciter ?")
	fmt.Println("                  ")
	fmt.Println("1 - 			  		Oui !")
	fmt.Println("2 - 		Non, ça ira, le jeu est nul.")
	fmt.Println("                  ")
	var input string
	fmt.Scanln(&input)
	if input == "1" {
		p.hp = p.hpmax / 2
		fmt.Println("Vous allez réssuciter avec 50% de vos points de vie")
		fmt.Println("	Essayez d'être plus prudent à l'avenir !")
		p.MainMenu()
	} else if input == "2" {
		fmt.Println("Merci d'avoir jouer tout de même !")
	} else {
		fmt.Println("                  ")
		fmt.Println("Avez-vous bien lu vos choix possibles ??")
		fmt.Println("                  ")
		p.Mort()
	}
}
