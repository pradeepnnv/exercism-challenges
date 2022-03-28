package lasagna

func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		avgPrepTime = 2
	}
	return avgPrepTime * len(layers)
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	for _, layer := range layers {
		if layer == "sauce" {
			sauce += 0.2
		} else if layer == "noodles" {
			noodles += 50
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, scale int) []float64 {
	newAmounts := make([]float64, len(amounts))
	for i := range amounts {
		newAmounts[i] = amounts[i] * float64(scale) / 2
	}
	return newAmounts
}
