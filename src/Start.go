package Projet_Red

import (
	"fmt"
	"os"
)

func (p *Player) StartMenu() {
	fmt.Println("\033[1m \033[96m                         ")
	fmt.Println(`	     _       __     __   ____    _        _____ 
     /\     | |      \ \   / /  / __ \  | |      |_   _|
    /  \    | |       \ \_/ /  | |  | | | |        | |  
   / /\ \   | |        \   /   | |  | | | |        | |  
  / ____ \  | |____     | |    | |__| | | |____   _| |_ 
 /_/    \_\ |______|    |_|     \____/  |______| |_____|
                                                        `)
	fmt.Println("\033[0m\033[1m")
	fmt.Println(" ")
	fmt.Println(" ")
	fmt.Println("    \033[96m1.\033[0m Nouvelle Partie                       ")
	fmt.Println("                           ")
	fmt.Println("    \033[96m2.\033[0m Options                      ")
	fmt.Println("                           ")
	fmt.Println("    \033[96m3.\033[0m Quitter                      ")
	fmt.Println("                           ")
	fmt.Println("                           ")
	fmt.Println("                           ")
	fmt.Println("    \033[96m4.\033[0m Crédits                     ")
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
		p.StartMenu()
		return
	default:
		fmt.Println("Les options proposées sont 1, 2, 3 et 4 ...")
		p.StartMenu()
		return
	}
}
