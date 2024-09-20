package Projet_Red

import (
	"fmt"
	"time"
)

func (p *Player) DebutAventure() {
	fmt.Println("\033[1mVous vous réveillez doucement ...")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Depuis combien de temps dormez vous ? ")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Où êtes-vous ? ... ")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("*Vous fouillez vos poches, rien à part ",p.money,"*")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Vous regardez autour de vous, vous êtes dans une petite grotte")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("et en regardant par la sortie, des arbres à perte de vue.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Bien qu'un petit chemin semble se dessiner sur la droite.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Le soleil est haut dans le ciel.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Le choix est vôtre.\033[0m")
	fmt.Println(" ")
	fmt.Println(" ")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("     ____________________________________________________________________     ")
	fmt.Println("<== || \033[34mA gauche à travers la forêt |1\033[0m|||\033[35m2| A droite par le petit chemin\033[0m || ==>")
	fmt.Println("     ====================================================================     ")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		p.TraversForet()
	case "2":
		p.CheminForet()
	default:
		fmt.Println("Il n'y a rien du tout à part la grotte de ce côté")
		p.DebutAventure()
	}
}

func (p *Player) CheminForet() {
	fmt.Println("\033[1mVous suivez le chemin prudemment, qui sait ce qui habite cette forêt.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Après une heure de marche, et une frayeur à l'apparition d'un sanglier,")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("vous arrivez dans une clairière, longée par un petit ruisseau.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Vous pouvez continuer sur le chemin,")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("ou bien vous reposer un instant dans ce calme apparent.\033[0m")
	fmt.Println(" ")
	time.Sleep(500 * time.Millisecond)
	p.MenuAventure()
	fmt.Println("\033[1mVous pouvez décider de continuer sur le chemin, ou bien suivre le ruisseau\033[0m")
	time.Sleep(500 * time.Millisecond)
	fmt.Println(" ")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("     ______________________________________________________     ")
	fmt.Println("<== || \033[34mContinuer sur le chemin |1\033[0m|||\033[35m2| Suivre le ruisseau\033[0m || ==>")
	fmt.Println("     ======================================================     ")
	time.Sleep(500 * time.Millisecond)
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		p.SuiteChemin()
	case "2":
		p.Ruisseau()
	default:
		fmt.Println("Suivez les instructions du narrateur !")
		p.CheminForet()
		return
	}
}

