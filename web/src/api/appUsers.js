import service from '@/utils/request'

// @Tags AppUsers
// @Summary 创建AppUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppUsers true "创建AppUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUsers/createAppUsers [post]
export const createAppUsers = (data) => {
  return service({
    url: '/appUsers/createAppUsers',
    method: 'post',
    data
  })
}

// @Tags AppUsers
// @Summary 删除AppUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppUsers true "删除AppUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUsers/deleteAppUsers [delete]
export const deleteAppUsers = (data) => {
  return service({
    url: '/appUsers/deleteAppUsers',
    method: 'delete',
    data
  })
}

// @Tags AppUsers
// @Summary 删除AppUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appUsers/deleteAppUsers [delete]
export const deleteAppUsersByIds = (data) => {
  return service({
    url: '/appUsers/deleteAppUsersByIds',
    method: 'delete',
    data
  })
}

// @Tags AppUsers
// @Summary 更新AppUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppUsers true "更新AppUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appUsers/updateAppUsers [put]
export const updateAppUsers = (data) => {
  return service({
    url: '/appUsers/updateAppUsers',
    method: 'put',
    data
  })
}

// @Tags AppUsers
// @Summary 用id查询AppUsers
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppUsers true "用id查询AppUsers"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appUsers/findAppUsers [get]
export const findAppUsers = (params) => {
  return service({
    url: '/appUsers/findAppUsers',
    method: 'get',
    params
  })
}

// @Tags AppUsers
// @Summary 分页获取AppUsers列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppUsers列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appUsers/getAppUsersList [get]
export const getAppUsersList = (params) => {
  return service({
    url: '/appUsers/getAppUsersList',
    method: 'get',
    params
  })
}
