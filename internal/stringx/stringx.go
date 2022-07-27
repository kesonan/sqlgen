package stringx

// Contains returns true if the string contains the sub-item string.
func Contains(list []string, item string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}
