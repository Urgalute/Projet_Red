package Projet_Red

import "fmt"

//Menu Principal
func (p *Player) MainMenu() {
	var input string
	fmt.Println("---------------------------------")
	fmt.Println("--- Menu Principal ---")
	fmt.Println("1: 	Information personnage")
	fmt.Println("2:		Inventaire")
	fmt.Println("3: 	Marchand")
	fmt.Println("4: 	Forgeron")
	fmt.Println("5: 	Combat")
	fmt.Println("0: 	Quitter")
	fmt.Print("Votre choix : ")
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("----------------")
		p.Display()
		p.MainMenu()
	case "2":
		fmt.Println("----------------")
		p.InventoryDisplay()
	case "3":
		fmt.Println("----------------")
		//Market()
	case "4":
		fmt.Println("----------------")
		//Forge()
	case "5":
		fmt.Println("----------------")
		//Combat()
	case "0":
		fmt.Println("Merci d'avoir utilisé notre jeu!")
		//Exit
	default:
		fmt.Println("----------------")
		fmt.Println("Commande inconnue, réessayez.")
		p.MainMenu()
	}
}
