package app

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
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
func (appCasesService *AppCasesService) DeleteAppCases(appCases app.AppCases) (err error) {
	err = global.GVA_DB.Delete(&appCases).Error
	return err
}

// DeleteAppCasesByIds 批量删除AppCases记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCasesService *AppCasesService) DeleteAppCasesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]app.AppCases{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppCases 更新AppCases记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCasesService *AppCasesService) UpdateAppCases(appCases app.AppCases) (err error) {
	err = global.GVA_DB.Save(&appCases).Error
	return err
}

// GetAppCases 根据id获取AppCases记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCasesService *AppCasesService) GetAppCases(id uint) (appCases app.AppCases, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&appCases).Error
	return
}

// GetAppCasesInfoList 分页获取AppCases记录
// Author [piexlmax](https://github.com/piexlmax)
func (appCasesService *AppCasesService) GetAppCasesInfoList(info appReq.AppCasesSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&app.AppCases{})
	var appCasess []app.AppCases
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&appCasess).Error
	return appCasess, total, err
}

func (appCasesService *AppCasesService) GetData(userId uint, data appReq.CaseData) (cases []app.AppCases, caseItems []app.AppCaseItems, err error) {
	var caseIds []uint
	var caseItemIds []uint
	for _, value := range data.Cases {
		caseDb := global.GVA_DB.Model(&app.AppCases{})
		value.UserId = &userId
		var r *app.AppCases
		r = &value
		err := caseDb.Where("name = ? and icon = ?", r.Name, r.Icon).Debug().First(&r).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			caseDb.Create(&r)
		}
		caseIds = append(caseIds, r.ID)
	}
	for _, itemValue := range data.CaseItems {
		caseItemDb := global.GVA_DB.Model(&app.AppCaseItems{})
		itemValue.UserId = &userId
		var r *app.AppCaseItems
		r = &itemValue
		err := caseItemDb.Where("user_id = ? and case_id = ? and start = ? and date = ? and end = ?", r.UserId, r.CaseId, r.Start, r.Date, r.End).Debug().First(&r).Error

		if errors.Is(err, gorm.ErrRecordNotFound) {
			caseItemDb.Create(&r)
		}
		//caseItemDb.Create(&itemValue)
		caseItemIds = append(caseItemIds, r.ID)
		global.GVA_LOG.Info(string(itemValue.ID) + "sdhskjdhfskdfhksjd")
	}

	caseDb := global.GVA_DB.Model(&app.AppCases{})
	caseItemDb := global.GVA_DB.Model(&app.AppCaseItems{})
	caseDb.Where("user_id = ?", userId).Not(caseIds)
	caseItemDb.Where("user_id = ?", userId).Not(caseItemIds)

	err = caseDb.Debug().Find(&cases).Error
	err = caseItemDb.Debug().Find(&caseItems).Error
	return cases, caseItems, err
}
