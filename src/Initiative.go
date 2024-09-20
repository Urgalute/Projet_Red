package Projet_Red

import (
	"math/rand"
	"time"
)

func (p *Player) InitiativeP() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(20) + 1
}

func (p *Player) InitiativeM() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(20) + 1
}

func (p *Player) JetdeDé() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(20) + 1
}

func (p *Player) InitiativeG1() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(20) + 1
}

func (p *Player) InitiativeG2() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(20) + 1
}