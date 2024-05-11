package helpers

import (
	"time"

	"gorm.io/datatypes"
)

// FormatDate formats the given date in the specified layout. If layout is empty, it uses the default layout "2006-01-02".
func FormatDate(date datatypes.Date, layout string) datatypes.Date {
	if layout == "" {
		layout = "2006-01-02"
	}

	// Convert datatypes.Date to time.Time
	timeDate := time.Time(date)

	// Format the date
	formattedDate := timeDate.Format(layout)

	// Parse the formatted date back to time.Time
	parsedDate, err := time.Parse(layout, formattedDate)
	ErrorPanic(err)

	// Convert time.Time back to datatypes.Date
	return datatypes.Date(parsedDate)
}
