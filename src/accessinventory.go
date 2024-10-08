package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) InventoryDisplay() {
	fmt.Println("=========================================================")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("||                 I N V E N T A I R E                 ||")
	fmt.Println("---------------------------------------------------------")
	fmt.Println("=========================================================")
	fmt.Println("                                                         ")
	fmt.Println(" Tapez le chiffre correspondant pour utiliser votre objet")
	fmt.Println("                                                         ")
	if len(p.inventory) == 0 {
		fmt.Println("Votre inventaire est vide !")
		time.Sleep(1 * time.Second)
		fmt.Println("  ")
		fmt.Println("Retour au menu principal")
		time.Sleep(1 * time.Second)
		p.MainMenu()
		return
	} else {
		fmt.Println("\033[92mPlace dans votre inventaire :\033[0m ", p.CheckQuantityInventory(), "/", p.inventorymax)
		fmt.Println("Vos Pièces d'or: ","\033[93m" ,p.money, "\033[93m")
		fmt.Println("                  ")
		for i, objet := range p.inventory {
			fmt.Println("\033[96m",i+1,"\033[0m", "- 	", objet.name, "x", objet.quantity, "	", objet.description)
		}
	}
	fmt.Println("--------------------------------------------------------")
	fmt.Println("       Tapez \033[96m0\033[0m pour revenir au menu principal        ")
	fmt.Println("________________________________________________________")
	fmt.Println("--------------------------------------------------------")
	p.MenuInventaire()
}
func (p *Player) MenuInventaire() {
	var input int
	var selectedItem Inventory
	var items []Inventory = p.inventory
	fmt.Println("                  ")
	fmt.Scanln(&input)
	if input == 0 {
		p.MainMenu()
		fmt.Println("                  ")
		return
	} else if input < len(p.inventory)+1 {
		fmt.Println("                  ")
		selectedItem = items[input-1]
		switch selectedItem.name {
		case "Potion de poison":
			ClearTerminal()
			p.Poison()
			fmt.Println("                  ")
			p.InventoryDisplay()
			return
		case "Potion de santé":
			ClearTerminal()
			p.TakePot()
			fmt.Println("                  ")
			p.InventoryDisplay()
			return
		case "Potion de mana":
			ClearTerminal()
			p.Mana()
			fmt.Println("                  ")
			p.InventoryDisplay()
		case "Livre de sort : Boule de feu":
			ClearTerminal()
			p.Spell()
			p.InventoryDisplay()
			return
		case "Casque en acier":
			ClearTerminal()
			if p.EquipGear("Casque en acier", 1, "+15pvmax") {
				fmt.Println("Vous avez déja un équipement sur cet emplacement")
			} else {
				p.RemoveItem("Casque en acier", 1)
			}
			p.InventoryDisplay()
			return
		case "Robe magique":
			ClearTerminal()
			if p.EquipGear("Robe magique", 1, "+25pvmax") {
				fmt.Println("Vous avez déja un équipement sur cet emplacement")
			} else {
				p.RemoveItem("Robe magique", 1)
			}
			p.InventoryDisplay()
			return
		case "Bottes en cuir":
			ClearTerminal()
			if p.EquipGear("Bottes en cuir", 1, "+10pvmax") {
				fmt.Println("Vous avez déja un équipement sur cet emplacement")
			} else {
				p.RemoveItem("Bottes en cuir", 1)
			}
			p.InventoryDisplay()
			return
		case "Augmentation d'inventaire":
			p.UpgradeInventorySlot()
			p.InventoryDisplay()
		default:
			p.InventoryDisplay()
			return
		}
	} else {
		ClearTerminal()
		fmt.Println("Cette fois-ci veuillez choisir le bon chiffre !")
		p.InventoryDisplay()
		return
	}
}

// Y faut verifier que la quantité dans l'inventaire sois inférieur a 10 pour mettre des objets sinon je doit signaler que l'inventaire est plein !
func (p *Player) CheckInventory() bool {
	count := 0
	for _, items := range p.inventory {
		count = count + items.quantity
		if count >= p.inventorymax {
			return true
		}
	}
	return false
}

// Check et renvoie le nombre d'item dans l'inventaire
func (p *Player) CheckQuantityInventory() int {
	count := 0
	for _, items := range p.inventory {
		count = count + items.quantity
	}
	return count
}
