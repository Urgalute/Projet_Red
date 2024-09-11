package Projet_Red

import "fmt"

func InitPlayer() {
	item01 := Inventory{"Potion", 3}

	player1 := Player{
		name:  "bob",
		class: "guerrier",
		level: 1,
		hpmax: 100,
		hp:    40,
		inventory: []Inventory{
			item01,
		},
	}
	fmt.Println(player1)
}
