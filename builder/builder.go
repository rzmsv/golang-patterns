package main

import (
	"fmt"
)

type Car struct {
	brand string
	color string
	year  int
}

type CarBuilder interface {
	SetBrand(brand string) CarBuilder
	SetColor(color string) CarBuilder
	SetYear(year int) CarBuilder
	Build() Car
}

type ConcreteCarBuilder struct {
	car Car
}

func (b *ConcreteCarBuilder) SetBrand(brand string) CarBuilder {
	b.car.brand = brand
	return b
}
func (b *ConcreteCarBuilder) SetColor(color string) CarBuilder {
	b.car.color = color
	return b
}
func (b *ConcreteCarBuilder) SetYear(year int) CarBuilder {
	b.car.year = year
	return b
}
func (b *ConcreteCarBuilder) Build() Car {
	return b.car
}

func main() {
	builder := &ConcreteCarBuilder{}
	car1 := builder.SetBrand("Benz").SetColor("Blue").SetYear(1989).Build()
	car2 := builder.SetBrand("BMW").SetColor("Blue").SetYear(1989).Build()
	fmt.Printf("%+v", car1)
	fmt.Printf("%+v", car2)
}
