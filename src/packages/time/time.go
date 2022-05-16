package timepkg

import "time"

func ParseHour(val string) (*time.Time, error) {
	t1, err := time.Parse("15:00", val)

	if err != nil {
		return nil, err
	}

	return &t1, nil
}

func CheckValidHour(val string) bool {
	_, err := ParseHour(val)

	if err != nil {
		return false
	}

	return true
}

func ParseDate(val string) (*time.Time, error) {
	layout := "02/01/2006"
	t1, err := time.Parse(layout, val)

	if err != nil {
		return nil, err
	}

	return &t1, nil
}
