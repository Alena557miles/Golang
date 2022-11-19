package main

import (
	"art-strategy/core"
	"art-strategy/strategies"
	"art-strategy/workshop"
	"time"
)

func main() {
	buildings := []core.Building{
		strategies.Bank{},
		strategies.Cafe{},
		strategies.Club{},
		strategies.Gallery{},
		strategies.Home{},
		strategies.Office{},
	}
	CheckArt(buildings)

}

func CheckArt(buildings []core.Building) {
	for {
		ca := workshop.CreateArt()
		for _, building := range buildings {
			if building.ChoseArt(&ca) {
				ca.Goto(building)
			}
		}
		time.Sleep(time.Second)
	}
}
