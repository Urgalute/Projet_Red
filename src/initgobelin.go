package Projet_Red

import "fmt"

func (p *Player) InitGobelin() {
	fmt.Println(" ")
	fmt.Println("Nom du monstre :", p.mname)
	fmt.Println("Pv max du monstre :", p.mhpmax)
	fmt.Println("Pv actuels du monstre :", p.mhp)
	fmt.Println("Dégats du monstre :", p.mdamage)
	fmt.Println(" ")
}