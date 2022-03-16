package twelve_days

import "fmt"

var gifts = map[int]string{
	1:  "a Partridge in a Pear Tree.",
	2:  "two Turtle Doves, and ",
	3:  "three French Hens, ",
	4:  "four Calling Birds, ",
	5:  "five Gold Rings, ",
	6:  "six Geese-a-Laying, ",
	7:  "seven Swans-a-Swimming, ",
	8:  "eight Maids-a-Milking, ",
	9:  "nine Ladies Dancing, ",
	10: "ten Lords-a-Leaping, ",
	11: "eleven Pipers Piping, ",
	12: "twelve Drummers Drumming, ",
}

func Verse(i int) string {
	var day string
	var verseItems string
	switch i {
	case 1:
		day = "first"
	case 2:
		day = "second"
	case 3:
		day = "third"
	case 4:
		day = "fourth"
	case 5:
		day = "fifth"
	case 6:
		day = "sixth"
	case 7:
		day = "seventh"
	case 8:
		day = "eighth"
	case 9:
		day = "ninth"
	case 10:
		day = "tenth"
	case 11:
		day = "eleventh"
	case 12:
		day = "twelfth"
	}
	for j := i; j > 0; j-- {
		verseItems += gifts[j]
	}
	return "On the " + day + " day of Christmas my true love gave to me: " + verseItems
}

func Song() string {
	var ans string
	for i := 1; i <= 12; i++ {
		verse := Verse(i)
		if i == 12 {
			ans += fmt.Sprint(verse)
		} else {
			ans += fmt.Sprintln(verse)
		}

	}
	return ans
}
