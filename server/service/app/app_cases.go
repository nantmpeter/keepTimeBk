package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
)

type AppCasesService struct {
}

// CreateAppCases 创建AppCases记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCasesService *AppCasesService) CreateAppCases(appCases app.AppCases) (err error) {
	err = global.GVA_DB.Create(&appCases).Error
	return err
}

// DeleteAppCases 删除AppCases记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCasesService *AppCasesService)DeleteAppCases(appCases app.AppCases) (err error) {
	err = global.GVA_DB.Delete(&appCases).Error
	return err
}

// DeleteAppCasesByIds 批量删除AppCases记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCasesService *AppCasesService)DeleteAppCasesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppCases{},"id in ?",ids.Ids).Error
	return err
}

// UpdateAppCases 更新AppCases记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCasesService *AppCasesService)UpdateAppCases(appCases app.AppCases) (err error) {
	err = global.GVA_DB.Save(&appCases).Error
	return err
}

// GetAppCases 根据id获取AppCases记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCasesService *AppCasesService)GetAppCases(id uint) (appCases app.AppCases, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appCases).Error
	return
}

// GetAppCasesInfoList 分页获取AppCases记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCasesService *AppCasesService)GetAppCasesInfoList(info appReq.AppCasesSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&app.AppCases{})
    var appCasess []app.AppCases
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&appCasess).Error
	return  appCasess, total, err
}
