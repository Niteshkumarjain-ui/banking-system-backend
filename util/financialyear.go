package util

import (
	"time"
)

func GetFinancialYear() (time.Time, time.Time) {
	now := time.Now()
	currentYear := now.Year()

	var startFinancialYear, endFinancialYear time.Time

	if now.Month() < time.April {
		startFinancialYear = time.Date(currentYear-1, time.April, 1, 0, 0, 0, 0, time.UTC)
		endFinancialYear = time.Date(currentYear, time.March, 31, 23, 59, 59, 0, time.UTC)
	} else {
		startFinancialYear = time.Date(currentYear, time.April, 1, 0, 0, 0, 0, time.UTC)
		endFinancialYear = time.Date(currentYear+1, time.March, 31, 23, 59, 59, 0, time.UTC)
	}

	return startFinancialYear, endFinancialYear
}
