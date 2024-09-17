package Projet_Red

type Player struct {
	name         string
	class        string
	level        int
	hpmax        int
	hp           int
	money        int
	inventorymax int
	inventory    []Inventory
	skill        []string
	equipement   []Equipement
	mname        string
	mhpmax       int
	mhp          int
	mdamage      int
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
