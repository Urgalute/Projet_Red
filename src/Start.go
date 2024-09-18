package Projet_Red

import (
	"fmt"
	"os"
)

func (p *Player) StartMenu() {
	fmt.Println("\033[1m                           ")
	fmt.Println(`	     _       __     __   ____    _        _____ 
     /\     | |      \ \   / /  / __ \  | |      |_   _|
    /  \    | |       \ \_/ /  | |  | | | |        | |  
   / /\ \   | |        \   /   | |  | | | |        | |  
  / ____ \  | |____     | |    | |__| | | |____   _| |_ 
 /_/    \_\ |______|    |_|     \____/  |______| |_____|
                                                        `)
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("    1. Nouvelle Partie                       ")
	fmt.Println("                           ")
	fmt.Println("    2. Options                      ")
	fmt.Println("                           ")
	fmt.Println("    3. Quitter                      ")
	fmt.Println("                           ")
	fmt.Println("    4. Crédits                     ")
	fmt.Println("                       	    \033[0m")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		p.PersoInit()
		return
	case "2":
		p.MainMenu()
		return
	case "3":
		fmt.Println(" ")
		fmt.Println("Au revoir !")
		fmt.Println(" ")
		os.Exit(0)
	case "4":
		p.Credits()
		return
	default:
		fmt.Println("Les options proposées sont 1, 2, 3 et 4 ...")
		p.StartMenu()
		return
	}
}
