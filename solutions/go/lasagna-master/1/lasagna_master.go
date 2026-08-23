package lasagnamaster

func PreparationTime(layers []string, time int) int {
    if time == 0 {
        time = 2
    }
    return len(layers) * time
}

func Quantities(layers []string) (int, float64) {
    noodles := 0
    sauce := 0.0
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodles += 50
        } else if layers[i] == "sauce" {
            sauce += 0.2
        }
    }
    return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) []string {
    myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
    return myList
}

func ScaleRecipe(quantities []float64, amount int) []float64 {
    result := []float64 {}
    for i := 0; i < len(quantities); i++ {
        result = append(result, quantities[i] * float64(amount) * 0.5)
    }
    return result
}