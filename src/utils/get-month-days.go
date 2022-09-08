package utils

import (
	"strconv"
	"strings"
)

var (
	days28 = ranger(28)
	days29 = ranger(29)
	days30 = ranger(30)
	days31 = ranger(31)
)

func ranger(n uint8) []uint8 {
	list := make([]uint8, n)

	var index uint8
	for index = 0; index < n; index++ {
		list[index] = index + 1
	}

	return list
}

func GetMonthDays(m string) []uint8 {
	splittedList := strings.Split(m, "-")
	month, _ := strconv.Atoi(splittedList[0])
	year, _ := strconv.Atoi(splittedList[1])

	switch month {
	case 2:
		if year%4 == 0 {
			return days29
		}
		return days28

	case 1, 3, 5, 7, 8, 10, 12:
		return days31
	default:
		return days30
	}
}
