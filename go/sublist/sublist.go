package sublist

type Relation string

func Sublist(l1, l2 []int) string {
	var r Relation
	// checking basic conditions
	if l1 == nil && l2 == nil || len(l1) == 0 && len(l2) == 0 {
		return "equal"
	} else if len(l1) != 0 && (len(l2) == 0 || l2 == nil) {
		return "superlist"
	} else if (len(l1) == 0 || l1 == nil) && len(l2) != 0 {
		return "sublist"
	} else if len(l1) != len(l2) {
		return "unequal"
	}

	// lists of unequal length. Can only be a sublist or superlist.

	if len(l1) != len(l2) {
		if len(l1) > len(l2) {

		}
	}

	return string(r)
}
