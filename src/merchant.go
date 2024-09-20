package Projet_Red

import (
	"fmt"
)

var n int

// Menu marchand
func (p *Player) Market() {
	var input string
	fmt.Println("             ")
	fmt.Println("Place dans votre inventaire : ", p.CheckQuantityInventory(), "/", p.inventorymax)
	fmt.Println("Voici votre nombre de pièces d'or : ", p.money)
	fmt.Println("             ")
	fmt.Println("             ")
	fmt.Println("\033[96m1\033[0m: 	Potion de \033[31mvie\033[0m || \033[33m3po\033[0m")
	fmt.Println("\033[96m2\033[0m: 	Potion de \033[32mpoison\033[0m || \033[33m6po\033[0m")
	fmt.Println("\033[96m3\033[0m: 	Potion de \033[34mmana\033[0m || \033[33m10po\033[0m")
	if n < 1 {
		fmt.Println("\033[96m4\033[0m: 	Potion de \033[31mvie\033[0m GRATUITE !! (1 fois seulement)")
	}
	if n == 1 {
		fmt.Println("\033[31m4: 	Potion de vie GRATUITE !! (1 fois seulement)\033[0m")
	}
	fmt.Println("\033[96m5\033[0m: 	Livre de Sort : Boule de Feu || \033[33m25po\033[0m")
	fmt.Println("\033[96m6\033[0m: 	Fourrure de loup || \033[33m4po\033[0m")
	fmt.Println("\033[96m7\033[0m: 	Peau de troll 	 || \033[33m7po\033[0m")
	fmt.Println("\033[96m8\033[0m: 	Cuir de sanglier || \033[33m3po\033[0m")
	fmt.Println("\033[96m9\033[0m: 	Plume de corbeau || \033[33m1po\033[0m")
	fmt.Println("\033[96m10\033[0m:	Augmentation d'inventaire || \033[33m30po\033[0m")
	fmt.Println("\033[96m0\033[0m: 	Retour")
	fmt.Print("Votre choix : ")
	fmt.Scanln(&input)
	switch input {
	case "1":
		ClearTerminal()
		if p.Moneyy(3) {
			if !p.CheckInventory() {
				fmt.Println("             ")
				fmt.Println("Vous avez achetez une potion de santé !")
				p.AddItem("Potion de santé", 1, "(+50pv)")
				p.money -= 3
			} else {
				fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			}
		} else {
			fmt.Println("             ")
			fmt.Println("Vous n'avez pas assez d'or")
		}
		p.Market()
		return
	case "2":
		ClearTerminal()
		if p.Moneyy(6) {
			if !p.CheckInventory() {
				fmt.Println("             ")
				fmt.Println("Vous avez achetez une potion de poison !")
				p.AddItem("Potion de poison", 1, "(-10pv/s pour 3s)")
				p.money -= 6
			} else {
				fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			}
		} else {
			fmt.Println("             ")
			fmt.Println("Vous n'avez pas assez d'or")
		}
		p.Market()
		return
	case "3":
		ClearTerminal()
		if p.Moneyy(10) {
			if !p.CheckInventory() {
				fmt.Println("             ")
				fmt.Println("Vous avez achetez une potion de mana !")
				p.AddItem("Potion de mana", 1, "(+25mana)")
				p.money -= 10
			} else {
				fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			}
		} else {
			fmt.Println("             ")
			fmt.Println("Vous n'avez pas assez d'or")
		}
		p.Market()
		return
	case "4":
		ClearTerminal()
		if !p.CheckInventory() {
			if n < 1 {
				fmt.Println("             ")
				fmt.Println("Vous obtenez votre potion de santé gratuite !")
				p.AddItem("Potion de santé", 1, "(+50pv)")
				p.CheckInventory()
				fmt.Println("             ")
				n++
			} else if n >= 1 {
				fmt.Println("             ")
				fmt.Println("\033[31mVous avez déjà eu votre potion gratuite !\033[0m")
				fmt.Println("             ")
			}
		} else {
			fmt.Println("             ")
			fmt.Println("Vous n'avez plus de place dans votre inventaire ")
		}
		p.Market()
		return
	case "5":
		ClearTerminal()
		if p.Moneyy(25) {
			if !p.CheckInventory() {
				fmt.Println("             ")
				fmt.Println("Vous avez achetez un Livre de sort : Boule de Feu !")
				fmt.Println("             ")
				p.AddItem("Livre de sort : Boule de feu", 1, "(+18 de dégats)")
				p.money -= 25
			} else {
				fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			}
		} else {
			fmt.Println("             ")
			fmt.Println("Vous n'avez pas assez d'or")
		}
		p.Market()
		return
	case "6":
		ClearTerminal()
		if p.Moneyy(4) {
			if !p.CheckInventory() {
				fmt.Println("             ")
				fmt.Println("Vous avez achetez une fourrue de loup !")
				fmt.Println("             ")
				p.AddItem("Fourrure de loup", 1, "(C'est doux)")
				p.money -= 4
			} else {
				fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			}
		} else {
			fmt.Println("             ")
			fmt.Println("Vous n'avez pas assez d'or")
		}
		p.Market()
		return
	case "7":
		ClearTerminal()
		if p.Moneyy(7) {
			if !p.CheckInventory() {
				fmt.Println("             ")
				fmt.Println("Vous avez achetez une peau de troll !")
				fmt.Println("             ")
				p.AddItem("Peau de troll", 1, "(C'est légérement poisseux, de la bave probablement)")
				p.money -= 7
			} else {
				fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			}
		} else {
			fmt.Println("             ")
			fmt.Println("Vous n'avez pas assez d'or")
		}
		p.Market()
		return
	case "8":
		ClearTerminal()
		if p.Moneyy(3) {
			if !p.CheckInventory() {
				fmt.Println("             ")
				fmt.Println("Vous avez achetez du cuir de sanglier !")
				fmt.Println("             ")
				p.AddItem("Cuir de sanglier", 1, "(Aussi robuste qu'un nain !)")
				p.money -= 3
			} else {
				fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			}
		} else {
			fmt.Println("             ")
			fmt.Println("Vous n'avez pas assez d'or")
		}
		p.Market()
		return
	case "9":
		ClearTerminal()
		if p.Moneyy(1) {
			if !p.CheckInventory() {
				fmt.Println("             ")
				fmt.Println("Vous avez achetez une plume de corbeau !")
				fmt.Println("             ")
				p.AddItem("Plume de corbeau", 1, "(Le chatouilleur 2000 ultime)")
				p.money -= 1
			} else {
				fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			}
		} else {
			fmt.Println("             ")
			fmt.Println("Vous n'avez pas assez d'or")
		}
		p.Market()
		return
	case "10":
		ClearTerminal()
		if p.Moneyy(30) {
			if !p.CheckInventory() {
				fmt.Println("             ")
				fmt.Println("Vous avez achetez une Augmentation d'inventaire !")
				fmt.Println("             ")
				p.AddItem("Augmentation d'inventaire", 1, "(+10 de taille d'inventaire)")
				p.money -= 30
			} else {
				fmt.Println("Vous n'avez plus de place dans votre inventaire ")
			}
		} else {
			fmt.Println("             ")
			fmt.Println("Vous n'avez pas assez d'or")
		}
		p.Market()
		return
	case "0":
		ClearTerminal()
		p.MainMenu()
		return
	default:
		ClearTerminal()
		fmt.Println("Cette fois-ci veuillez choisir le bon chiffre !")
		p.Market()
		return
	}
}

// Augmente l'inventaire de 10 objets supplémentaires
func (p *Player) UpgradeInventorySlot() {
	if p.inventorymax >= 40 {
		fmt.Println("\033[101mVous ne pouvez plus augmenter votre Inventaire ! \033[0m")
	} else {
		p.inventorymax += 10
		fmt.Println("Bien joué vous avez augmenté votre Inventaire de 10 places supplémentaires !")
		p.RemoveItem("Augmentation d'inventaire", 1)
	}
}

// Check si on a assez d'or

	func (p *Player) Moneyy(cost int) bool {
		return p.money >= cost
	}
