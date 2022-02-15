package gross_store

func Units() map[string]int {
	measurement := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return measurement
}
func NewBill() map[string]int {
	return make(map[string]int)
}
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, unitExists := units[unit]
	_, itemExists := bill[item]

	if !unitExists {
		return false
	}
	if !itemExists {
		bill[item] = units[unit]
	} else {
		bill[item] += units[unit]
	}
	return true
}
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, unitExists := units[unit]
	value, itemExists := bill[item]
	if !itemExists || !unitExists {
		return false
	}
	value -= units[unit]

	if value < 0 {
		return false
	}
	if value == 0 {
		delete(bill, item)
	} else {
		bill[item] = value
	}
	return true
}
func GetItem(bill map[string]int, item string) (int, bool) {
	_, itemExists := bill[item]
	if !itemExists {
		return 0, false
	}
	return bill[item], true
}
