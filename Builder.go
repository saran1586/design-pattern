package main

import "fmt"

type Pizza struct {
	Size    string

}

type PizzaBuilder struct {
	pizza Pizza
}

func NewPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{}
}

func (b *PizzaBuilder) SetSize(size string) *PizzaBuilder {
	b.pizza.Size = size
	return b
}


func (b *PizzaBuilder) Build() Pizza {
	return b.pizza
}

func main() {

	pizza := NewPizzaBuilder().
		SetSize("Large").
	
		Build()

	fmt.Println(pizza)
}
