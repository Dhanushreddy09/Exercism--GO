package gross_store

import "fmt"

func Gross_store() {
	units := Units()
	bill := NewBill()
	fmt.Println(bill)
	fmt.Println(AddItem(bill, units, "tomato", "half_of_a_dozen"))
	fmt.Println(bill)
	fmt.Println(RemoveItem(bill, units, "tomato", "quarter_of_a_dozen"))
}
