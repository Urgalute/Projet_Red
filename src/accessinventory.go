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
		time.Sleep(2 * time.Second)
		fmt.Println("Retour au menu principal")
		time.Sleep(2 * time.Second)
		p.MainMenu()
		return
	} else {
		fmt.Println("Place dans votre inventaire : ", p.CheckQuantityInventory(), "/", p.inventorymax)
		fmt.Println("                  ")
		for i, objet := range p.inventory {
			fmt.Println(i+1, "- 	", objet.name, "x", objet.quantity, "	", objet.description)
		}
	}
	fmt.Println("--------------------------------------------------------")
	fmt.Println("       Tapez 0 pour revenir au menu principal        ")
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
	} else if input < len(p.inventory)+1 {
		fmt.Println("                  ")
		selectedItem = items[input-1]
		switch selectedItem.name {
		case "Potion de poison":
			p.Poison()
			fmt.Println("Vous avez bu une potion de poison.")
			fmt.Println("                  ")
			p.InventoryDisplay()
		case "Potion de santé":
			p.TakePot()
			fmt.Println("Vous avez bu une potion de santé")
			fmt.Println("                  ")
			p.InventoryDisplay()
		case "Livre de sort : Boule de feu":
			p.Spell()
			p.InventoryDisplay()
		case "Casque en acier":
			p.EquipGear("Casque en acier", 1, "+15pvmax")
			p.RemoveItem("Casque en acier", 1)
			p.InventoryDisplay()
		case "Robe magique":		
			p.EquipGear("Robe magique", 1, "+25pvmax")
			p.RemoveItem("Robe magique", 1)
			p.InventoryDisplay()
		case "Bottes en cuir":
			p.EquipGear("Bottes en cuir", 1, "+10pvmax")
			p.RemoveItem("Bottes en cuir", 1)
			p.InventoryDisplay()
		default:
			p.InventoryDisplay()
				return
		}
	}else {
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
	fmt.Println(count, "/", p.inventorymax)
	return false
}

func (p *Player) CheckQuantityInventory() int {
	count := 0
	for _, items := range p.inventory {
		count = count + items.quantity
	}
	return count
}
