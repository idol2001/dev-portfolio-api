package models

import (
	"database/sql/driver"
	"dev-portfolio-api/pkg/global"
	"fmt"
	"time"
)

// BaseModel 基座模型（对齐实际数据库：小写列名 + 软删除）
type BaseModel struct {
	ID        uint       `gorm:"column:id;primary_key;autoIncrement;comment:'自增编号'" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at;comment:'创建时间'" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;comment:'更新时间'" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index;comment:'删除时间(软删除)'" json:"deleted_at,omitempty"`
}

// 表名设置
func (BaseModel) TableName(name string) string {
	if global.Conf.Mysql != nil {
		return fmt.Sprintf("%s%s", global.Conf.Mysql.TablePrefix, name)
	}
	if global.Conf.Pgsql != nil {
		return fmt.Sprintf("%s%s", global.Conf.Pgsql.TablePrefix, name)
	}
	return name
}

// 自定义时间 JSON 转换
const TimeFormat = "2006-01-02 15:04:05"

type LocalTime struct {
	time.Time
}

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = LocalTime{Time: time.Time{}}
		return
	}
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = LocalTime{Time: now}
	return
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf("\"%s\"", t.Format(TimeFormat))
	return []byte(output), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (t LocalTime) String() string {
	return t.Format(TimeFormat)
}
