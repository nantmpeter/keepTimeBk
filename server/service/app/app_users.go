package app

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type AppUsersService struct {
}

// CreateAppUsers 创建AppUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUsersService *AppUsersService) CreateAppUsers(appUsers app.AppUsers) (err error) {
	err = global.GVA_DB.Create(&appUsers).Error
	return err
}

// DeleteAppUsers 删除AppUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUsersService *AppUsersService) DeleteAppUsers(appUsers app.AppUsers, userId uint) (err error) {
	err = global.GVA_DB.Debug().Where("1 = 1").Delete(&appUsers, userId).Error
	return err
}

// DeleteAppUsersByIds 批量删除AppUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUsersService *AppUsersService) DeleteAppUsersByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppUsers{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppUsers 更新AppUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUsersService *AppUsersService) UpdateAppUsers(appUsers app.AppUsers) (err error) {
	err = global.GVA_DB.Save(&appUsers).Error
	return err
}

// GetAppUsers 根据id获取AppUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUsersService *AppUsersService) GetAppUsers(id uint) (appUsers app.AppUsers, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appUsers).Error
	return
}

// GetAppUsersInfoList 分页获取AppUsers记录
// Author [piexlmax](https://github.com/piexlmax)
func (appUsersService *AppUsersService) GetAppUsersInfoList(info appReq.AppUsersSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppUsers{})
	var appUserss []app.AppUsers
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appUserss).Error
	return appUserss, total, err
}

func (appUserService *AppUsersService) Register(u app.AppUsers) (userInter app.AppUsers, err error) {
	var user app.AppUsers
	if !errors.Is(global.GVA_DB.Where("email = ?", u.Email).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("该邮箱已注册")
	}
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.Uuid = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return u, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Login
//@description: 用户登录
//@param: u *model.SysUser
//@return: err error, userInter *model.SysUser

func (appUserService *AppUsersService) Login(u *app.AppUsers) (userInter *app.AppUsers, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user app.AppUsers
	err = global.GVA_DB.Where("email = ?", u.Email).First(&user).Error
	global.GVA_LOG.Info("upwd:" + u.Password + "\n userpwd+ " + user.Password + "\n email: " + u.Email)
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}

	return &user, err
}
