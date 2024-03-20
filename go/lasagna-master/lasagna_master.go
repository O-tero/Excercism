package lasagna

// PreparationTime estimates the total preparation time based on the number of layers
func PreparationTime(layers []string, avgPerLayer int) int {
	if avgPerLayer == 0 {
		avgPerLayer = 2
	}
	return len(layers) * avgPerLayer
}

// Quantities determines the quantity of noodles and sauce needed to make your meal
func Quantities(layers []string) (int, float64) {
	noodles, sauce := 0, float64(0)
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// AddSecretIngredient adds the secret ingredient to your recipe
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// ScaleRecipe gives the amounts needed for the desired number of portions
func ScaleRecipe(list []float64, portions int) []float64 {
	output := make([]float64, len(list))
	for i := 0; i < len(list); i++ {
		output[i] = list[i] * float64(portions) / 2
	}
	return output
}
