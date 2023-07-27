package lasagna

func PreparationTime(layers []string, avg int) int {
	if avg == 0 {
		avg = 2
	}

	return len(layers) * avg
}

func Quantities(layers []string) (int, float64) {

	sumNoodles := 0
	sumSauce := 0.0

	for _, v := range layers {
		if v == "noodles" {
			sumNoodles += 50
		} else if v == "sauce" {
			sumSauce += 0.2
		}

	}

	return sumNoodles, sumSauce
}

func AddSecretIngredient(friendsList, myList []string) []string {
	lenFriendList := len(friendsList)
	lenMyList := len(myList)

	myList[lenMyList-1] = friendsList[lenFriendList-1]

	return nil
}

func ScaleRecipe(quantities []float64, portion int) []float64 {
	newQuantities := []float64{}
	for _, quantity := range quantities {
		q := quantity / 2
		q = q * float64(portion)
		newQuantities = append(newQuantities, q)
	}
	return newQuantities
}
