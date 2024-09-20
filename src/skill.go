package Projet_Red

import (
	"fmt"
)

func (p *Player) Spell() {
	
	var input string
	fmt.Println("Voulez-vous apprendre le sort Boule de Feu ?")
	fmt.Println("\033[96m1.\033[0m Oui")
	fmt.Println("\033[96m2.\033[0m Non")
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("             ")
		p.LearnSpell("Boule de feu")
		fmt.Println("             ")
	case "2":
		fmt.Println("             ")
		fmt.Println("Vous avez décider de ne pas apprendre le sort Boule de feu !")
		fmt.Println("             ")
	default:
		fmt.Println("Cette fois-ci veuillez choisir le bon chiffre !")
		p.Spell()
	}
}
func (p *Player) LearnSpell(name string){
	if !p.CheckSpell(name){
	p.skill = append(p.skill, name)
	fmt.Println("Vous avez appris le sort Boule de feu !")
	p.RemoveItem("Livre de sort : Boule de feu", 1)
	} else {
	fmt.Println("Vous connaissez déjà cette compétence")		
	}

}

func (p *Player) CheckSpell(name string) bool{
	for _, skill := range p.skill {
		if skill == name {
			return true
		}
	}
	return false
}

