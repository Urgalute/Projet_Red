package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) Donjon() {
	fmt.Println(" ")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("\033[1mVous arrivez devant la grande porte du donjon. Vous essayez de la pousser,")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("mais elle est bien fermée à clé. Vous apercevez une serrure en son centre.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Il vous faudra trouver la clé pour y entrer.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Vous repartez donc sur le chemin longé par un ruisseau, dans la deuxième direction qu'indiquait le panneau.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("La ville d'Alyoli\033[0m")
	fmt.Println(" ")
	p.Ruisseau()
}

func (p *Player) EntreeDonjon() {

}
