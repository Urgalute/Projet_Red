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
	fmt.Println("\033[93mServeuse\033[0m - Qu'est ce que je vous sers ?")
	fmt.Println(" ")
	fmt.Println("\033[96m1.\033[0m- Acheter une bière locale")
	fmt.Println("\033[96m2.\033[0m- Demander des informations sur le donjon")
	fmt.Println(" ")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[94m", p.name, "\033[0m", "- Je vais vous prendre votre meilleur bière s'il vous plaît !")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[93mServeuse\033[0m - Très bien, ce sera 5 pièces d'or.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[94m", p.name, "\033[0m", "- Voici.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("\033[1m*Vous lui glissez 3 pièces d'or sur la table*\033[0m")
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[93mServeuse\033[0m - Merci, bonne dégustation")
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[1mElle s'éloigne et reprend son travail avec d'autres clients.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Vous regardez l'aubergiste qui à l'air extrêmement préoccupé,")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("probablement l'intensité du service.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Vous quittez l'auberge après avoir bu votre délicieuse bière.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Vous vous soignez de 10 pv mais perdez 10 de mana. L'alcool et la")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("magie ne font pas très bon ménage.\033[0m")
		p.hp += 10
		p.mana -= 10
		p.money -= 3
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		p.MainMenu()
		return
	case "2":
		if !p.CheckItem("Clé du Donjon Obscur") {
			fmt.Println(" ")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[94m", p.name, "\033[0m", "- Je souhaiterai avoir des informations sur le donjon pas")
			fmt.Println("loin d'ici s'il vous plaît.")
			time.Sleep(500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("\033[1m*Elle lève les yeux au ciel en soufflant*\033[0m")
			fmt.Println(" ")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[93mServeuse\033[0m - Allez voir l'aubergiste pour ça.")
			time.Sleep(500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("\033[1mElle s'éloigne et reprend son travail avec d'autres clients.")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Vous regardez l'aubergiste qui à l'air extrêmement préoccupé,")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("probablement l'intensité du service. Mais sur le conseil de la serveuse")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("vous allez à sa rencontre.")
			time.Sleep(500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("\033[94m", p.name, "\033[0m", "- Bonjour, votre serveuse m'a indiquée que vous auriez des renseignements !")
			fmt.Println("sur le donjon pas loin ?")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[93mAubergiste\033[0m - Oh, oui, c'était mon manoir il y a encore quelques mois ...")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[94m", p.name, "\033[0m", "- Que s'est-il passé ?")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[93mAubergiste\033[0m - Un monstre énorme est arrivé et à commencer à ravager les environs.")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[94m", p.name, "\033[0m", "- Je vois.")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[93mAubergiste\033[0m - Il a investi mon manoir, et il est impossible de l'en déloger, il est trop puissant.")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[94m", p.name, "\033[0m", "- Que diriez-vous que je m'en charge ?")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[93mAubergiste\033[0m - Vous feriez ça ? Très bien, je vais poser mes espoirs en vous. Voici la clé de l'entrée.")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[94m", p.name, "\033[0m", "- Merci de votre confiance, je m'y rend dès que je me suis équipé et que je suis prêt.")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[93mAubergiste\033[0m - Merci ... à bientôt j'espère ...")
			time.Sleep(500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("\033[1mVous jetez un oeil à l'aubergiste en sortant de l'auberge.")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Il a l'air moins préoccupé qu'avant, mais très légérement seulement.\033[0m")
			time.Sleep(500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("Vous obtenez la \033[93mClé du Donjon Obscur\033[0m !")
			fmt.Println(" ")
			time.Sleep(1000 * time.Millisecond)
			p.AddItem("Clé du Donjon Obscur", 1, "La clé du donjon donnée par l'aubergiste")
			p.MainMenu()
		
		}else {
			fmt.Println(" ")
			fmt.Println("\033[93mAubergiste\033[0m - Alors qu'en est-il de mon problème de donjon ?")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[94m", p.name, "\033[0m", "- j'y travaille !")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[93mAubergiste\033[0m - Revenez me voir quand vous aurez du nouveau !")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[94m", p.name, "\033[0m", "- Pas de problème !")
			time.Sleep(500 * time.Millisecond)
			fmt.Println(" ")
			p.MainMenu()
		}
	}
}
