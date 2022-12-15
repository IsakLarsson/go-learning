package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		return len(layers) * 2

	}
	return len(layers) * timePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	sauce := 0.0
	noodles := 0
	for _, v := range layers {
		switch {
		case v == "sauce":
			sauce += 0.2
		case v == "noodles":
			noodles += 50

		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendIngredients []string, myIngredients []string) []string {
	myIngredients[len(friendIngredients)] = friendIngredients[len(friendIngredients)-1]
	return myIngredients
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	newQuantities := make([]float64, len(amounts))
	for i, v := range amounts {
		newQuantities[i] = v * float64(portions) / 2
	}
	return newQuantities
}
