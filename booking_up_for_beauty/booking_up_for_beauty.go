package booking_up_for_beauty

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func Schedule(date string) time.Time {
	res1 := strings.Split(date, " ")
	res2 := strings.Split(res1[0], "/")
	day := res2[1]
	month := res2[0]
	year := res2[2]
	int_month, _ := strconv.Atoi(res2[0])

	if int_month < 10 {
		month = "0" + month
	}

	revisedDate := year + "-" + month + "-" + day + "T" + res1[1] + "Z"
	tm, _ := time.Parse(time.RFC3339, revisedDate)
	return tm
}

func HasPassed(date string) bool {
	actualDate := Schedule(date)
	today := time.Now()
	return actualDate.Before(today)
}

func IsAfternoonAppointment(date string) bool {
	appointmentDate := Schedule(date)
	hour := appointmentDate.Hour()
	return hour >= 12 && hour <= 18
}

func Description(date string) string {
	actualDate := Schedule(date)
	ans := fmt.Sprintf("You have an appointment on %v, %v %v, at %v:%v.", actualDate.Weekday(), actualDate.Month(), actualDate.Day(), actualDate.Hour(), actualDate.Minute())
	return ans
}

func AnniversaryDate() time.Time {
	return time.Date(2020, time.September, 5, 0, 0, 0, 0, time.UTC)
}
