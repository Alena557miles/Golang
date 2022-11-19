package core

type Building interface {
	GetName() string
	ChoseArt(Art) bool
}

type Art interface {
	Author() string
	Befits() string
	Goto(Building)
}
