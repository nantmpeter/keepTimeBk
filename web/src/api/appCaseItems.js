import service from '@/utils/request'

// @Tags AppCaseItems
// @Summary 创建AppCaseItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppCaseItems true "创建AppCaseItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCaseItems/createAppCaseItems [post]
export const createAppCaseItems = (data) => {
  return service({
    url: '/appCaseItems/createAppCaseItems',
    method: 'post',
    data
  })
}

// @Tags AppCaseItems
// @Summary 删除AppCaseItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppCaseItems true "删除AppCaseItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appCaseItems/deleteAppCaseItems [delete]
export const deleteAppCaseItems = (data) => {
  return service({
    url: '/appCaseItems/deleteAppCaseItems',
    method: 'delete',
    data
  })
}

// @Tags AppCaseItems
// @Summary 删除AppCaseItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppCaseItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appCaseItems/deleteAppCaseItems [delete]
export const deleteAppCaseItemsByIds = (data) => {
  return service({
    url: '/appCaseItems/deleteAppCaseItemsByIds',
    method: 'delete',
    data
  })
}

// @Tags AppCaseItems
// @Summary 更新AppCaseItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AppCaseItems true "更新AppCaseItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appCaseItems/updateAppCaseItems [put]
export const updateAppCaseItems = (data) => {
  return service({
    url: '/appCaseItems/updateAppCaseItems',
    method: 'put',
    data
  })
}

// @Tags AppCaseItems
// @Summary 用id查询AppCaseItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.AppCaseItems true "用id查询AppCaseItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appCaseItems/findAppCaseItems [get]
export const findAppCaseItems = (params) => {
  return service({
    url: '/appCaseItems/findAppCaseItems',
    method: 'get',
    params
  })
}

// @Tags AppCaseItems
// @Summary 分页获取AppCaseItems列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取AppCaseItems列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appCaseItems/getAppCaseItemsList [get]
export const getAppCaseItemsList = (params) => {
  return service({
    url: '/appCaseItems/getAppCaseItemsList',
    method: 'get',
    params
  })
}
