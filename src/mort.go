package Projet_Red

import "fmt"

func (p *Player) Mort() {
	fmt.Println("\033[1m	Vous êtes mort, souhaitez vous réssuciter ?")
	fmt.Println("                  ")
	fmt.Println("\033[96m1\033[0m\033[1m - 			  Oui !			   	")
	fmt.Println("\033[96m2\033[0m\033[1m - 		Non, ça ira, le jeu est nul.	   \033[0m")
	fmt.Println("                  ")
	var input string
	fmt.Scanln(&input)
	fmt.Println("                  ")
	if input == "1" {
		p.hp = p.hpmax / 2
		fmt.Println("                  ")
		fmt.Println("\033[1mVous allez réssuciter avec 50% de vos points de vie")
		fmt.Println("	Essayez d'être plus prudent à l'avenir !\033[0m")
		fmt.Println("                  ")
		fmt.Println("                  ")
		p.MainMenu()
		return
	} else if input == "2" {
		fmt.Println("\033[1mMerci d'avoir jouer tout de même !\033[0m")
		fmt.Println("                  ")
		p.StartMenu()
		return
	} else {
		fmt.Println("                  ")
		fmt.Println("\033[1mAvez-vous bien lu vos choix possibles ??\033[0m")
		fmt.Println("                  ")
		p.Mort()
		return
	}
}

