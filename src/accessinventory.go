package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) InventoryDisplay() {
	fmt.Println("=========================================")
	fmt.Println("-----------------------------------------")
	fmt.Println("||        I N V E N T A I R E          ||")
	fmt.Println("-----------------------------------------")
	fmt.Println("=========================================")
	fmt.Println("                                         ")
	if len(p.inventory) == 0 {
		fmt.Println("Votre inventaire est vide !")
		time.Sleep(3 * time.Second)
		p.MainMenu()
	} else {
			for i, objet := range p.inventory {
				fmt.Println( i+1, "- ", objet.name, "x", objet.quantity)
			}
		}
	}
func (p *Player) MenuInventaire() {
	var input string

	fmt.Scanln(&input)
	switch input {
	case "1":
	}
}
