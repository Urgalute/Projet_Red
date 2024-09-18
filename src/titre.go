package Projet_Red

<<<<<<< HEAD
//func (p *Player) Titre() {
//	if killgob {
//		p.name += "Goblin Slayer"
//	}else if deathgob {
//		p.name += "Le Faible"
//	}
//}
=======
func (p *Player) Titre() {
	if killgob == 1 {
		p.name += " Goblin Slayer"
		killgob++
	}
	if deathgob == 1 {
		p.name += " Le Faible"
		deathgob++
	}
}
>>>>>>> 369316ec8be3a0c6e2fe15ddaf477888e319561a
