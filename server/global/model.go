package global

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type GVA_MODEL struct {
	ID        uint `gorm:"primarykey"` // 主键ID
	CreatedAt Time // 创建时间
	//CreatedAt XTime // 创建时间
	UpdatedAt XTime // 创建时间
	DeletedAt XTime // 创建时间
	//UpdatedAt time.Time      // 更新时间
	//DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime{value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)

}

func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixMilli() == zeroTime.UnixMilli() {
		return nil, nil
	}
	return t.Time, nil
}

func (t XTime) MarshalJSON() ([]byte, error) {
	//output := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	output := fmt.Sprintf("%v", t.UnixMilli())
	return []byte(output), nil
}

type XTime struct {
	time.Time
}
