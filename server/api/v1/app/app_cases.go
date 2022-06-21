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

type AppCasesApi struct {
}

var appCasesService = service.ServiceGroupApp.AppServiceGroup.AppCasesService


// CreateAppCases 创建AppCases
// @Tags AppCases
// @Summary 创建AppCases
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppCases true "创建AppCases"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCases/createAppCases [post]
func (appCasesApi *AppCasesApi) CreateAppCases(c *gin.Context) {
	var appCases app.AppCases
	_ = c.ShouldBindJSON(&appCases)
	if err := appCasesService.CreateAppCases(appCases); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppCases 删除AppCases
// @Tags AppCases
// @Summary 删除AppCases
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppCases true "删除AppCases"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appCases/deleteAppCases [delete]
func (appCasesApi *AppCasesApi) DeleteAppCases(c *gin.Context) {
	var appCases app.AppCases
	_ = c.ShouldBindJSON(&appCases)
	if err := appCasesService.DeleteAppCases(appCases); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppCasesByIds 批量删除AppCases
// @Tags AppCases
// @Summary 批量删除AppCases
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppCases"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appCases/deleteAppCasesByIds [delete]
func (appCasesApi *AppCasesApi) DeleteAppCasesByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := appCasesService.DeleteAppCasesByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppCases 更新AppCases
// @Tags AppCases
// @Summary 更新AppCases
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppCases true "更新AppCases"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appCases/updateAppCases [put]
func (appCasesApi *AppCasesApi) UpdateAppCases(c *gin.Context) {
	var appCases app.AppCases
	_ = c.ShouldBindJSON(&appCases)
	if err := appCasesService.UpdateAppCases(appCases); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppCases 用id查询AppCases
// @Tags AppCases
// @Summary 用id查询AppCases
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppCases true "用id查询AppCases"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appCases/findAppCases [get]
func (appCasesApi *AppCasesApi) FindAppCases(c *gin.Context) {
	var appCases app.AppCases
	_ = c.ShouldBindQuery(&appCases)
	if reappCases, err := appCasesService.GetAppCases(appCases.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappCases": reappCases}, c)
	}
}

// GetAppCasesList 分页获取AppCases列表
// @Tags AppCases
// @Summary 分页获取AppCases列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppCasesSearch true "分页获取AppCases列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCases/getAppCasesList [get]
func (appCasesApi *AppCasesApi) GetAppCasesList(c *gin.Context) {
	var pageInfo appReq.AppCasesSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := appCasesService.GetAppCasesInfoList(pageInfo); err != nil {
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
