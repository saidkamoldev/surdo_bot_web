package i18n

import (
	"fmt"
	"surdo-bot/internal/i18n/lng"
	"time"
)

var weekdays = map[time.Weekday]map[lng.Language]string{
	time.Monday: {
		lng.RUSSIAN: "Понедельник",
	},
	time.Tuesday: {
		lng.RUSSIAN: "Вторник",
	},
	time.Wednesday: {
		lng.RUSSIAN: "Среда",
	},
	time.Thursday: {
		lng.RUSSIAN: "Четверг",
	},
	time.Friday: {
		lng.RUSSIAN: "Пятница",
	},
	time.Saturday: {
		lng.RUSSIAN: "Суббота",
	},
	time.Sunday: {
		lng.RUSSIAN: "Воскресенье",
	},
}

var today = map[lng.Language]string{
	lng.RUSSIAN: "Сегодня",
}
var tomorrow = map[lng.Language]string{
	lng.RUSSIAN: "Завтра",
}

func FormatDate(date time.Time, lang lng.Language) string {
	addition := ""

	if isToday(date) {
		addition = "(" + today[lang] + ")"
	}
	if isTomorrow(date) {
		addition = "(" + tomorrow[lang] + ")"
	}
	return fmt.Sprintf("%s %s %s", weekdays[date.Weekday()][lang], date.Format("02.01.2006"), addition)
}

func isTomorrow(d time.Time) bool {
	year1, month1, day1 := d.Date()
	year2, month2, day2 := time.Now().Date()
	return year1 == year2 && month1 == month2 && day1 == day2+1
}

func isToday(d time.Time) bool {
	return EqDate(d, time.Now())
}

func EqDate(d1, d2 time.Time) bool {
	year1, month1, day1 := d1.Date()
	year2, month2, day2 := d2.Date()
	return year1 == year2 && month1 == month2 && day1 == day2
}
