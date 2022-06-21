import service from '@/utils/request'

// @Tags AppCases
// @Summary 创建AppCases
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppCases true "创建AppCases"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCases/createAppCases [post]
export const createAppCases = (data) => {
  return service({
    url: '/appCases/createAppCases',
    method: 'post',
    data
  })
}

// @Tags AppCases
// @Summary 删除AppCases
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppCases true "删除AppCases"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appCases/deleteAppCases [delete]
export const deleteAppCases = (data) => {
  return service({
    url: '/appCases/deleteAppCases',
    method: 'delete',
    data
  })
}

// @Tags AppCases
// @Summary 删除AppCases
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppCases"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appCases/deleteAppCases [delete]
export const deleteAppCasesByIds = (data) => {
  return service({
    url: '/appCases/deleteAppCasesByIds',
    method: 'delete',
    data
  })
}

// @Tags AppCases
// @Summary 更新AppCases
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppCases true "更新AppCases"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appCases/updateAppCases [put]
export const updateAppCases = (data) => {
  return service({
    url: '/appCases/updateAppCases',
    method: 'put',
    data
  })
}

// @Tags AppCases
// @Summary 用id查询AppCases
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppCases true "用id查询AppCases"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appCases/findAppCases [get]
export const findAppCases = (params) => {
  return service({
    url: '/appCases/findAppCases',
    method: 'get',
    params
  })
}

// @Tags AppCases
// @Summary 分页获取AppCases列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppCases列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCases/getAppCasesList [get]
export const getAppCasesList = (params) => {
  return service({
    url: '/appCases/getAppCasesList',
    method: 'get',
    params
  })
}
