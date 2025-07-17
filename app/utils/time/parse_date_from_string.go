package timeutil

import "time"

func ParseDateFromString(dateStr string) (*time.Time, error) {
	if dateStr == "" {
		return nil, nil
	}
	
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, err
	}

	return &date, nil
}
