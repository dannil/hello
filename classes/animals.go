package classes

import "fmt"

type Dog struct {
	Name string
	Age int
}

func (dog Dog) Bark() {
	fmt.Println("WOOF!")
}