package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
	for _, val := range s {
		acc = fn(acc, val)
	}
	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
	for i := len(s) - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	var result IntList
	for _, val := range s {
		if fn(val) {
			result = append(result, val)
		}
	}

	if result == nil {
		return IntList{}
	}
	return result
}


func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	result := make(IntList, len(s))
	for i, val := range s {
		result[i] = fn(val)
	}
	return result
}

func (s IntList) Reverse() IntList {
	result := make(IntList, len(s))
	for i, val := range s {
		result[len(s)-1-i] = val
	}
	return result
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	result := s
	for _, list := range lists {
		result = append(result, list...)
	}
	return result
}
