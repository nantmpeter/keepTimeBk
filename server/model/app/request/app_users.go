package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppUsersSearch struct {
	app.AppUsers
	request.PageInfo
}

type AppUserLogin struct {
	Email string `json:"email"` // 邮箱
	//Username string `json:"email"`    // 邮箱
	Password string `json:"password"` // 密码
	//Captcha   string `json:"captcha"`   // 验证码
	//CaptchaId string `json:"captchaId"` // 验证码ID
}

type AppUserRegister struct {
	Username  string `json:"userName"`
	Email     string `json:"email"`
	Password  string `json:"passWord"`
	NickName  string `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg string `json:"headerImg" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"`
	//AuthorityId  string   `json:"authorityId" gorm:"default:888"`
	//AuthorityIds []string `json:"authorityIds"`
}

type CaseData struct {
	Cases     []app.AppCases
	CaseItems []app.AppCaseItems
}
