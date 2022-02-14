package lasagna_master

func PreparationTime(layers []string, averagePreparationTime int) int {
	if averagePreparationTime == 0 {
		averagePreparationTime = 2
	}
	return len(layers) * averagePreparationTime
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, value := range layers {
		if value == "noodles" {
			noodles += 50
		}
		if value == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsRecipe []string, myRecipe []string) {
	myRecipe[len(myRecipe)-1] = friendsRecipe[len(friendsRecipe)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	scale := float64(portions)
	scaledRecipe := []float64{}
	for i := 0; i < len(quantities); i++ {
		scaledRecipe = append(scaledRecipe, quantities[i]*(scale/2))
	}
	return scaledRecipe
}
