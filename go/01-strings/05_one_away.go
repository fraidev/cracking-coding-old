package strings

func OneAway(i1, i2 string) bool {

	//Can Add?
	if (len(i1) - 1) == len(i2) {
		return oneEditInsert(i2, i1)
	}

	//Can Chenge?
	if len(i1) == len(i2) {
		return oneEditReplace(i1, i2)
	}

	//Can Remove?
	if (len(i1) + 1) == len(i2) {
		return oneEditInsert(i1, i2)
	}

	return false
}

func oneEditReplace(i1, i2 string) bool {
	foundDiff := false

	for i := range i1 {
		if i1[i] != i2[i] {
			if foundDiff {
				return false
			}
			foundDiff = true
		}
	}
	return true
}

func oneEditInsert(i1, i2 string) bool {
	index1 := 0
	index2 := 0

	for index2 < len(i2) && index1 < len(i1) {
		if i1[index1] != i2[index2] {
			if index1 != index2 {
				return false
			}
			index2++
		}
		index1++
		index2++
	}

	return true
}
