package Projet_Red

import "fmt"

func (p *Player) AchatEquipements() {
	var input string
	fmt.Scanln(&input)
	switch input {
	case "5":
		fmt.Println("Vous avez crafté un Casque en acier, pensez à l'équiper !")
		p.AddItem("Casque en acier", 1, "+15pvmax || 20po")
		p.money -= 5

	case "6":
		fmt.Println("Vous avez crafté une Robe magique, pensez à l'équiper !")
		p.AddItem("Robe magique", 1, "+25pvmax || 30po")
		p.money -= 5

	case "7":
		fmt.Println("Vous avez crafté des Bottes en cuir, pensez à l'équiper !")
		p.AddItem("Bottes en cuir", 1, "+10pvmax || 15po")
		p.money -= 5

	}
}
