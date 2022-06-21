package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AppUsersRouter struct {
}

// InitAppUsersRouter 初始化 AppUsers 路由信息
func (s *AppUsersRouter) InitAppUsersRouter(Router *gin.RouterGroup) {
	appUsersRouter := Router.Group("appUsers").Use(middleware.OperationRecord())
	appUsersRouterWithoutRecord := Router.Group("appUsers")
	var appUsersApi = v1.ApiGroupApp.AppApiGroup.AppUsersApi
	{
		appUsersRouter.POST("createAppUsers", appUsersApi.CreateAppUsers)             // 新建AppUsers
		appUsersRouter.DELETE("deleteAppUsers", appUsersApi.DeleteAppUsers)           // 删除AppUsers
		appUsersRouter.DELETE("deleteAppUsersByIds", appUsersApi.DeleteAppUsersByIds) // 批量删除AppUsers
		appUsersRouter.PUT("updateAppUsers", appUsersApi.UpdateAppUsers)              // 更新AppUsers
		appUsersRouter.POST("data", appUsersApi.Data)
		//appUsersRouter.POST("login", appUsersApi.Login)
		//appUsersRouter.POST("register", appUsersApi.Register)
	}
	{
		appUsersRouterWithoutRecord.POST("register", appUsersApi.Register)
		appUsersRouterWithoutRecord.GET("findAppUsers", appUsersApi.FindAppUsers)       // 根据ID获取AppUsers
		appUsersRouterWithoutRecord.GET("getAppUsersList", appUsersApi.GetAppUsersList) // 获取AppUsers列表
	}
}
