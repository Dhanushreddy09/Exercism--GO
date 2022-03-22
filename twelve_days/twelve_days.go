package twelve_days

import "fmt"

func Verse(i int) string {

	var gifts = []string{"a Partridge in a Pear Tree.", "two Turtle Doves, and ", "three French Hens, ", "four Calling Birds, ", "five Gold Rings, ", "six Geese-a-Laying, ", "seven Swans-a-Swimming, ", "eight Maids-a-Milking, ", "nine Ladies Dancing, ", "ten Lords-a-Leaping, ", "eleven Pipers Piping, ", "twelve Drummers Drumming, "}

	var days = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

	var verseItems string
	for j := i; j > 0; j-- {
		verseItems += gifts[j-1]
	}
	return "On the " + days[i-1] + " day of Christmas my true love gave to me: " + verseItems
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
