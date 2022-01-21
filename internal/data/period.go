package data

import (
	"fmt"
	"regexp"
)

// Period represents a time period
type Period string

var dateFormatRegex = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
var dateRangeFormatRegex = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}:\d{4}-\d{2}-\d{2}$`)

func checkDateFormat(date string) bool {
	if !dateFormatRegex.MatchString(date) && !dateRangeFormatRegex.MatchString(date) {
		return false
	}
	return true
}

// NewPeriod returns a Period from a string or an error if an invalid period
func NewPeriod(p string) (Period, error) {
	if checkDateFormat(p) {
		return Period(p), nil
	}
	switch p {
	case "last-week", "last-month", "last-year":
		return Period(p), nil
	default:
		return Period(""), fmt.Errorf("invalid period: %v", p)
	}
}
