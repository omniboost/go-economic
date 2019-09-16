package economic

import (
	"encoding/json"
	"strconv"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format("2006-01-02"))
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	// try iso8601 date format
	d.Time, err = time.Parse("2006-01-02", value)
	return err
}

type DateTime struct {
	time.Time
}

func (d *DateTime) MarshalJSON() ([]byte, error) {
	if d.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(d.Time)
}

func (d DateTime) IsEmpty() bool {
	return d.Time.IsZero()
}

func (d *DateTime) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	d.Time, err = time.Parse("2006-01-02 15:04:05", value)
	return err
}

func (d DateTime) MarshalSchema() string {
	return d.Format("2006-01-02 15:04:05")
}

type Int int

func (i *Int) UnmarshalJSON(data []byte) error {
	realInt := 0
	err := json.Unmarshal(data, &realInt)
	if err == nil {
		*i = Int(realInt)
		return nil
	}

	// error, so maybe it isn't an int
	str := ""
	err = json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	if str == "" {
		*i = 0
		return nil
	}

	realInt, err = strconv.Atoi(str)
	if err != nil {
		return err
	}

	i2 := Int(realInt)
	*i = i2
	return nil
}

func (i Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(i))
}

type Bool bool

func (b Bool) MarshalJSON() ([]byte, error) {
	if b == true {
		return json.Marshal("1")
	}
	return json.Marshal("0")
}

func (b Bool) MarshalSchema() string {
	if b == true {
		return "true"
	}
	return "false"
}

type LogTime struct {
	time.Time
}

func (l *LogTime) MarshalJSON() ([]byte, error) {
	if l.Time.IsZero() {
		return json.Marshal(nil)
	}

	return json.Marshal(l.Time)
}

func (l LogTime) IsEmpty() bool {
	return l.Time.IsZero()
}

func (l *LogTime) UnmarshalJSON(text []byte) (err error) {
	var value string
	err = json.Unmarshal(text, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	l.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	l.Time, err = time.Parse("2006-01-02T15:04:05", value)
	return err
}

func (l LogTime) MarshalSchema() string {
	return l.Format("2006-01-02T15:04:05")
}
