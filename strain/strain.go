package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var ans Ints
	for _, val := range i {
		if filterInt(val) {
			ans = append(ans, val)
		}
	}
	return ans
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var ans Ints
	for _, val := range i {
		if !filter(val) {
			ans = append(ans, val)
		}
	}
	return ans
}

func (l Lists) Keep(filter func(int) bool) Lists {
	var ans Lists
	for _, val := range l {
		var arr []int
		for _, item := range val {
			if filter(item) {
				arr = append(arr, item)
			}
		}
		ans = append(ans, arr)
	}
	return ans
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var ans Strings
	for _, val := range s {
		if filter(val) {
			ans = append(ans, val)
		}
	}
	return ans
}

func filterInt(i int) bool {
	return i%2 == 0
}

func filterString(s string) bool {
	return len(s) > 6
}
