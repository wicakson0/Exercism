package expenses

import (
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var out []Record

	for _, val := range in {
		if predicate(val) {
			out = append(out, val)
		}
	}

	return out
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	low := p.From
	high := p.To

	return func(r Record) bool {
		if r.Day <= high && r.Day >= low {
			return true
		}

		return false
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	runningTotal := 0.0
	periodChecker := ByDaysPeriod(p)

	for _, val := range in {
		if periodChecker(val) {
			runningTotal += val.Amount
		}
	}

	return float64(runningTotal)
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	categoryChecker := ByCategory(c)
	categoryChecked := Filter(in, categoryChecker)

	if categoryChecked == nil {
		return 0.0, fmt.Errorf("error(unknown category %s)", c)
	}

	runningTotal := TotalByPeriod(categoryChecked, p)

	return runningTotal, nil
}
