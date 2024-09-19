package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) TransitDonjon() {
	fmt.Println("\033[1mVous sortez de la ville, en regardant au loin vous voyez un panneau")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("indiquant le donjon obscur et la ville. Vous pourriez aller au donjon, voir de quoi")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("il retourne. Ou vous pouvez toujours retourner en ville.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println(" ")
	fmt.Println("     ____________________________________________________________________     ")
	fmt.Println("<== || \033[34mDirection le donjon |1\033[0m|||\033[35m2| Retourner en ville\033[0m || ==>")
	fmt.Println("     ====================================================================     ")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("\033[1mVous progressez tranquillement sur le chemin menant au donjon,")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("croisant un deuxième panneau indicateur, vous continuez à suivre")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("les indications de ce panneau, et vous apercevez enfin le donjon")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("poindre à l'horizon. Il est collossal, et sa grande porte noire")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("n'invite pas à être ouverte.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Vous approchez de la porte et remarquez qu'il faut une clé.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Vous remarquez aussi des traces de sang sur la poignée et au sol.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("L'ambiance de ce donjon s'annonce festive !\033[0m")
		time.Sleep(500 * time.Millisecond)
		p.EntreeDonjon()
	case "2":
		p.MainMenu()
	default:
		fmt.Println("Il n'y a rien du tout à part ces deux chemins et de la forêt")
		p.TransitDonjon()
	}
}
