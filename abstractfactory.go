package main

import "fmt"

type Animal interface {
	Sound()
}

type Dog struct{}

func (Dog) Sound() {
	fmt.Println("Bow Bow")
}

type Cat struct{}

func (Cat) Sound() {
	fmt.Println("Meow")
}

type Factory interface {
	CreateAnimal() Animal
}

type DogFactory struct{}

func (DogFactory) CreateAnimal() Animal {
	return Dog{}
}

type CatFactory struct{}

func (CatFactory) CreateAnimal() Animal {
	return Cat{}
}

func main() {
	var factory Factory

	factory = DogFactory{}
	factory.CreateAnimal().Sound()

	factory = CatFactory{}
	factory.CreateAnimal().Sound()
}
