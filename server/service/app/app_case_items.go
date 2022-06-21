package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
)

type AppCaseItemsService struct {
}

// CreateAppCaseItems 创建AppCaseItems记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCaseItemsService *AppCaseItemsService) CreateAppCaseItems(appCaseItems app.AppCaseItems) (err error) {
	err = global.GVA_DB.Create(&appCaseItems).Error
	return err
}

// DeleteAppCaseItems 删除AppCaseItems记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCaseItemsService *AppCaseItemsService)DeleteAppCaseItems(appCaseItems app.AppCaseItems) (err error) {
	err = global.GVA_DB.Delete(&appCaseItems).Error
	return err
}

// DeleteAppCaseItemsByIds 批量删除AppCaseItems记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCaseItemsService *AppCaseItemsService)DeleteAppCaseItemsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppCaseItems{},"id in ?",ids.Ids).Error
	return err
}

// UpdateAppCaseItems 更新AppCaseItems记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCaseItemsService *AppCaseItemsService)UpdateAppCaseItems(appCaseItems app.AppCaseItems) (err error) {
	err = global.GVA_DB.Save(&appCaseItems).Error
	return err
}

// GetAppCaseItems 根据id获取AppCaseItems记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCaseItemsService *AppCaseItemsService)GetAppCaseItems(id uint) (appCaseItems app.AppCaseItems, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appCaseItems).Error
	return
}

// GetAppCaseItemsInfoList 分页获取AppCaseItems记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCaseItemsService *AppCaseItemsService)GetAppCaseItemsInfoList(info appReq.AppCaseItemsSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&app.AppCaseItems{})
    var appCaseItemss []app.AppCaseItems
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&appCaseItemss).Error
	return  appCaseItemss, total, err
}
