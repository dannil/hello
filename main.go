package main

import (
	"fmt"

	"github.com/dannil/hello/classes"
	"github.com/dannil/hello/util"
	)

func main() {
	fmt.Printf("Hello, world PI - 44.\n")

	util.Hello_util()
	fmt.Printf(util.ToUpperCase("This should be upper-case\n"))

	d := &classes.Dog{Name: "Shredder", Age: 10}
	d.Bark()
	d.Rename("Bob")
	d.Info()

	fmt.Println(d.DogYears())

	fmt.Println(d.ToString())
}

