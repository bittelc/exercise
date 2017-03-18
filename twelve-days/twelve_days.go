// Package twelve returns a single verse for any of the twelve days of Christmas
package twelve

import (
	"fmt"
)

const testVersion = 1

type dailyGift struct {
	day  string
	gift string
}

var twelveGifts = []dailyGift{
	{"first", "a Partridge in a Pear Tree."},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

// Verse returns a single verse corresponding to the i'th numbered verse
func Verse(i int) string {
	verse := verbiage(twelveGifts[i-1])

	for j := i; j >= 2; j-- {
		verse += " " + twelveGifts[j-1].gift + ","
	}
	if i != 1 {
		verse += " and " + twelveGifts[0].gift
	} else {
		verse += " " + twelveGifts[0].gift
	}

	return verse
}

// Song returns the entire song, line seperating each verse
func Song() string {

	var song string
	for i := 1; i <= 12; i++ {
		song += Verse(i) + "\n"
	}
	return song
}

// verbiage provides the initial entry verbiage into the 12 days of Christmas song
func verbiage(daysGift dailyGift) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me,", daysGift.day)
}
