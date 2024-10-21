package lasagna

// TODO: define the 'OvenTime' constant
const (
    OvenTime = 40
    layerTime = 2
    )
func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

func PreparationTime(numberOfLayers int) int {
	return  numberOfLayers * layerTime
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return actualMinutesInOven + PreparationTime(numberOfLayers)
}
