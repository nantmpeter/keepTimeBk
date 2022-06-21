package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppCaseItemsRouter struct {
}

// InitAppCaseItemsRouter 初始化 AppCaseItems 路由信息
func (s *AppCaseItemsRouter) InitAppCaseItemsRouter(Router *gin.RouterGroup) {
	appCaseItemsRouter := Router.Group("appCaseItems").Use(middleware.OperationRecord())
	appCaseItemsRouterWithoutRecord := Router.Group("appCaseItems")
	var appCaseItemsApi = v1.ApiGroupApp.AppApiGroup.AppCaseItemsApi
	{
		appCaseItemsRouter.POST("createAppCaseItems", appCaseItemsApi.CreateAppCaseItems)   // 新建AppCaseItems
		appCaseItemsRouter.DELETE("deleteAppCaseItems", appCaseItemsApi.DeleteAppCaseItems) // 删除AppCaseItems
		appCaseItemsRouter.DELETE("deleteAppCaseItemsByIds", appCaseItemsApi.DeleteAppCaseItemsByIds) // 批量删除AppCaseItems
		appCaseItemsRouter.PUT("updateAppCaseItems", appCaseItemsApi.UpdateAppCaseItems)    // 更新AppCaseItems
	}
	{
		appCaseItemsRouterWithoutRecord.GET("findAppCaseItems", appCaseItemsApi.FindAppCaseItems)        // 根据ID获取AppCaseItems
		appCaseItemsRouterWithoutRecord.GET("getAppCaseItemsList", appCaseItemsApi.GetAppCaseItemsList)  // 获取AppCaseItems列表
	}
}
