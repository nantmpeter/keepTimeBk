// 自动生成模板AppUsers
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	uuid "github.com/satori/go.uuid"
)

// AppUsers 结构体
// 如果含有time.Time 请自行import time包
type AppUsers struct {
	global.GVA_MODEL
	Uuid        uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户UUID;size:191;"`
	Username    string    `json:"username" form:"username" gorm:"column:username;comment:用户登录名;size:191;"`
	Password    string    `json:"password" form:"password" gorm:"column:password;comment:用户登录密码;size:191;"`
	NickName    string    `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;size:191;"`
	HeaderImg   string    `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:用户头像;size:191;"`
	AuthorityId string    `json:"authorityId" form:"authorityId" gorm:"column:authority_id;comment:用户角色ID;size:191;"`
	Phone       string    `json:"phone" form:"phone" gorm:"column:phone;comment:用户手机号;size:191;"`
	Email       string    `json:"email" form:"email" gorm:"column:email;comment:用户邮箱;size:191;"`
}

type AppUserShow struct {
	Username  string `json:"username" form:"username" gorm:"column:username;comment:用户登录名;size:191;"`
	Email     string `json:"email" form:"email" gorm:"column:email;comment:用户邮箱;size:191;"`
	Phone     string `json:"phone" form:"phone" gorm:"column:phone;comment:用户手机号;size:191;"`
	NickName  string `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;size:191;"`
	HeaderImg string `json:"headerImg" form:"headerImg" gorm:"column:header_img;comment:用户头像;size:191;"`
}

// TableName AppUsers 表名
func (AppUsers) TableName() string {
	return "app_users"
}
