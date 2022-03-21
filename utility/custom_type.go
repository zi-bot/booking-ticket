package utility

import (
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) MarshalJSON() ([]byte, error) {
	date := d.Time.Format("2006-01-02")
	date = fmt.Sprintf("\"%s\"", date)
	return []byte(date), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1]
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}
func (d *Date) ToTime() time.Time {
	return d.Time
}
