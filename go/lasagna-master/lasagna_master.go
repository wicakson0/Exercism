package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2
	}

	return len(layers) * avgTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	const noodlesMeasure = 50
	const sauceMeasure = 0.2

	noodles := 0
	sauce := 0.0

	for _, v := range layers {
		if v == "sauce" {
			sauce += sauceMeasure
		}
		if v == "noodles" {
			noodles += noodlesMeasure
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaledAmounts := make([]float64, len(amounts))

	for i, v := range amounts {
		scaledAmounts[i] = v * float64(portions)
		scaledAmounts[i] /= 2
	}
	return scaledAmounts
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
