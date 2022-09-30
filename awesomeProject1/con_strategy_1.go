package main

import "fmt"

type microwave struct {
	name string
}

func (d *microwave) cook() {
	fmt.Printf("Your %s was successfully heated on microwave", d.name)
}

func newMicrowave(name string) *microwave {
	return &microwave{
		name: name,
	}
}
