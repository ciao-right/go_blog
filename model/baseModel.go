package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type FTime struct {
	time.Time
}
type BaseModel struct {
	ID         uint  `gorm:"primaryKey;not null;autoIncrement" json:"id" `
	CreatedOn  FTime `json:"created_on" gorm:"autoCreateTime"`
	ModifiedOn FTime `json:"modified_on" gorm:"autoUpdateTime"`
}

type PageListModel struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type OverrideTimeModel struct {
	CreatedOn  string `json:"created_on" `
	ModifiedOn string `json:"modified_on"`
}

// MarshaJSON 为 FTime 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t *FTime) MarshaJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	fmt.Println(output)
	return []byte(output), nil
}

// Value 3. 为 FTime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t FTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 4. 为 FTime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *FTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = FTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
