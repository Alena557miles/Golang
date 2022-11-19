package strategies

import "art-strategy/core"

type Home struct {
}

func (c Home) GetName() string {
	return `Home`
}

func (c Home) ChoseArt(a core.Art) bool {
	return a.Befits() == `rest`
}
