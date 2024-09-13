package Projet_Red

import (
	"fmt"
)
var n int
func (p *Player) Market() {
	var input string
	fmt.Println("1: 	Potion de \033[31mvie\033[0m")
	fmt.Println("2: 	Potion de \033[32mpoison\033[0m")
	fmt.Println("2: 	Potion de \033[34mmana\033[0m")
	if n < 1 {
		fmt.Println("3: 	Potion de vie GRATUITE !! (1 fois seulement)")
	}
	if n == 1 {
		fmt.Println("\033[31m3: 	Potion de vie GRATUITE !! (1 fois seulement)\033[0m")
	}
	fmt.Println("0: 	Retour")
	fmt.Print("Votre choix : ")
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("             ")
		fmt.Println("Vous avez achetez une potion de vie !")
		fmt.Println("             ")
		p.AddItem("Potion de santé", 1, "(+50pv)")
		p.Market()
	case "2":
		fmt.Println("             ")
		fmt.Println("Vous avez achetez une potion de poison !")
		fmt.Println("             ")
		p.AddItem("Potion de poison", 1, "(-10pv/s pour 3s)")
		p.Market()
	case "3" :
		fmt.Println("             ")
		fmt.Println("Vous avez achetez une potion de mana !")
		fmt.Println("             ")
		p.AddItem("Potion de mana", 1, "(+35 mana)")
		p.Market()
	case "4":
			if n < 1 {
				fmt.Println("             ")
				fmt.Println("Vous obtenez votre potion de vie gratuite !")
				fmt.Println("             ")
				p.AddItem("Potion de santé", 1, "(+50pv)")
				n++
			} else if n == 1 {
				fmt.Println("             ")
				fmt.Println("\033[31mVous avez déjà eu votre potion gratuite !\033[0m")
				fmt.Println("             ")
			} else if n > 1 {
				fmt.Println(" ")
			}
		p.Market()
	case "0":
		p.MainMenu()
	default:
		fmt.Println("Cette fois-ci veuillez choisir le bon chiffre !")
		p.Market()
	}
}
