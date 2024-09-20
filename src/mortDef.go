package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) MortDefinitive() {
	fmt.Println(" ")
	time.Sleep(2 * time.Second)
	fmt.Println("\033[1m	   Vous êtes mort     ")
	time.Sleep(2 * time.Second)
	fmt.Println(" ")
	fmt.Println("       Merci d'avoir joué   ")
	time.Sleep(2 * time.Second)
	fmt.Println(" ")
	fmt.Println("      Retour à l'écran titre  \033[0m")
	time.Sleep(5 * time.Second)
	fmt.Println(" ")
	p.StartMenu()
}
