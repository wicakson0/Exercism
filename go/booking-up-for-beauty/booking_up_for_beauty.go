package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/02/2006 15:04:05", date)

	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	now := time.Now()
	dateToCompare, _ := time.Parse("January 2, 2006 15:04:05", date)

	return dateToCompare.Before(now)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	dateString, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)

	return (dateString.Hour() >= 12) && (dateString.Hour() < 18)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	appointmentTime, _ := time.Parse("1/2/2006 15:04:05", date)

	return fmt.Sprintf(
		"You have an appointment on %s, %s %d, %d, at %d:%d.",
		appointmentTime.Weekday(),
		appointmentTime.Month(),
		appointmentTime.Day(),
		appointmentTime.Year(),
		appointmentTime.Hour(),
		appointmentTime.Minute(),
	)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	nowTime := time.Now()

	return time.Date(nowTime.Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
