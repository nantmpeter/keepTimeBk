// 自动生成模板AppCaseItems
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// AppCaseItems 结构体
// 如果含有time.Time 请自行import time包
type AppCaseItems struct {
	global.GVA_MODEL
	CaseId *int       `json:"caseId" form:"caseId" gorm:"column:case_id;comment:;size:10;"`
	UserId *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:;size:10;"`
	Date   *time.Time `json:"date" form:"date" gorm:"column:date;comment:;"`
	Start  *time.Time `json:"start" form:"start" gorm:"column:start;comment:;"`
	End    *time.Time `json:"end" form:"end" gorm:"column:end;comment:;"`
}

// TableName AppCaseItems 表名
func (AppCaseItems) TableName() string {
	return "app_case_items"
}
