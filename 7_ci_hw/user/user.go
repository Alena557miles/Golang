package user

type User struct {
	name  string
	email string
	td    TODOer
	On    Courser
}

func (u *User) ShowMyTasks() {
	u.td.ShowTasks()
}

func (u *User) SetTODO(sl Getter) {
	var t TODOer
	sl.Get(&t)
	u.td = t
}

func (u *User) SetCourse(sl Getter) {
	var c Courser
	sl.Get(&c)
	u.On = &c
}

func NewUser(n, e string) *User {
	return &User{name: n, email: e}
}

type TODOer interface {
	ShowTasks()
}

type Courser interface {
}
type ShowMyTasker interface {
}

type Getter interface {
	Get(some any) bool
}
