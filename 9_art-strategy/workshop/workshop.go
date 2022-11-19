package workshop

import (
	"art-strategy/core"
	"fmt"
	"math/rand"
	"time"
)

type Art struct {
	author   string
	befits   string
	location core.Building
}

var Authors = []string{"Kevin", "Michail", "Eleonora", "Marichka", "Alexandr", "Oscar"}

func (a Art) Author() string {
	return a.author
}

func (a Art) Befits() string {
	return a.befits
}

func (a Art) Goto(b core.Building) {
	a.location = b
	fmt.Printf("Art with author %s go to  %s\n", a.Author(), b.GetName())
}

func CreateArt() Art {
	rand.Seed(time.Now().Unix())
	a := rand.Intn(6)
	b := BefitsGenerator()
	return Art{author: Authors[a], befits: b}
}

func BefitsGenerator() string {
	rand.Seed(time.Now().Unix())
	r := rand.Intn(6)
	switch r {
	case 0:
		return `eat`
	case 1:
		return `joy`
	case 2:
		return `beauty`
	case 3:
		return `rest`
	case 4:
		return `financial organization`
	default:
		fallthrough
	case 5:
		return `rest`
	}
}
