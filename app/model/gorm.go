package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type JsonTime time.Time

// MarshalJSON 模型时间格式化公共方法
func (t *JsonTime) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte(`""`), nil
	}
	formatted := fmt.Sprintf("\"%s\"", time.Time(*t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t JsonTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t *JsonTime) Scan(value interface{}) error {
	if value == nil {
		*t = JsonTime(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*t = JsonTime(v)
		return nil
	case []byte:
		tt, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			return err
		}
		*t = JsonTime(tt)
		return nil
	case string:
		tt, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			return err
		}
		*t = JsonTime(tt)
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

type JsonString []string

func (j JsonString) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JsonString) Scan(value interface{}) error {
	if value == nil {
		*j = JsonString{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return fmt.Errorf("cannot scan %T into JsonString", value)
	}
	return json.Unmarshal(bytes, j)
}

type JsonInt64 []int64

func (j JsonInt64) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (j *JsonInt64) Scan(value interface{}) error {
	if value == nil {
		*j = JsonInt64{}
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case string:
		bytes = []byte(v)
	case []byte:
		bytes = v
	default:
		return fmt.Errorf("cannot scan %T into JsonInt64", value)
	}
	return json.Unmarshal(bytes, j)
}
