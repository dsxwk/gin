package model

import (
	"database/sql/driver"
	"fmt"
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

// DeletedAt 自定义删除时间类型，用于 swagger 展示
type DeletedAt struct {
	Time  *time.Time `json:"time,omitempty"`
	Valid bool       `json:"valid"`
}

// Scan 实现 sql.Scanner
func (d *DeletedAt) Scan(value interface{}) error {
	if value == nil {
		d.Time = nil
		d.Valid = false
		return nil
	}

	switch v := value.(type) {
	case time.Time:
		d.Time = &v
		d.Valid = true
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}

	return nil
}

// Value 实现 driver.Valuer
func (d DeletedAt) Value() (driver.Value, error) {
	if !d.Valid || d.Time == nil {
		return nil, nil
	}
	return *d.Time, nil
}
