package classes

import "fmt"

type Dog struct {
	Name string
	Age int
}

func (dog *Dog) Bark() {
	fmt.Println("WOOF!")
}

func (dog *Dog) Rename(name string) {
	dog.Name = name
}

func (dog *Dog) Info() {
	fmt.Printf("Dog {name=%s, age=%d}\n", dog.Name, dog.Age)
}

func (dog *Dog) ToString() (string) {
	return fmt.Sprintf("%+v", dog)
}

func (dog *Dog) DogYears() (int) {
	return dog.Age * 7
}