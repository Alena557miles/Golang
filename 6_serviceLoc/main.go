package main

import (
	"fmt"
	"reflect"
)

type ServiceLocator struct {
	services []any
	names    []string
}

func (s *ServiceLocator) Register(name string, some any) {
	s.services = append(s.services, some)
	s.names = append(s.names, name)
}

func (s *ServiceLocator) Get(name string) any {
	for i, t := range s.names {
		if t == name {
			return s.services[i]
		}
	}
	return nil
}

type Concrete struct {
	Name string
}

func (c *Concrete) Do() {}

type Doer interface {
	Do()
}

func main() {
	l := ServiceLocator{}

	a := &Concrete{"first"}
	a.Do()
	fmt.Println("Method Name: ", a.Name)
	l.Register(`Aname`, a)
	aM := l.Get(`Aname`)
	fmt.Println("Method from SL: ", aM)
	fmt.Println("Type of aM: ", reflect.TypeOf(aM))

	b := Concrete{"second"}
	b.Do()
	fmt.Println("Method Name: ", b.Name)
	l.Register(`Bname`, b)
	bM := l.Get(`Bname`)
	fmt.Println("Method from SL: ", bM)
	fmt.Println("Type of aM: ", reflect.TypeOf(bM))

	// var c Doer
	// l.Register("cname", c)
	// cM := l.Get("cname")

	// fmt.Println("c: ", cM)
	// fmt.Println("Type of c: ", reflect.TypeOf(cM))
	fmt.Println(l)

}
