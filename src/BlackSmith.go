package Projet_Red

import "fmt"

func (p *Player) BlackSmithMenu() {
	fmt.Println("Bienvenue chez le forgeron !", "\n", "Contre une légère commission, ma forge est à votre disposition.")
	if p.money >= 5 {
		fmt.Println("Votre or: ", p.money)
		fmt.Println("Capacité de l'inventaire: ")
		fmt.Println("Voilà ce que je peut vous proposer.")
		if  p.CheckQtyItem("Plume de corbeau") >= 1 && p.CheckQtyItem("Cuir de sanglier") >= 1{
			fmt.Println("1: Chapeau de l'aventurier (+15 PV, 5 Pièces d'or, 1 Plume de corbeau, 1 Cuir de sanglier)")
		} else if p.CheckQtyItem("Plume de corbeau") < 1 || p.CheckQtyItem("Cuir de sanglier") < 1 {
				fmt.Println("Chapeau de l'aventurier: Il vous manque quelques composants pour le fabriquer: Plume de corbeau:", p.CheckQtyItem("Plume de corbeau"), "/ 1, Cuir de sanglier", p.CheckQtyItem("Cuir de sanglier"),"/ 1")
			}
		if p.CheckQtyItem("Fourrure de loup") >= 2 && p.CheckQtyItem("Peau de troll") >= 1 {
			fmt.Println("2: Tunique de l'aventurier (+25 PV, 5 Pièces d'or, 2 Fourrures de loup, 1 Peau de troll)")
		} else if p.CheckQtyItem("Fourrure de loup") < 2 || p.CheckQtyItem("Peau de troll") < 1 {
			fmt.Println("Tunique de l'aventurier: Il vous manque quelques composants pour le fabriquer: Fourrure de loup: ", p.CheckQtyItem("Fourrure de loup"), "/ 2, Peau de troll:", p.CheckQtyItem("Peau de troll"), "/ 1")
		}
		if p.CheckQtyItem("Fourrure de loup") >= 1 && p.CheckQtyItem("Cuir de sanglier") >= 1 {
			fmt.Println("3: Bottes de l'aventurier (+15 PV, 5 pièces d'or, 1 Fourrure de loup, 1 Cuir de sanglier)")
		} else if p.CheckQtyItem("Fourrure de loup") < 1 || p.CheckQtyItem("Cuir de sanglier") >= 1 {
			fmt.Println(" Bottes de l'aventurier: Il vous manque quelques composants pour le fabriquer: Fourrure de loup: ", p.CheckQtyItem("Fourrure de loup"), "/ 1, Cuir de sanglier: ", p.CheckQtyItem("Cuir de sangier"), "/ 1")
		}
		var input string
		fmt.Println("0: Retour")
		fmt.Scanln(&input)
		switch input {
		case "1":
			if p.CheckQtyItem("Plume de corbeau") >= 1 && p.CheckQtyItem("Cuir de sanglier") >= 1 {
				p.money -= 5
                p.AddItem("Chapeau de l'aventurier", 1, "Un chapeau de base pour tout aventurier (+15 PV)")
				p.RemoveItem("Plume de corbeau", 1)
				p.RemoveItem("Cuir de sanglier", 1)   
                fmt.Println("Vous avez fabriqué un chapeau de l'aventurier avec succès!")
                p.BlackSmithMenu()
			}
		}
	} else {
		fmt.Println("Désolé mais je ne travaille pas gratuitement..", "\n", "(Il vous faut 5 Pièces d'or au minimum pour demander au forgeron de vous fabriquer un objet. Vous avez:", p.money, ")")
		fmt.Println("Retour au menu principal")
		p.MainMenu()
	}
}