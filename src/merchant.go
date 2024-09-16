package Projet_Red

import (
	"fmt"
)

var n int

func (p *Player) Market() bool {
	var input string
	fmt.Println("             ")
	fmt.Println("Place dans votre inventaire : ", p.CheckQuantityInventory(), "/", p.inventorymax)
	fmt.Println("             ")
	fmt.Println("Voici votre nombre de pièces d'or : ", p.money)
	fmt.Println("             ")
	fmt.Println("1: 	Potion de \033[31mvie\033[0m || \033[33m3po\033[0m")
	fmt.Println("2: 	Potion de \033[32mpoison\033[0m || \033[33m6po\033[0m")
	fmt.Println("3: 	Potion de \033[34mmana\033[0m || \033[33m10po\033[0m")
	if n < 1 {
		fmt.Println("4: 	Potion de \033[31mvie\033[0m GRATUITE !! (1 fois seulement)")
	}
	if n == 1 {
		fmt.Println("\033[31m4: 	Potion de vie GRATUITE !! (1 fois seulement)\033[0m")
	}
	fmt.Println("5: 	Livre de Sort : Boule de Feu || \033[33m25po\033[0m")
	fmt.Println("6: 	Fourrure de loup || \033[33m4po\033[0m")
	fmt.Println("7: 	Peau de troll 	 || \033[33m7po\033[0m")
	fmt.Println("8: 	Cuir de sanglier || \033[33m3po\033[0m")
	fmt.Println("9: 	Plume de corbeau || \033[33m1po\033[0m")
	fmt.Println("0: 	Retour")
	fmt.Print("Votre choix : ")
	fmt.Scanln(&input)
	switch input {
	case "1":
		if !p.CheckInventory() {
			fmt.Println("Vous avez assez de place dans votre inventaire !")
			p.money -= 3
			fmt.Println("             ")
			fmt.Println("Vous avez achetez une potion de santé !")
			fmt.Println("             ")
			p.AddItem("Potion de santé", 1, "(+50pv)")
			fmt.Println("             ")
			fmt.Println("Votre inventaire s'agrandit")
			p.CheckInventory()
			fmt.Println("             ")
			fmt.Println("Vous avez maintenant", p.money, "Pièces d'or")
			fmt.Println("             ")
			p.Market()
		} else {
			fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			p.Market()
		}
	case "2":
		fmt.Println("Place dans votre inventaire : ")
		if !p.CheckInventory() {
			fmt.Println("Vous avez assez de place dans votre inventaire !")
			p.money -= 6
			fmt.Println("             ")
			fmt.Println("Vous avez achetez une potion de poison !")
			fmt.Println("             ")
			p.AddItem("Potion de poison", 1, "(-10pv/s pour 3s)")
			fmt.Println("             ")
			fmt.Println("Votre inventaire s'agrandit")
			p.CheckInventory()
			fmt.Println("             ")
			fmt.Println("Vous avez maintenant", p.money, "Pièces d'or")
			fmt.Println("             ")
			p.Market()
		} else {
			fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			p.Market()
		}
	case "3":
		fmt.Println("Place dans votre inventaire : ")
		if !p.CheckInventory() {
			fmt.Println("Vous avez assez de place dans votre inventaire !")
			fmt.Println("             ")
			fmt.Println("Vous avez achetez une potion de mana !")
			fmt.Println("             ")
			p.AddItem("Potion de mana", 1, "(+25mana)")
			fmt.Println("             ")
			fmt.Println("Votre inventaire s'agrandit")
			p.CheckInventory()
			fmt.Println("             ")
			fmt.Println("Vous avez maintenant", p.money, "Pièces d'or")
			fmt.Println("             ")
			p.Market()
		} else {
			fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			p.Market()
		}
	case "4":
		fmt.Println("Place dans votre inventaire : ")
		if !p.CheckInventory() {
		if n < 1 {
			fmt.Println("             ")
			fmt.Println("Vous obtenez votre potion de santé gratuite !")
			fmt.Println("             ")
			p.AddItem("Potion de santé", 1, "(+50pv)")
			fmt.Println("             ")
			fmt.Println("Votre inventaire s'agrandit")
			p.CheckInventory()
			fmt.Println("             ")
			n++
		} else if n == 1 {
			fmt.Println("             ")
			fmt.Println("\033[31mVous avez déjà eu votre potion gratuite !\033[0m")
			fmt.Println("             ")
		} else if n > 1 {
			fmt.Println(" ")
			} else {
				fmt.Println("Vous n'avez plus de place dans votre inventaire ")
		}
	}
		p.Market()
	case "5":
		fmt.Println("Place dans votre inventaire : ")
		if !p.CheckInventory() {
			fmt.Println("Vous avez assez de place dans votre inventaire !")
			p.money -= 25
			fmt.Println("             ")
			fmt.Println("Vous avez achetez un Livre de sort : Boule de Feu !")
			fmt.Println("             ")
			p.AddItem("Livre de sort : Boule de feu", 1, "(+18 de dégats)")
			fmt.Println("             ")
			fmt.Println("Votre inventaire s'agrandit")
			p.CheckInventory()
			fmt.Println("             ")
			fmt.Println("Vous avez maintenant", p.money, "Pièces d'or")
			fmt.Println("             ")
			p.Market()
		} else {
			fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			p.Market()
		}
	case "6":
		fmt.Println("Place dans votre inventaire : ")
		if !p.CheckInventory() {
			fmt.Println("Vous avez assez de place dans votre inventaire !")
			p.money -= 4
			fmt.Println("             ")
			fmt.Println("Vous avez achetez une fourrue de loup !")
			fmt.Println("             ")
			p.AddItem("Fourrure de loup", 1, "(C'est doux)")
			fmt.Println("             ")
			fmt.Println("Votre inventaire s'agrandit")
			p.CheckInventory()
			fmt.Println("             ")
			fmt.Println("Vous avez maintenant", p.money, "Pièces d'or")
			fmt.Println("             ")
			p.Market()
		} else {
			fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			p.Market()
		}
	case "7":
		fmt.Println("Place dans votre inventaire : ")
		if !p.CheckInventory() {
			fmt.Println("Vous avez assez de place dans votre inventaire !")
			p.money -= 7
			fmt.Println("             ")
			fmt.Println("Vous avez achetez une peau de troll !")
			fmt.Println("             ")
			p.AddItem("Peau de troll", 1, "(C'est légérement poisseux, de la bave probablement)")
			fmt.Println("             ")
			fmt.Println("Votre inventaire s'agrandit")
			p.CheckInventory()
			fmt.Println("             ")
			fmt.Println("Vous avez maintenant", p.money, "Pièces d'or")
			fmt.Println("             ")
			p.Market()
		} else {
			fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			p.Market()
		}
	case "8":
		fmt.Println("Place dans votre inventaire : ")
		if !p.CheckInventory() {
			fmt.Println("Vous avez assez de place dans votre inventaire !")
			p.money -= 3
			fmt.Println("             ")
			fmt.Println("Vous avez achetez du cuir de sanglier !")
			fmt.Println("             ")
			p.AddItem("Cuir de sanglier", 1, "(Aussi robuste qu'un nain !)")
			fmt.Println("             ")
			fmt.Println("Votre inventaire s'agrandit")
			p.CheckInventory()
			fmt.Println("             ")
			fmt.Println("Vous avez maintenant", p.money, "Pièces d'or")
			fmt.Println("             ")
			p.Market()
		} else {
			fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			p.Market()
		}
	case "9":
		fmt.Println("Place dans votre inventaire : ")
		if !p.CheckInventory() {
			fmt.Println("Vous avez assez de place dans votre inventaire !")
			p.money -= 1
			fmt.Println("             ")
			fmt.Println("Vous avez achetez une plume de corbeau !")
			fmt.Println("             ")
			p.AddItem("Plume de corbeau", 1, "(Le chatouilleur 2000 ultime)")
			fmt.Println("             ")
			fmt.Println("Votre inventaire s'agrandit")
			p.CheckInventory()
			fmt.Println("             ")
			fmt.Println("Vous avez maintenant", p.money, "Pièces d'or")
			fmt.Println("             ")
			p.Market()
		} else {
			fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			p.Market()
		}
	case "0":
		p.MainMenu()
	default:
		fmt.Println("Cette fois-ci veuillez choisir le bon chiffre !")
		p.Market()
	}
	return true
}
