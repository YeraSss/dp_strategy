package main

func startCooking(food string, cooking Cooking) {
	cooking.cook()
}

type food struct {
	dish string
}

func chooseDish(name string) *food {
	return &food{
		dish: name,
	}
}
