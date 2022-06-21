package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/app"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type AppCaseItemsApi struct {
}

var appCaseItemsService = service.ServiceGroupApp.AppServiceGroup.AppCaseItemsService


// CreateAppCaseItems 创建AppCaseItems
// @Tags AppCaseItems
// @Summary 创建AppCaseItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppCaseItems true "创建AppCaseItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCaseItems/createAppCaseItems [post]
func (appCaseItemsApi *AppCaseItemsApi) CreateAppCaseItems(c *gin.Context) {
	var appCaseItems app.AppCaseItems
	_ = c.ShouldBindJSON(&appCaseItems)
	if err := appCaseItemsService.CreateAppCaseItems(appCaseItems); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppCaseItems 删除AppCaseItems
// @Tags AppCaseItems
// @Summary 删除AppCaseItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppCaseItems true "删除AppCaseItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appCaseItems/deleteAppCaseItems [delete]
func (appCaseItemsApi *AppCaseItemsApi) DeleteAppCaseItems(c *gin.Context) {
	var appCaseItems app.AppCaseItems
	_ = c.ShouldBindJSON(&appCaseItems)
	if err := appCaseItemsService.DeleteAppCaseItems(appCaseItems); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppCaseItemsByIds 批量删除AppCaseItems
// @Tags AppCaseItems
// @Summary 批量删除AppCaseItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppCaseItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appCaseItems/deleteAppCaseItemsByIds [delete]
func (appCaseItemsApi *AppCaseItemsApi) DeleteAppCaseItemsByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := appCaseItemsService.DeleteAppCaseItemsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppCaseItems 更新AppCaseItems
// @Tags AppCaseItems
// @Summary 更新AppCaseItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppCaseItems true "更新AppCaseItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appCaseItems/updateAppCaseItems [put]
func (appCaseItemsApi *AppCaseItemsApi) UpdateAppCaseItems(c *gin.Context) {
	var appCaseItems app.AppCaseItems
	_ = c.ShouldBindJSON(&appCaseItems)
	if err := appCaseItemsService.UpdateAppCaseItems(appCaseItems); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppCaseItems 用id查询AppCaseItems
// @Tags AppCaseItems
// @Summary 用id查询AppCaseItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppCaseItems true "用id查询AppCaseItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appCaseItems/findAppCaseItems [get]
func (appCaseItemsApi *AppCaseItemsApi) FindAppCaseItems(c *gin.Context) {
	var appCaseItems app.AppCaseItems
	_ = c.ShouldBindQuery(&appCaseItems)
	if reappCaseItems, err := appCaseItemsService.GetAppCaseItems(appCaseItems.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappCaseItems": reappCaseItems}, c)
	}
}

// GetAppCaseItemsList 分页获取AppCaseItems列表
// @Tags AppCaseItems
// @Summary 分页获取AppCaseItems列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppCaseItemsSearch true "分页获取AppCaseItems列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCaseItems/getAppCaseItemsList [get]
func (appCaseItemsApi *AppCaseItemsApi) GetAppCaseItemsList(c *gin.Context) {
	var pageInfo appReq.AppCaseItemsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := appCaseItemsService.GetAppCaseItemsInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
