package sublist

func Sublist(l1, l2 []int) Relation {
	switch {
	case isEqual(l1, l2):
		return RelationEqual

	case isSublist(l1, l2):
		return RelationSublist

	case isSublist(l2, l1):
		return RelationSuperlist

	default:
		return RelationUnequal
	}
}

func isEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}
	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}
	return true
}

func isSublist(a, b []int) bool {
	if len(a) > len(b) {
		return false
	}
	for i := 0; i <= len(b)-len(a); i++ {
		if isEqual(a, b[i:i+len(a)]) {
			return true
		}
	}
	return false
}
