package helper

import (
	"fmt"
	"time"
)

// CalculateAge takes a birth date and returns the age as a string (e.g., "2 years" or "6 months")
func CalculateAge(birthDate time.Time) string {
	now := time.Now()
	years := now.Year() - birthDate.Year()

	if now.YearDay() < birthDate.YearDay() {
		years--
	}

	if years > 0 {
		if years == 1 {
			return "1 year"
		}
		return fmt.Sprintf("%d years", years)
	}

	months := int(now.Sub(birthDate).Hours() / (24 * 30))
	if months == 0 || months == 1 {
		return fmt.Sprintf("%d month", months)
	}
	return fmt.Sprintf("%d months", months)
}
