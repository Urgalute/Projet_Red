package Projet_Red

type Player struct {
	name      string
	class     string
	level     int
	hpmax     int
	hp        int
	inventory []Inventory
	skill []string
}

type Inventory struct {
	name        string
	quantity    int
	description string
}
