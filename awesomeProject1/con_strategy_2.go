package main

import "fmt"

type pan struct {
	name string
}

func (d *pan) cook() {
	fmt.Printf("Your %s was successfully roasted on pan", d.name)
}

func newPan(name string) *pan {
	return &pan{
		name: name,
	}
}
