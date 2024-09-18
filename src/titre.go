package Projet_Red

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
