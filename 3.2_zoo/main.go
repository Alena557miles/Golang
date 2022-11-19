package main

import "fmt"

// Із зоопарком - створюємо програму із кількома структурами, де є абстрактний наглядач і
//абстрактні звири, та те, що ще буде потрібно. І наглядач збирає цих звірів і розташовую по клітках. Ясно, що властивості нагладачя,
// звірів та інших структур - це філди, а їх дії (“розбіжатись“, “зібрати“) - це методи.
// Тобто програма це такий собі кодовий мультфільм, який виконується

type animal struct {
	kind string
	cage *Cage
}
type Cage struct {
	animal *animal
}
type zooKeeper struct {
	cages []Cage
}

func (c Cage) IsEmpty() bool {
	return c.animal == nil
}

func (c Cage) Keep(a animal) {
	c.animal = &a
	a.cage = &c
	fmt.Println("Cage is Keeping:", c.animal)
	fmt.Println("Cage:", c)
	fmt.Println("Animal is keeping on cage:", a.cage)
}

func (z zooKeeper) Catch(a animal) {
	for i, cage := range z.cages {
		if cage.IsEmpty() {
			cage.Keep(a)
			return
		}
		i++
	}
}

func (a animal) Run() {
	a.cage = nil
	fmt.Println("Animal is running", a)
}

func main() {

	// ANIMALS
	tiger := animal{kind: "tiger"}
	hippo := animal{kind: "hippo"}
	giraffe := animal{kind: "giraffe"}
	raccoon := animal{kind: "racoon"}
	lion := animal{kind: "lion"}
	elephant := animal{kind: "racoon"}
	wolf := animal{kind: "wolf"}
	peacock := animal{kind: "peacock"}
	// fmt.Println(tiger, hippo, giraffe, raccoon, lion, elephant, wolf, peacock)

	// ZOO
	zoo := []animal{tiger, hippo, giraffe, raccoon, lion, elephant, wolf, peacock}
	amountAnimals := len(zoo)
	// fmt.Println(zoo)

	// ZOOKEEPER
	keeper := zooKeeper{make([]Cage, amountAnimals)}

	// ZOOKEEPER IS CATHING A TIGER
	keeper.Catch(tiger)
	//  TIGER IS RUNNING OUT
	tiger.Run()
	lion.Run()
	wolf.Run()
	keeper.Catch(lion)
}
