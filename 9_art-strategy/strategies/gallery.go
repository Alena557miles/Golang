package strategies

import "art-strategy/core"

type Gallery struct {
}

func (c Gallery) GetName() string {
	return `Gallery`
}

func (c Gallery) ChoseArt(a core.Art) bool {
	return a.Befits() == `beauty`
}
