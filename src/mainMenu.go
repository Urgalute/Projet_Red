package Projet_Red

import (
	"fmt"
)

// Menu Principal
func (p *Player) MainMenu() {
	var input string
	fmt.Println("==============================")
	fmt.Println("=========  \033[96mA L Y O L I\033[0m ========")
	fmt.Println("==============================")
	fmt.Println(" ")
	fmt.Println("\033[96m1\033[0m: 	Information personnage")
	fmt.Println("\033[96m2\033[0m: 	Inventaire")
	fmt.Println("\033[96m3\033[0m: 	Marchand")
	fmt.Println("\033[96m4\033[0m: 	Forgeron")
	fmt.Println("\033[96m5\033[0m: 	Equipements")
	fmt.Println("\033[96m6\033[0m: 	Entraînement au combat")
	fmt.Println("\033[96m7\033[0m: 	Auberge")
	fmt.Println(" ")
	fmt.Println("\033[96m8\033[0m:     Quitter la ville")
	fmt.Println(" ")
<<<<<<< HEAD
	fmt.Println("\033[96m0\033[0m: 	Quitter")
=======
	fmt.Println("0:     Quitter le jeu")
>>>>>>> 87f608dc522a2f057ccbf0737e46ff014cc2e93d
	fmt.Print("Votre choix : ")
	fmt.Scanln(&input)
	switch input {
	case "1":
		ClearTerminal()
		fmt.Println("----------------")
		p.Display()
		p.MainMenu()
		return
	case "2":
		ClearTerminal()
		fmt.Println("----------------")
		p.InventoryDisplay()
		return
	case "3":
		ClearTerminal()
		fmt.Println("----------------")
		fmt.Println("Bienvenue chez le Marchand ! ")
		p.Market()
		return
	case "4":
		ClearTerminal()
		fmt.Println("----------------")
		p.BlackSmithMenu()
		p.MainMenu()
		return
	case "5":
		ClearTerminal()
		fmt.Println("----------------")
		p.GearSlot()
		return
	case "6":
		ClearTerminal()
		fmt.Println("----------------")
		p.Entrainement()
		return
	case "7":
		ClearTerminal()
		fmt.Println("----------------")
		p.Auberge()
		return
	case "8":
		ClearTerminal()
		fmt.Println("----------------")
		p.TransitDonjon()
		return
	case "0":
		ClearTerminal()
		var input string
		fmt.Println("----------------")
		fmt.Println(" ")
		fmt.Println("Êtes vous sûr de vouloir quitter. Il n'y a pas de sauvegarde !")
		fmt.Println("\033[96m1.\033[0m-	Oui")
		fmt.Println("\033[96m2.\033[0m-	Non")
		fmt.Println(" ")
		fmt.Scanln(&input)
		switch input {
		case "1":
			fmt.Println("Merci d'avoir joué !")
			p.StartMenu()
			return
		case "2":
			fmt.Println("On y retourne !")
			p.MainMenu()
			return
		default:
			fmt.Println("----------------")
			fmt.Println("Commande inconnue, réessayez.")
			p.MainMenu()
			return
		}
	default:
		ClearTerminal()
		fmt.Println("----------------")
		fmt.Println("Commande inconnue, réessayez.")
		p.MainMenu()
		return
	}
}
