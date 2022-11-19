package course

type Course struct {
	title string
	users []Userer
}

func (c *Course) addUser(u Userer) {
	c.users = append(c.users, u)
}

func (c *Course) Users() []Userer {
	return c.users
}

func NewCourse(title string, u ...Userer) *Course {
	c := &Course{title: title}
	for _, a := range u {
		c.addUser(a)
	}

	return c
}

type Getter interface {
	Get(some any) bool
}
type Userer interface {
	ShowMyTasks()
}
