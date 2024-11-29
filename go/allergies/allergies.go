package allergies

var allergenMap = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}


func Allergies(allergies uint) []string {
	var result []string
	for allergen, value := range allergenMap {
		if allergies&value > 0 { // Check if the bit for this allergen is set
			result = append(result, allergen)
		}
	}
	return result
}


func AllergicTo(allergies uint, allergen string) bool {
	value, exists := allergenMap[allergen]
	if !exists {
		return false 
	}
	return allergies&value > 0 // Check if the bit for this allergen is set
}
