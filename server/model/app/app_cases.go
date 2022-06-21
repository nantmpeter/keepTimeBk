// 自动生成模板AppCases
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// AppCases 结构体
// 如果含有time.Time 请自行import time包
type AppCases struct {
      global.GVA_MODEL
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:10;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;size:100;"`
      Icon  *int `json:"icon" form:"icon" gorm:"column:icon;comment:图标;size:10;"`
}


// TableName AppCases 表名
func (AppCases) TableName() string {
  return "app_cases"
}

