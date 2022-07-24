package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, servings int) int {
	if servings == 0 {
		servings = 2
	}
	return len(layers) * servings
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	sauce := 0.0
	noodles := 0
	// loop over layers using range
	for _, layer := range layers {
		if layer == "sauce" {
			sauce += 0.2
		} else if layer == "noodles" {
			noodles += 50
		}
	}
	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	newQuantities := make([]float64, len(quantities))
	for i := 0; i < len(quantities); i++ {
		newQuantities[i] = quantities[i] * float64(portions) / 2
	}
	return newQuantities
}
