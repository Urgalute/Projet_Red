package Projet_Red

import "fmt"

//Menu du Forgeron
func (p *Player) BlackSmithMenu() {
	fmt.Println("Bienvenue chez le forgeron !", "\n", "Contre une légère commission, ma forge est à votre disposition.")
	if p.money >= 5 {
		fmt.Println("Votre or: ", "\033[93m", p.money, "\033[0m")
		fmt.Println("\033[92mCapacité de l'inventaire: \033[0m", p.CheckQuantityInventory(), "/", p.inventorymax)
		fmt.Println("Voilà ce que je peut vous proposer.")
		fmt.Println("")
		fmt.Println("")
		if p.CheckQtyItem("Plume de corbeau") >= 1 && p.CheckQtyItem("Cuir de sanglier") >= 1 {
			fmt.Println("\033[96m1\033[0m: Casque en acier (+15 PV, 5 Pièces d'or, 1 Plume de corbeau, 1 Cuir de sanglier)")
			fmt.Println("")
		} else if p.CheckQtyItem("Plume de corbeau") < 1 || p.CheckQtyItem("Cuir de sanglier") < 1 {
			fmt.Println("Casque en acier: Il vous manque quelques composants pour le fabriquer:")
			fmt.Println("Plume de corbeau:", p.CheckQtyItem("Plume de corbeau"), "/ 1, Cuir de sanglier", p.CheckQtyItem("Cuir de sanglier"), "/ 1")
			fmt.Println("")
		}
		if p.CheckQtyItem("Fourrure de loup") >= 2 && p.CheckQtyItem("Peau de troll") >= 1 {
			fmt.Println("\033[96m2\033[0m: Robe magique (+25 PV, 5 Pièces d'or, 2 Fourrures de loup, 1 Peau de troll)")
			fmt.Println("")
		} else if p.CheckQtyItem("Fourrure de loup") < 2 || p.CheckQtyItem("Peau de troll") < 1 {
			fmt.Println("Robe magique: Il vous manque quelques composants pour le fabriquer:")
			fmt.Println("Fourrure de loup: ", p.CheckQtyItem("Fourrure de loup"), "/ 2, Peau de troll:", p.CheckQtyItem("Peau de troll"), "/ 1")
			fmt.Println("")
		}
		if p.CheckQtyItem("Fourrure de loup") >= 1 && p.CheckQtyItem("Cuir de sanglier") >= 1 {
			fmt.Println("3: Bottes en cuir (+15 PV, 5 pièces d'or, 1 Fourrure de loup, 1 Cuir de sanglier)")
			fmt.Println("")
		} else if p.CheckQtyItem("Fourrure de loup") < 1 || p.CheckQtyItem("Cuir de sanglier") < 1 {
			fmt.Println("Bottes en cuir: Il vous manque quelques composants pour le fabriquer:")
			fmt.Println("Fourrure de loup: ", p.CheckQtyItem("Fourrure de loup"), "/ 1, Cuir de sanglier: ", p.CheckQtyItem("Cuir de sangier"), "/ 1")
			fmt.Println("")
		}
		var input string
		fmt.Println("\033[96m0\033[0m: Retour")
		fmt.Scanln(&input)
		switch input {
		case "1":
			ClearTerminal()
			if p.CheckQtyItem("Plume de corbeau") >= 1 && p.CheckQtyItem("Cuir de sanglier") >= 1 {
				if !p.CheckInventory() {
					p.money -= 5
					p.AddItem("Casque en acier", 1, "Un chapeau de base pour tout aventurier (+15 PV)")
					p.RemoveItem("Plume de corbeau", 1)
					p.RemoveItem("Cuir de sanglier", 1)
					fmt.Println("Vous avez fabriqué un casque en acier avec succès!")
				} else {
					fmt.Println("\033[101mVotre inventaire est plein\033[0m, pensez à le vider ou à l'agrandir chez le marchand")
				}
			} else {
				fmt.Println("Bien essayé mais il vous manque quelques composants")
			}
		case "2":
			ClearTerminal()
			if p.CheckQtyItem("Fourrure de loup") >= 2 && p.CheckQtyItem("Peau de troll") >= 1 {
				if !p.CheckInventory() {
					p.money -= 5
					p.AddItem("Robe magique", 1, "Une tunique de base pour tout aventurier (+25 PV)")
					p.RemoveItem("Fourrure de loup", 2)
					p.RemoveItem("Peau de troll", 1)
					fmt.Println("Vous avez fabriqué une Robe magique avec succès!")
				} else {
					fmt.Println("\033[101mVotre inventaire est plein\033[0m, pensez à le vider ou à l'agrandir chez le marchand")
				}
			} else {
				fmt.Println("Bien essayé mais il vous manque quelques composants")
			}
		case "3":
			ClearTerminal()
			if p.CheckQtyItem("Fourrure de loup") >= 1 && p.CheckQtyItem("Cuir de sanglier") >= 1 {
				if !p.CheckInventory() {
					p.money -= 5
					p.AddItem("Bottes en cuir", 1, "Une botte de base pour tout aventurier (+15 PV)")
					p.RemoveItem("Fourrure de loup", 1)
					p.RemoveItem("Cuir de sanglier", 1)
					fmt.Println("Vous avez fabriqué des Bottes en cuir avec succès!")
				} else {
					fmt.Println("\033[101mVotre inventaire est plein\033[0m, pensez à le vider ou à l'agrandir chez le marchand")
				}
			} else {
				fmt.Println("Bien essayé mais il vous manque quelques composants")
			}
		case "0":
			ClearTerminal()
			fmt.Println("Retour au menu principal")
			p.MainMenu()
		default:
			ClearTerminal()
			fmt.Println("Commande inconnue. Veuillez réessayer.")
		}
		p.BlackSmithMenu()
	} else {
		fmt.Println("Désolé mais je ne travaille pas gratuitement.. Revenez quand vous aurez plus de pièces !!", "\n", "(Il vous faut \033[103m5 Pièces d'or\033[0m au minimum pour demander au forgeron de vous fabriquer un objet. Vous avez:", "\033[103m", p.money, "\033[0m")
		fmt.Println("Retour au menu principal")
		p.MainMenu()
	}
}
