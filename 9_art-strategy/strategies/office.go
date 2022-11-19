package strategies

import "art-strategy/core"

type Office struct {
}

func (c Office) GetName() string {
	return `Office`
}

func (c Office) ChoseArt(a core.Art) bool {
	return a.Befits() == `work`
}
