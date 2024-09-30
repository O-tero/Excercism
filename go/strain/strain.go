package strain

// Keep function returns a new collection with elements where predicate is true
func Keep[T any](collection []T, predicate func(T) bool) []T {
    var result []T
    for _, element := range collection {
        if predicate(element) {
            result = append(result, element)
        }
    }
    return result
}

func Discard[T any](collection []T, predicate func(T) bool) []T {
    var result []T
    for _, element := range collection {
        if !predicate(element) {
            result = append(result, element)
        }
    }
    return result
}
