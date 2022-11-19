package strategies

import "art-strategy/core"

type Club struct {
}

func (c Club) GetName() string {
	return `Club`
}

func (c Club) ChoseArt(a core.Art) bool {
	return a.Befits() == `joy`
}
