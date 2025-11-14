package model

import (
	"database/sql/driver"
	"fmt"
	"github.com/goccy/go-json"
	"gorm.io/gorm"
	"time"
)

type DateTime time.Time

func (t *DateTime) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`""`), nil
	}
	formatted := fmt.Sprintf("\"%s\"", time.Time(*t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t DateTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t *DateTime) Scan(value interface{}) error {
	if value == nil {
		*t = DateTime(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*t = DateTime(v)
		return nil
	case []byte:
		tt, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			return err
		}
		*t = DateTime(tt)
		return nil
	case string:
		tt, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return err
		}
		*t = DateTime(tt)
		return nil
	default:
		return fmt.Errorf("cannot convert %v to timestamp", value)
	}
}

type DeletedAt struct {
	gorm.DeletedAt
}

func (d DeletedAt) MarshalJSON() ([]byte, error) {
	if !d.Valid {
		return []byte(`null`), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, d.Time.Format("2006-01-02 15:04:05"))), nil
}

type ArrayString []string

func (j ArrayString) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *ArrayString) Scan(value interface{}) error {
	if value == nil {
		*j = ArrayString{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return fmt.Errorf("cannot scan %T into ArrayString", value)
	}
	return json.Unmarshal(bytes, j)
}

type ArrayInt64 []int64

func (j ArrayInt64) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *ArrayInt64) Scan(value interface{}) error {
	if value == nil {
		*j = ArrayInt64{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return fmt.Errorf("cannot scan %T into ArrayInt64", value)
	}
	return json.Unmarshal(bytes, j)
}
