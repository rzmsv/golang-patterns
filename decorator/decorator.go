package main

import "fmt"

type Coffee interface {
	Cost() float64
	Description() string
}

type SimpleCoffee struct{}

func (s *SimpleCoffee) Cost() float64 {
	return 0.5
}
func (s *SimpleCoffee) Description() string {
	return "Simple Coffee "
}

/* -------------------------------- Decorator ------------------------------- */
type CoffeeWithMilk struct {
	coffee Coffee
}

func (c *CoffeeWithMilk) Cost() float64 {
	return c.coffee.Cost() + 5
}
func (c *CoffeeWithMilk) Description() string {
	return c.coffee.Description() + "With Milk"
}

/* -------------------------------- Decorator ------------------------------- */

func main() {
	simpleCoffee := &SimpleCoffee{}
	fmt.Printf("Coffee: %s Cost: %v \n", simpleCoffee.Description(), simpleCoffee.Cost())

	coffeeWithMilk := &CoffeeWithMilk{coffee: simpleCoffee}
	fmt.Printf("Coffee: %s Cost: %v \n", coffeeWithMilk.Description(), coffeeWithMilk.Cost())
}
