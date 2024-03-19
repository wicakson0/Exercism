package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	number := float64(nb.Number())
	return fmt.Sprintf("This is a box containing the number %.1f", number)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch v := fnb.(type) {
	case FancyNumber:
		num, err := strconv.Atoi(v.Value())
		if err != nil {
			return 0
		}
		return num
	}
	return 0
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	switch v := fnb.(type) {
	case FancyNumber:
		num, err := strconv.Atoi(v.Value())
		if err != nil {
			return fmt.Sprintf("This is a fancy box containing the number 0.0, error: %s", err)
		}
		return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(num))
	}
	return fmt.Sprint("This is a fancy box containing the number 0.0")
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
	case int:
		numFloat := float64(v)
		return DescribeNumber(numFloat)
	case float64:
		return DescribeNumber(v)
	case NumberBox:
		return DescribeNumberBox(v)
	case FancyNumberBox:
		return DescribeFancyNumberBox(v)
	}
	return "Return to sender"
}
