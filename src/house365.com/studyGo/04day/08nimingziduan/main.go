package main

import "fmt"

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s会动!\n", a.name)
}

type dog struct {
	feet uint8
	animal
}

func (d dog) wang() {
	fmt.Printf("%s汪汪汪~~~!\n", d.name)
}
func main() {
	d := dog{
		feet:   4,
		animal: animal{name: "周"},
	}
	d.wang()
	d.move()

}
