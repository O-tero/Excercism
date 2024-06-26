package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	return min(option1, option2) + " is clearly the better choice."
}

func min(first string, second string) string {
	if first < second {
		return first
	}
	return second
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return .8 * originalPrice
	} else if age < 10 {
		return .7 * originalPrice
	} else {
		return .5 * originalPrice
	}
}
