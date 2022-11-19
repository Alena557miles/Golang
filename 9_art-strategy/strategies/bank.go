package strategies

import "art-strategy/core"

type Bank struct {
}

func (c Bank) GetName() string {
	return `Bank`
}

func (c Bank) ChoseArt(a core.Art) bool {
	return a.Befits() == `financial organization`
}
