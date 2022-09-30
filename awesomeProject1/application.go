package main

func main() {
	food := chooseDish("egg")
	cook := 2
	var cooking Cooking
	switch cook {
	case 1:
		cooking = newPan(food.dish)
	case 2:
		cooking = newPot(food.dish)
	case 3:
		cooking = newMicrowave(food.dish)
	}

	startCooking(food.dish, cooking)
}
