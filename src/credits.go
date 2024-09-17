package Projet_Red

import (
	"fmt"
	"strings"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func (p *Player) Credits() {
	credits := `
============== CREDITS ==============

     Un jeu exceptionnel

    Codé intégralement en
	      Golang

   Avec la participation de

         Yoan Commeau

   	 Alexis Dupin

   	Lilian Le Piver

     Merci d'avoir joué !

========== ATTRIBUTION DES TACHES : ==========

	TÂCHE 1 : CRÉATION DU MENU
	Yoan



	TÂCHE 2 : CRÉATION DU PERSONNAGE
	Lilian


	TÂCHE 3 : INITIALISATION DU PERSONNAGE
	Lilian


	TÂCHE 4 - AFFICHAGE DES INFORMATIONS DU PERSONNAGE
	Lilian
	

	TÂCHE 5 - ACCÈS À L'INVENTAIRE
	Alexis


	TÂCHE 6 : POTION DE VIE
	Yoan


	TÂCHE 7 : MARCHAND
	Lilian


	TÂCHE 8 : WASTED !
	Alexis


	TÂCHE 9 : POTION DE POISON
	Alexis


	TÂCHE 10 : WINGARDIUM LEVIOSA
	Lilian


	MISSION 1 : AMÉLIORATION DE LA CRÉATION DE PERSONNAGE
	Alexis

	
	MISSION 2 : LIMITE D'INVENTAIRE
	Lilian


	TÂCHE 11 : MONEY, MONEY, MONEY
	Alexis


	TÂCHE 12 : TWO FOR THE PRICE OF ONE
	Alexis


	TÂCHE 13 : GIMME! GIMME! GIMME! 
	Yoan


	TÂCHE 14 : I SAW IT IN THE MIRROR
	Alexis


	TÂCHE 15 : MAMMA MIA
	Alexis


	MISSION 2.2 : ON AND ON AND ON
	Lilian


	TÂCHE 16 : LA CHOSE
	Alexis


	TÂCHE 17 : FIGHTER SQUAD
	Alexis
	

	TÂCHE 18 : A.I. INTELLIGENCE ARTIFICIELLE 
	Alexis


	TÂCHE 19 : READY PLAYER ONE
	Lilian


	TÂCHE 20 : DUEL
	Alexis, Lilian



	MISSION 3 : INITIATIVE
	Alexis


	MISSION 4 : EXPÉRIENCE



	MISSION 5.1 : COMBAT MAGIQUE



	MISSION 5.2 : RESSOURCE DE MANA



TRAVAIL SUR FICHIERS :

	accessinventory.go 
	Alexis

	Avec la participation de
	Yoan, Lilian



	BlackSmith.go
	Yoan



	credits.go
	Alexis



	equipementslot.go
	Alexis



	initgobelin.go
	Alexis

	Avec le soutien de
	Yoan



	Initiative.go
	Alexis, ChatGPT



	mainMenu.go
	Yoan

	Avec la participation de
	Alexis, Lilian



	merchant.go
	Lilian

	Avec la participation de
	Alexis, Yoan



	mort.go
	Alexis



	persoinit.go
	Alexis

	Avec le soutien de
	Yoan



	player.go
	Lilian

	Avec la participation de
	Alexis, Yoan



	skill.go
	Lilian

	Avec le soutien de
	Yoan


	Start.go
	Alexis

	Avec le soutien de
	Yoan



	struct.go
	Lilian

	Avec la participation de 
	Alexis, Yoan



	TakeItem.go
	Yoan

	Avec la participation de
	Lilian, Alexis



	TrainingFight.go
	Alexis

	Avec le soutien de
	Yoan


	Utility.go
	Yoan



	Création du git
	Yoan



	Création du dossier src et docs
	Yoan



	Organisation via Trello
	Yoan



	 Et un grand merci à

	  ABBA et Spielberg

	pour nous avoir gratifié
	  de leur présence !!


    Alyoli(c) 2024 Tous droits réservés
    `
	displayCredits(strings.TrimSpace(credits))
}

func displayCredits(credits string) {
	lines := strings.Split(credits, "\n")
	for i, line := range lines {
		switch {
		case i == 9:
			fmt.Println(Purple + line + Reset)
		case i == 11:
			fmt.Println(Green + line + Reset)
		case i == 13:
			fmt.Println(Cyan + line + Reset)
		default:
			fmt.Println(White + line + Reset)
		}
		time.Sleep(220 * time.Millisecond)
	}
}
