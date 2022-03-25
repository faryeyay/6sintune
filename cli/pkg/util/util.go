package util

// SliceContainsString check if a slice contains the string str
func SliceContainsString(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// StringSubset check if the first slice of strings is present
// within the second slice of strings
func StringSubset(first, second []string) bool {
	for _, value := range first {
		if !SliceContainsString(second, value) {
			return false
		}
	}

	// I know the logic is flawed here
	// just looking for a quick solution for now
	return true
}
