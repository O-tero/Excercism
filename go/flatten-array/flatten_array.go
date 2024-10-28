package flatten

func Flatten(nested interface{}) []interface{} {
	var result []interface{}

	
	var flattenHelper func(interface{})
	flattenHelper = func(item interface{}) {
		switch v := item.(type) {
		case []interface{}:
					for _, elem := range v {
				flattenHelper(elem) 
			}
		case nil:
						return
		default:
					result = append(result, v)
		}
	}

	flattenHelper(nested)

	if result == nil {
		return []interface{}{}
	}
	return result
}
