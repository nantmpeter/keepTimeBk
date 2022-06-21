package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	appReq "github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
	appRes "github.com/flipped-aurora/gin-vue-admin/server/model/app/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type AppUsersApi struct {
}

var appUsersService = service.ServiceGroupApp.AppServiceGroup.AppUsersService
var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

// CreateAppUsers 创建AppUsers
// @Tags AppUsers
// @Summary 创建AppUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppUsers true "创建AppUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUsers/createAppUsers [post]
func (appUsersApi *AppUsersApi) CreateAppUsers(c *gin.Context) {
	var appUsers app.AppUsers
	_ = c.ShouldBindJSON(&appUsers)
	if err := appUsersService.CreateAppUsers(appUsers); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppUsers 删除AppUsers
// @Tags AppUsers
// @Summary 删除AppUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppUsers true "删除AppUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUsers/deleteAppUsers [delete]
func (appUsersApi *AppUsersApi) DeleteAppUsers(c *gin.Context) {
	var appUsers app.AppUsers
	_ = c.ShouldBindJSON(&appUsers)
	if err := appUsersService.DeleteAppUsers(appUsers); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppUsersByIds 批量删除AppUsers
// @Tags AppUsers
// @Summary 批量删除AppUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appUsers/deleteAppUsersByIds [delete]
func (appUsersApi *AppUsersApi) DeleteAppUsersByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := appUsersService.DeleteAppUsersByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppUsers 更新AppUsers
// @Tags AppUsers
// @Summary 更新AppUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body app.AppUsers true "更新AppUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appUsers/updateAppUsers [put]
func (appUsersApi *AppUsersApi) UpdateAppUsers(c *gin.Context) {
	var appUsers app.AppUsers
	_ = c.ShouldBindJSON(&appUsers)
	if err := appUsersService.UpdateAppUsers(appUsers); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppUsers 用id查询AppUsers
// @Tags AppUsers
// @Summary 用id查询AppUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query app.AppUsers true "用id查询AppUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appUsers/findAppUsers [get]
func (appUsersApi *AppUsersApi) FindAppUsers(c *gin.Context) {
	var appUsers app.AppUsers
	_ = c.ShouldBindQuery(&appUsers)
	if reappUsers, err := appUsersService.GetAppUsers(appUsers.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappUsers": reappUsers}, c)
	}
}

// GetAppUsersList 分页获取AppUsers列表
// @Tags AppUsers
// @Summary 分页获取AppUsers列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query appReq.AppUsersSearch true "分页获取AppUsers列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUsers/getAppUsersList [get]
func (appUsersApi *AppUsersApi) GetAppUsersList(c *gin.Context) {
	var pageInfo appReq.AppUsersSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := appUsersService.GetAppUsersInfoList(pageInfo); err != nil {
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

func (appUsersApi *AppUsersApi) Login(c *gin.Context) {
	var loginReq appReq.AppUserLogin
	_ = c.ShouldBindJSON(&loginReq)
	if err := utils.Verify(loginReq, utils.AppLoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		u := &app.AppUsers{Username: loginReq.Email, Password: loginReq.Password}
		if user, err := appUsersService.Login(u); err != nil {
			global.GVA_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
		} else {
			appUsersApi.TokenNext(c, *user)
		}
	}
}

func (appUserApi *AppUsersApi) Register(c *gin.Context) {
	var r appReq.AppUserRegister
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &app.AppUsers{Username: r.Username, NickName: r.NickName, Password: r.Password, HeaderImg: r.HeaderImg, Email: r.Email}
	userReturn, err := appUsersService.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(appRes.AppUserResponse{User: userReturn}, err.Error(), c)
	} else {
		appUserApi.TokenNext(c, *user)
		response.OkWithDetailed(appRes.AppUserResponse{User: userReturn}, "注册成功", c)
	}
}

// 登录以后签发jwt
func (b *AppUsersApi) TokenNext(c *gin.Context, user app.AppUsers) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:        user.Uuid,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	showUser := app.AppUserShow{
		Username:  user.Username,
		NickName:  user.NickName,
		Phone:     user.Phone,
		Email:     user.Email,
		HeaderImg: user.HeaderImg,
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(appRes.LoginResponse{
			User:      showUser,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(appRes.LoginResponse{
			User:      showUser,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(appRes.LoginResponse{
			User:      showUser,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

func (appUsersApi *AppUsersApi) Data(c *gin.Context) {
	var caseData appReq.CaseData
	_ = c.ShouldBindJSON(&caseData)
	response.OkWithData(gin.H{"CaseData": caseData}, c)
}
