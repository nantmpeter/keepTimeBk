// 自动生成模板AppCaseItems
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppCaseItems 结构体
// 如果含有time.Time 请自行import time包
type AppCaseItems struct {
	global.GVA_MODEL
	CaseId *uint   `json:"caseId" form:"caseId" gorm:"column:case_id;comment:;size:10;" json:"case_id,omitempty"`
	UserId *uint   `json:"userId" form:"userId" gorm:"column:user_id;comment:;size:10;" json:"user_id,omitempty"`
	Date   *string `json:"date" form:"date" gorm:"column:date;comment:;size:30;" json:"date,omitempty"`
	Start  *string `json:"start" form:"start" gorm:"column:start;comment:;size:30;" json:"start,omitempty"`
	End    *string `json:"end" form:"end" gorm:"column:end;comment:;size:30;" json:"end,omitempty"`
	//Date   *global.XTime
	//Start  *global.XTime
	//End    *global.XTime
}

// TableName AppCaseItems 表名
func (AppCaseItems) TableName() string {
	return "app_case_items"
}
