package Projet_Red

type Player struct {
	name      string
	class     string
	level     int
	hpmax     int
	hp        int
	inventory []Inventory
}

type Inventory struct {
	name     string
	quantity int
}
