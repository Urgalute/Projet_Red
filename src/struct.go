package Projet_Red

type Player struct {
	name         string
	class        string
	level        int
	hpmax        int
	hp           int
	xp           int
	xpmax        int
	dammage      int
	money        int
	mana         int
	manamax      int
	inventorymax int
	inventory    []Inventory
	skill        []string
	equipement   []Equipement
	monster      []Monster


	
	mname        string
	mhpmax       int
	mhp          int
	mdammage     int
	mxp          int
	mlevel       int
	g1name       string
	g1hpmax      int
	g1hp         int
	g1damage     int
}

type Inventory struct {
	name        string
	quantity    int
	description string
}

type Equipement struct {
	name        string
	quantity    int
	description string
}

type Monster struct {
	name   string
	hp     int
	hpmax  int
	damage int
	xp     int
	level  int
}
