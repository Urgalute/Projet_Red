package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) Donjon() {
	fmt.Println(" ")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("\033[1m\033[104mVous arrivez devant la grande porte du donjon\033[0m\033[1m. Vous essayez de la pousser,")
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
	fmt.Println(" ")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("\033[1m\033[104mVous arrivez devant la grande porte du donjon\033[0m\033[1m. Vous essayez de la pousser,")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("mais elle est bien fermée à clé. Vous apercevez une serrure en son centre.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println(" ")
	if p.CheckItem("Clé du Donjon Obscur") {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[1mVous avez la clé donnée par l'aubergiste. Vous ouvrez la porte en la poussant,")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("elle grince extrêmement fort, et s'ouvre difficilement. Vous pénétrez le donjon,")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Et c'est ici que s'arrête votre aventure pour cette petite introduction.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Merci beaucoup d'avoir joué, nous espérons que cette petite aventure vous a")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("plus. N'hésitez pas à wishlist le jeu sur steam, et à le précommander")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("pour seulement 10€. Il n'en est encore qu'aux balbutiements de sa splendeur future.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Le mot des développeurs :")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Lilian : Bien.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Yoan : Laissez moi tranquille.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Alexis : Achète !")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		time.Sleep(3 * time.Millisecond)
		fmt.Println("Au revoir et à bientôt !!\033[0m")
		time.Sleep(3 * time.Second)
		p.StartMenu()
		return

	} else if !p.CheckItem("Clé du Donjon Obscur") {
		fmt.Println("\033[1mVous n'avez pas la clé, peut être devriez vous aller vous renseigner")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("sur le donjon à l'auberge.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("Vous retournez en ville")
		fmt.Println(" ")
		p.MainMenu()
		return

	}
}
