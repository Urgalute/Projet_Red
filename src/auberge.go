package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) Auberge() {
	fmt.Println(" ")
	fmt.Println("\033[1mVous pénétrez dans l'auberge. L'ambiance y est chaleureuse,")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("des gens dansent sur le rythme d'un groupe de ménéstrels avec un nom étrange,")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Futur Luth, original. Les discussions se veulent plus ou moins vive et mouvementé en ")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("fonction des tables. Vous en apercevez une de libre vers le fond.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Vous partez vous asseoir.")
	fmt.Println(" ")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Une serveuse vient à votre rencontre")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Serveuse - Qu'est ce que je vous sers ?")
	fmt.Println(" ")
	fmt.Println("1.- Acheter une bière locale")
	fmt.Println("2.- Demander des informations sur le donjon")
	fmt.Println(" ")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		time.Sleep(500 * time.Millisecond)
		fmt.Println(p.name, "- Je vais vous prendre votre meilleur bière s'il vous plaît !")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Serveuse - Très bien, ce sera 5 pièces d'or.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(p.name, "- Voici.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("\033[1m*Vous lui glissez 5 pièces d'or sur la table*\033[0m")
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Serveuse - Merci, bonne dégustation")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[1mElle s'éloigne et reprend son travail avec d'autres clients.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Vous regardez l'aubergiste qui à l'air extrêmement préoccupé,")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("probablement l'intensité du service.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("fonction des tables. Vous en apercevez une de libre vers le fond.")
		time.Sleep(500 * time.Millisecond)
	case "2":

	}
}
