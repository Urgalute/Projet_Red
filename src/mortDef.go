package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) MortDefinitive() {
	fmt.Println(" ")
	time.Sleep(2 * time.Second)
	fmt.Println("	   Vous êtes mort     ")
	time.Sleep(2 * time.Second)
	fmt.Println(" ")
	fmt.Println("       Merci d'avoir joué   ")
	time.Sleep(2 * time.Second)
	fmt.Println(" ")
	fmt.Println("      Retour à l'écran titre  ")
	time.Sleep(5 * time.Second)
	fmt.Println(" ")
	p.StartMenu()
}
