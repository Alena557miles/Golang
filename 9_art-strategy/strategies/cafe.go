package strategies

import "art-strategy/core"

type Cafe struct {
}

func (c Cafe) GetName() string {
	return `GallCafeery`
}

func (c Cafe) ChoseArt(a core.Art) bool {
	return a.Befits() == `eat`
}
