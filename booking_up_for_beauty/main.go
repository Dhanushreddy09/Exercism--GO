package booking_up_for_beauty

import "fmt"

func Booking_up_for_beauty() {
	fmt.Printf("Date is %v\n", Schedule("7/25/2019 13:45:00"))
	fmt.Printf("Appointment date has passed : %v\n", HasPassed("7/25/2019 13:45:00"))
	fmt.Printf("Appointment date is after noon : %v\n", IsAfternoonAppointment("7/25/2019 19:45:00"))
	fmt.Printf("Description : %v\n", Description("7/25/2019 13:45:00"))
	fmt.Printf("Anniversary date : %v\n", AnniversaryDate())
}
