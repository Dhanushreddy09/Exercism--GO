package card_tricks

func GetItem(slice []int, index int) (int, bool) {
	if index < 0 || index >= len(slice) {
		return 0, false
	}
	return slice[index], true
}
func SetItem(slice []int, index, value int) []int {
	if index >= 0 && index < len(slice) {
		slice[index] = value
	} else {
		slice = append(slice, value)
	}
	return slice
}
func PrefilledSlice(value, length int) []int {
	slice := []int{}
	for len(slice) < length {
		slice = append(slice, value)
	}
	return slice
}
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index > len(slice)-1 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
