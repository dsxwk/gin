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
