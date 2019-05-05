package sodukusolver

//contains indicates whether the array contains the specified value
func contains(array []int, value int) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}
