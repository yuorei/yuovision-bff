package lib

// stringを*stringに変換する
func StringPointer(s string) *string {
	return &s
}

func PointersToStrings(sliceOfPointers []*string) []string {
	sliceOfString := make([]string, len(sliceOfPointers))

	for i, v := range sliceOfPointers {
		if v != nil {
			sliceOfString[i] = *v
		}
	}

	return sliceOfString
}

func StringsToPointers(sliceOfString []string) []*string {
	sliceOfPointers := make([]*string, len(sliceOfString))

	for i, v := range sliceOfString {
		copy := v
		sliceOfPointers[i] = &copy
	}

	return sliceOfPointers
}