func (p *Player) TraversForet() {
	fmt.Println("\033[1mVous avancez péniblement au travers de la forêt,")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("peut être auriez vous dû suivre ce petit chemin qui avait l'air bien plus pratiquable.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Malgré votre marche lente, vous ne remarquez pas un trou conséquent,")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("vous risquez de vous tordre la cheville si vous ne l'évitez pas !")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("  ")
	fmt.Println("Lancez le dé (Entrez D) et obtenez un 12 ou plus.")
	fmt.Println("  ")
	var inputD string
	fmt.Scanln(&inputD)
	switch inputD {
	case "D":
		var résultat int
		résultat = p.JetdeDé()
		fmt.Println("\033[1mVous avez fait un \033[0m", résultat)
		fmt.Println("  ")
		if résultat >= 12 {
			fmt.Println("\033[1mVous vous rattrapez au dernier moment et évitez le trou")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Vous continuez votre route en surveillant plus intensément")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("là où vous posez vos pieds.\033[0m")
			time.Sleep(500 * time.Millisecond)
		} else {
			fmt.Println("\033[1mVous marchez en plein dans le trou, et entendez un craquement au niveau de votre cheville")
			fmt.Println("  ")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[91mVous perdez 5 PV\033[0m")
			fmt.Println("  ")
			time.Sleep(500 * time.Millisecond)
			p.hp -= 10
			fmt.Println("\033[1mVous repartez en boitillant, en espérant trouver une ville au plus vite.\033[0m")
		}
	default:
		fmt.Println("  ")
		fmt.Println("Suivez les instructions du narrateur !")
		fmt.Println("  ")
		p.StartMenu()
		return
	}
	fmt.Println("\033[1mVous marchez maintenant depuis une bonne heure dans cette forêt dense,")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("et rien ne vous laisse penser que vous en sortirez bientôt.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Quand, au détour d'un arbre, vous apercevez une vieille massure.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Vu son état, elle est probablement vide, mais vous pouvez tout de même")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("aller la voir si vous le souhaitez.\033[0m")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("     ____________________________________________________________     ")
	fmt.Println("<== || \033[34mContinuer dans la forêt |1\033[0m|||\033[35m2| Aller voir cette massure\033[0m || ==>")
	fmt.Println("     ============================================================     ")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		p.SuiteForet()
	case "2":
		p.Massure()
	default:
		fmt.Println("  ")
		fmt.Println("Suivez les instructions du narrateur !")
		fmt.Println("  ")
		p.StartMenu()
		return
	}
}

func (p *Player) SuiteForet() {
	fmt.Println("\033[1mLa forêt se fait de plus en plus dense, les arbres commencent à")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("obstruer la lumière du soleil. Plus vous progressez, et plus la forêt")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("devient inquiétante.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Soudainement vous entendez le hurlement d'un loup dans le lointain.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Il va falloir se mettre à courir, et vite, impossible de savoir si ils")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("sont déjà sur votre piste ou non.")
	fmt.Println(" ")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Vous courez droit devant sans regardez derrière vous, votre souffle")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("se fait de plus en plus court.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Après quelques minutes, vous arrivez devant un pont délabrée traversant une petite crevasse.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Vous entendez les hurlements des loups plus proche, ils sont bel et bien sur vos traces !")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Emprunter le pont serait plus rapide, mais risqué, faire le tour sera plus long, mais sûr.\033[0m")
	time.Sleep(500 * time.Millisecond)
	fmt.Println(" ")
	fmt.Println("     ________________________________________________     ")
	fmt.Println("<== || \033[34mPasser par le pont ! |1\033[0m|||\033[35m2| Faire le tour !\033[0m || ==>")
	fmt.Println("     ================================================     ")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("\033[94m", p.name, "\033[0m", "\033[1m - Allez, c'est parti !")
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Vous posez un pied sur le pont, la goutte de sueur sur votre front")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("se faisant de plus en plus proéminante.\033[0m")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("Lancez le dé (Entrez D), votre résultat influera sur la gravité des conséquences\033[0m")
		fmt.Println("  ")
		var inputD string
		fmt.Scanln(&inputD)
		switch inputD {
		case "D":
			var résultat int
			résultat = p.JetdeDé()
			fmt.Println("\033[1mVous avez fait un \033[0m", résultat)
			fmt.Println("  ")
			time.Sleep(500 * time.Millisecond)
			if résultat >= 17 {
				fmt.Println("\033[1mVous traversez le pont comme une flèche, posant les pieds juste là où il faut,")
				time.Sleep(500 * time.Millisecond)
				fmt.Println(" et en vous retournant, vous avez tout juste le temps de le voir finir")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("de s'effondrer suite à votre passage. Voilà qui devrait ralentir les loups !\033[0m")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("  ")
				fmt.Println("Vous gagnez le titre : L'Agile")
				fmt.Println("  ")
				fmt.Println("\033[1mDerrière vous, vous apercevez un chemin sinueux qui continue.\033[0m")
				p.SuiteChemin()
				return
			} else if résultat <= 16 && résultat >= 14 {
				fmt.Println("\033[1mVous traversez doucement le pont, posant prudemment vos pieds, en testant")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("chaque planche au préalable. Au bout d'un moment vous arrivez de l'autre côté,")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("transpirant à grosses gouttes. Vous espérez que les loups ne seront pas aussi")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("téméraires que vous, et qu'ils feront le tour.")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Derrière vous, vous apercevez un chemin sinueux qui continue.\033[0m")
				p.SuiteChemin()
				return
			} else if résultat <= 13 && résultat >= 10 {
				fmt.Println("\033[1mVous vous mettez à quatre pattes, aucun risque ne sera pris !")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Vous traversez extrêmement lentement, tellement, que vous commencez")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("à vous demander si faire le tour n'aurait pas été tout aussi rapide.")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Mais vous arrivez tout de même sain et sauf de l'autre côté.")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Derrière vous, vous apercevez un chemin sinueux qui continue.\033[0m")
				p.SuiteChemin()
				return
			} else if résultat <= 9 && résultat >= 5 {
				fmt.Println("\033[1mVous traversez avec grande difficulté le pont, lentement.")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Soudainement, une planche du pont que vous n'avez pas pensé à tester")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("casse sous votre pied, votre jambe passe au travers du bois, et le")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("bois vous écorche tout le long de la jambe.")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("  ")
				fmt.Println("\033[91mVous perdez 10 PV\033[0m")
				p.hp -= 10
				fmt.Println("  ")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("\033[1mVous vous sortez la jambe malgré tout, et continuez votre progression.")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Vous arrivez finalement de l'autre côté sans autre frayeur, mais avec une belle plaie.")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Derrière vous, vous apercevez un chemin sinueux qui continue.\033[0m")
				p.SuiteChemin()
				return
			} else if résultat < 5 {
				fmt.Println("\033[1mVous êtes confiant, vous marchez sur le pont en regardant où poser les pieds,")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("un véritable jeu d'enfant, il suffit de ne pas marcher sur les planches trop abîmées.")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Vous êtes à quelques mètres du bout du pont, la marche fière tellement la")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("facilité de l'opération est déconcertante.")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("  ")
				fmt.Println("\033[91mLes cordes du pont lachent.\033[0m")
				fmt.Println("  ")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("\033[1mVous vous voyez tomber au ralenti, votre vie défilant devant vos yeux.")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Finalement, vous étiez peut être un peu trop confiant dans vos capacités.")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("  ")
				fmt.Println("\033[91mVous mourrez écrasé au fond d'une crevasse.\033[0m")
				fmt.Println("  ")
				time.Sleep(500 * time.Millisecond)
				p.MortDefinitive()
				return
			}

		default:
			fmt.Println("  ")
			fmt.Println("Suivez les instructions du narrateur !")
			fmt.Println("  ")
			p.StartMenu()
			return
		}
	case "2":
		fmt.Println("\033[1mVous vous mettez en route pour faire le tour, mais plus vous")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("progressez, et plus les hurlements des loups se font insistants.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Vous regardez au loin pour trouver le bout de la crevasse, mais celle-ci")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("ne semble pas se finir. Vous continuez à marcher.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Vous entendez maintenant les hurlements juste derrière vous, mais")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("heureusement vous apercevez un endroit où la crevasse est plus resserée.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("C'est votre seule chance de traverser, si vous attendez plus les loups seront sur vous.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("1.- sauter")
		fmt.Println("2.- Patienter")
		fmt.Println(" ")
		var input string
		fmt.Scanln(&input)
		switch input {
		case "1":
			fmt.Println(" ")
			fmt.Println("Lancez le dé (Entrez D), faites 14 ou plus")
			fmt.Println(" ")
			var inputD string
			fmt.Scanln(&inputD)
			switch inputD {
			case "D":
				var résultat int
				résultat = p.JetdeDé()
				fmt.Println("\033[1mVous avez fait un \033[0m", résultat)
				fmt.Println("  ")
				time.Sleep(500 * time.Millisecond)
				if résultat >= 14 {
					fmt.Println("Vous sautez et arrivez tout juste de l'autre côté de la crevasse.")
					time.Sleep(500 * time.Millisecond)
					fmt.Println("Vous vous retournez et voyez les loups arriver de l'autre côté qui")
					time.Sleep(500 * time.Millisecond)
					fmt.Println("vous fixe intensément. Soudainement la meute saute par le même chemin que vous")
					time.Sleep(500 * time.Millisecond)
					fmt.Println("et la meute entière fond sur vous.")
					time.Sleep(500 * time.Millisecond)
					fmt.Println("Vous pensiez vraiment que des loups ne réussiraient pas à sauter ce qu'un")
					time.Sleep(500 * time.Millisecond)
					fmt.Println("humain peut sauter ?")
					time.Sleep(500 * time.Millisecond)
					fmt.Println(" ")
					fmt.Println("Vous êtes mort, dévoré par une meute de loups.")
					fmt.Println(" ")
					p.MortDefinitive()
					return
				} else {
					fmt.Println("Vous tentez le saut, mais ne touchez l'autre extrémité que du bout des doigts.")
					time.Sleep(500 * time.Millisecond)
					fmt.Println("Vous chutez vers ce qui est votre fin.")
					time.Sleep(500 * time.Millisecond)
					fmt.Println(" ")
					fmt.Println("Vous êtes mort, écrasé au fond d'une crevasse.")
					fmt.Println(" ")
					p.MortDefinitive()
					return
				}
			default:
				fmt.Println("  ")
				fmt.Println("Suivez les instructions du narrateur !")
				fmt.Println("  ")
				p.StartMenu()
				return
			}
		case "2":
			{
				fmt.Println("Vous vous retournez et voyez les loups arriver, ils courent en vous")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("fixant intensément. Sans arrêter leur course ils vous saute dessus")
				time.Sleep(500 * time.Millisecond)
				fmt.Println("et vous déchiquète.")
				time.Sleep(500 * time.Millisecond)
				fmt.Println(" ")
				fmt.Println("Vous êtes mort, dévoré par une meute de loups.")
				fmt.Println(" ")
				p.MortDefinitive()
				return
			}
		}
	}
}

func (p *Player) Massure() {
	fmt.Println("\033[1mVous approchez doucement de cette massure, plus vous approchez et moins")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("l'espoir de trouver quoi que ce soit se fait présent. Vous entendez le bruit de l'eau qui coule derrière la massure")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Vous arrivez devant la porte, vous pouvez toquer ou essayer d'ouvrir.\033[0m")
	time.Sleep(500 * time.Millisecond)
	fmt.Println(" ")
	fmt.Println("1.- Toquer")
	fmt.Println("2.- Ouvrir")
	fmt.Println(" ")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("\033[1mVous toquer à la porte et vous attendez ...")
		time.Sleep(5 * time.Second)
		fmt.Println("Après 5 bonnes secondes, vous entendez enfin du bruit.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("La porte s'ouvre sur un vieil homme, extrêmement courbé sous le poid de l'âge.\033[0m")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("\033[93mVieil Homme\033[0m - Bonjour, que puis-je faire pour vous mon brave ? ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[94m", p.name, "\033[0m", "- Bonjour, je suis un peu perdu, sauriez-vous où se situe la ville la plus proche ?")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[93mVieil Homme\033[0m - L'île la plus proche ? ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[94m", p.name, "\033[0m", "- Non, non, LA VILLE la plus proche !")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[93mVieil Homme\033[0m - Ah ! La ville, et bien vous avez la ville d'Alyoli si vous continuez plus au sud")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[94m", p.name, "\033[0m", "- Merci, quel serait le chemin le plus rapide ?")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[93mVieil Homme\033[0m - Huuuuummm ... ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[94m", p.name, "\033[0m", "- ...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[93mVieil Homme\033[0m - Et bien ... ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[94m", p.name, "\033[0m", "- ...")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[93mVieil Homme\033[0m - Vous devez suivre le ruisseau il me semble")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[94m", p.name, "\033[0m", "- Merci ! Bonne journ... ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[93mVieil Homme\033[0m - Ce ruisseau est tellement beau l'été, un bleu brillant et resplendissant")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[94m", p.name, "\033[0m", "- BONNE JOURNEE !")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[93mVieil Homme\033[0m - Oh ! Vous partez ? Bon voyage.")
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[1mVous repartez après cette discussion passionnante,")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Préférez-vous continuer dans la forêt, ou suivre le ruisseau ?\033[0m")
		fmt.Println(" ")
		fmt.Println("     ______________________________________________________     ")
		fmt.Println("<== || \033[34mContinuer dans la forêt |1\033[0m|||\033[35m2| Suivre le ruisseau\033[0m || ==>")
		fmt.Println("     ======================================================     ")
		var input string
		fmt.Scanln(&input)
		switch input {
		case "1":
			p.SuiteForet()
			return
		case "2":
			p.Ruisseau()
			return
		}

	case "2":
		fmt.Println("\033[1mVous ouvrez la porte sans ménagement")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Vous vous retrouvez face à un vieil homme, vraiment très âgé.\033[0m")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("\033[93mVieil Homme\033[0m - Que faites-vous chez moi !! Sortez immédiatement ! ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(p.name, " - Bonjour, désolé je suis perdu, je cherche une v...")
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[1mLe vieil homme vous frappe avec sa canne !\033[0m")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("\033[1m\033[91mvous perdez 5 PV\033[0m")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		p.hp -= 5
		fmt.Println("\033[93mVieil Homme\033[0m - Sortez immédiatement de chez moi !")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("\033[1mVous pouvez choisir de sortir de chez lui ou bien l'attaquer.\033[0m")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("1.- Sortir en courant !")
		fmt.Println("2.- L'attaquer !")
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		var input string
		fmt.Scanln(&input)
		switch input {
		case "1":
			fmt.Println("\033[1mVous sortez en courant, et vous continuez à vous enfoncer dans la forêt, ")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("vous courez bien plus vite, il ne vous rattrapera pas ! ")
			time.Sleep(500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("Vous avez gagné le titre : Le Malpoli !\033[0m")
			time.Sleep(500 * time.Millisecond)
			fmt.Println(" ")
			p.SuiteForet()
			return

		case "2":
			fmt.Println("\033[1mVous prenez une position de combat, vous lui lancez une énorme droite, ")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("le vieil homme s'écroule sous la force du coup.")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Vous le regardez inanimé par terre, et vous quittez les lieux.")
			time.Sleep(500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("Vous avez gagné le titre : Le Gérontophobe !\033[0m")
			time.Sleep(500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("\033[1mVous courez en direction de la forêt sans regarder derrière vous.\033[0m")
			time.Sleep(500 * time.Millisecond)
			p.SuiteForet()
			return
		}
	default:
		fmt.Println("Suivez les instructions du narrateur !")
		p.Massure()
	}
}

func (p *Player) Ruisseau() {
	fmt.Println("\033[1mVous commencez à longer le ruisseau, tout le monde sait que les villes")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("se construisent autour de points d'eau. Vous marchez pendant 2h.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Vous croisez un chemin avec deux panneaux.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("L'un indique Alyoli pointant vers l'ouest.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("L'autrre indique Donjon Obscur pointant vers l'Est.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Quel chemin souhaitez-vous emprunter ?")
	time.Sleep(500 * time.Millisecond)
	fmt.Println(" ")
	fmt.Println("                               ___                                    ")
	fmt.Println("                              |   |                                   ")
	fmt.Println("             _________________|___|___________                        ")
	fmt.Println("           <||       \033[34mA L Y O L I\033[0m\033[1m      || \033[34m1km\033[0m\033[1m ||                     ")
	fmt.Println("             =================================                        ")
	fmt.Println("                              |   |                                   ")
	fmt.Println("                              |   |                                   ")
	fmt.Println("                ______________|___|___________________                ")
	fmt.Println("               || \033[35m2km\033[0m\033[1m ||       \033[35mDonjon Obscur\033[0m\033[1m         ||>              ")
	fmt.Println("                ======================================               ")
	fmt.Println("                              |   |                                   ")
	fmt.Println("                              |   |                                   ")
	fmt.Println("                              |   |                                   ")
	fmt.Println("                              |   |                                   ")
	fmt.Println("                              |   |                                   ")
	fmt.Println(" ")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "1":
		fmt.Println("\033[1mVous suivez le chemin en direction d'ALYOLI")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Plus vous approchez, et plus vous apercevez de belles et grandes maisons au loin.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Vous arrivez devant la grande porte d'entrée de la ville.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Postés devant se trouvent 2 gardes qui vous regarde approcher.\033[0m")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("\033[93mGarde\033[0m - Halte là ! Que venez vous faire ici pouilleux ?")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[94m", p.name, "\033[0m", "- Bonjour ... Je cherchais la ville pour me restaurer et m'équiper.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[93mGarde\033[0m - Et avec quel argent minable ? ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[94m", p.name, "\033[0m", "- J'ai de l'argent !")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("*Vous lui montrez votre bourse, bien remplie.*")
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[93mGarde\033[0m - Bon si t'as de l'argent, tu peux payer le droit d'entrée alors ! Hein minable ?")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[93mGarde\033[0m - C'est 10 pièces d'or, mais comme je t'aime bien, ce sera 20 pièces d'or pour toi !")
		fmt.Println(" ")
		fmt.Println("*Il vous regarde en riant*")
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[1mQue souhaitez vous faire ?\033[0m")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("1.- Payer")
		fmt.Println("2.- L'attaquer !")
		fmt.Println(" ")
		var input string
		fmt.Scanln(&input)
		switch input {
		case "1":
			fmt.Println(" ")
			fmt.Println("\033[93mGarde\033[0m - Alors, qu'est ce que t'attends ?")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[94m", p.name, "\033[0m", "- C'est bon, je vais payer.")
			p.money -= 20
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[93mGarde\033[0m - Alors tu peux passer le bouseux.")
			time.Sleep(500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("\033[1m*Il vous regarde passer en riant encore plus fort*\033[0m")
			fmt.Println(" ")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[1mVous grommeler un merci les dents serrées, vous retenant de le frapper.")
			fmt.Println(" ")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Vous entrez dans la magnifique ville d'AYOLI, des couleurs splendides,")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("une multitude de personne, des maisons hautes avec balcons, surplombée")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("de verdure. Mais vous apercevez surtout un marchand et un forgeron qui")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("piquent à tout deux grandement votre interêt. Il ne me reste qu'à vous souhaiter à vous joueur ...\033[0m")
			time.Sleep(500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("\033[1mBIENVENUE DANS LA VILLE D'ALYOLI !!\033[0m")
			fmt.Println(" ")
			time.Sleep(500 * time.Millisecond)
			fmt.Println(" ")
			p.MainMenu()
			return
		case "2":
			fmt.Println("Ne supportant plus ses remarques, vous vous jetez sur lui !")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Préparez vous au combat !")
			p.MainMenu()
			return
		}
	case "2":
		p.Donjon()
		return
	default:
		fmt.Println("  ")
		fmt.Println("Suivez les instructions du narrateur !")
		fmt.Println("  ")
		p.StartMenu()
		return
	}

}

func (p *Player) SuiteChemin() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(" ")
	fmt.Println("\033[1mVous suivez le chemin qui commence doucement à s'entourer d'arbres morts.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Pas très rassurant, mais après quelques minutes de marche, vous arrivez en vue")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("d'un grand et magnifique donjon, totalement construit en pierre sombre.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Vous n'en aviez jamais vu de si haut.")
	fmt.Println(" ")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Votre chemin en rejoint un autre, et à cette jonction, vous apercevez un panneau.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("L'un indique Alyoli pointant vers l'ouest.")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("L'autre indique Donjon Obscur pointant vers l'Est. Probablement celui que vous apercevez au loin.")
	fmt.Println(" ")
	fmt.Println("                               ___                                    ")
	fmt.Println("                              |   |                                   ")
	fmt.Println("             _________________|___|___________                        ")
	fmt.Println("           <||       \033[34mA L Y O L I\033[0m\033[1m      || \033[34m2km\033[0m\033[1m ||                     ")
	fmt.Println("             =================================                        ")
	fmt.Println("                              |   |                                   ")
	fmt.Println("                              |   |                                   ")
	fmt.Println("                ______________|___|___________________                ")
	fmt.Println("               || \033[35m1km\033[0m\033[1m ||       \033[35mDonjon Obscur\033[0m\033[1m         ||>              ")
	fmt.Println("                ======================================               ")
	fmt.Println("                              |   |                                   ")
	fmt.Println("                              |   |                                   ")
	fmt.Println("                              |   |                                   ")
	fmt.Println("                              |   |                                   ")
	fmt.Println("                              |   |                                   ")
	fmt.Println(" ")
	var input string
	fmt.Scanln(&input)
	switch input {
	case "2":
		fmt.Println("\033[1mVous suivez le chemin en direction d'ALYOLI")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Plus vous approchez, et plus vous apercevez de belles et grandes maisons au loin.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Vous arrivez devant la grande porte d'entrée de la ville.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Postés devant se trouvent 2 gardes qui vous regarde approcher.\033[0m")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("\033[93mGarde\033[0m - Halte là ! Que venez vous faire ici pouilleux ?")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[94m", p.name, "\033[0m", "- Bonjour ... Je cherchais la ville pouir me restaurer et m'équiper.")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[93mGarde\033[0m - Et avec quel argent minable ? ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[94m", p.name, "\033[0m", "- J'ai de l'argent !")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("\033[1m*Vous lui montrez votre bourse, bien remplie.*\033[0m")
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[93mGarde\033[0m - Bon si t'as de l'argent, tu peux payer le droit d'entrée alors ! Hein minable ?")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[93mGarde\033[0m - C'est 10 pièces d'or, mais comme je t'aime bien, ce sera 20 pièces d'or pour toi !")
		fmt.Println(" ")
		fmt.Println("*Il vous regarde en riant*")
		fmt.Println(" ")
		time.Sleep(500 * time.Millisecond)
		fmt.Println("\033[1mQue souhaitez vous faire ?\033[0m")
		time.Sleep(500 * time.Millisecond)
		fmt.Println(" ")
		fmt.Println("1.- Payer")
		fmt.Println("2.- L'attaquer !")
		fmt.Println(" ")
		var input string
		fmt.Scanln(&input)
		switch input {
		case "1":
			fmt.Println(" ")
			fmt.Println("\033[93mGarde\033[0m - Alors, qu'est ce que t'attends ?")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[94m", p.name, "\033[0m", "- C'est bon, je vais payer.")
			p.money -= 20
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[93mGarde\033[0m - Alors tu peux passer le bouseux.")
			time.Sleep(500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("\033[1m*Il vous regarde passer en riant encore plus fort*\033[0m")
			fmt.Println(" ")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("\033[1mVous grommeler un merci les dents serrées, vous retenant de le frapper.")
			fmt.Println(" ")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Vous entrez dans la magnifique ville d'AYOLI, des couleurs splendides,")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("une multitude de personne, des maisons hautes avec balcons, surplombée")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("de verdure. Mais vous apercevez surtout un marchand et un forgeron qui")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("piquent à tout deux grandement votre interêt. Il ne me reste qu'à vous souhaiter à vous joueur ...\033[0m")
			time.Sleep(500 * time.Millisecond)
			fmt.Println(" ")
			fmt.Println("\033[1mBIENVENUE DANS LA VILLE D'ALYOLI !!\033[0m")
			fmt.Println(" ")
			time.Sleep(500 * time.Millisecond)
			fmt.Println(" ")
			p.MainMenu()
			return
		case "2":
			fmt.Println("Ne supportant plus ses remarques, vous vous jetez sur lui !")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("Préparez vous au combat !")
			fmt.Println(" ")
			p.CombatGardes()
			return
		}
	case "1":
		p.Donjon()
		return
	default:
		fmt.Println("  ")
		fmt.Println("Suivez les instructions du narrateur !")
		fmt.Println("  ")
		p.StartMenu()
		return
	}

}

func (p *Player) CombatGardes() {

}
