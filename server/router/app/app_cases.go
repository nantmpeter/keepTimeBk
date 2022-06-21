package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppCasesRouter struct {
}

// InitAppCasesRouter 初始化 AppCases 路由信息
func (s *AppCasesRouter) InitAppCasesRouter(Router *gin.RouterGroup) {
	appCasesRouter := Router.Group("appCases").Use(middleware.OperationRecord())
	appCasesRouterWithoutRecord := Router.Group("appCases")
	var appCasesApi = v1.ApiGroupApp.AppApiGroup.AppCasesApi
	{
		appCasesRouter.POST("createAppCases", appCasesApi.CreateAppCases)   // 新建AppCases
		appCasesRouter.DELETE("deleteAppCases", appCasesApi.DeleteAppCases) // 删除AppCases
		appCasesRouter.DELETE("deleteAppCasesByIds", appCasesApi.DeleteAppCasesByIds) // 批量删除AppCases
		appCasesRouter.PUT("updateAppCases", appCasesApi.UpdateAppCases)    // 更新AppCases
	}
	{
		appCasesRouterWithoutRecord.GET("findAppCases", appCasesApi.FindAppCases)        // 根据ID获取AppCases
		appCasesRouterWithoutRecord.GET("getAppCasesList", appCasesApi.GetAppCasesList)  // 获取AppCases列表
	}
}
