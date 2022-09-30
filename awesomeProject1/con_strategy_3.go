package main

import "fmt"

type pot struct {
	name string
}

func (d *pot) cook() {
	fmt.Printf("Your %s was successfully boiled on pot", d.name)
}

func newPot(name string) *pot {
	return &pot{
		name: name,
	}
}
