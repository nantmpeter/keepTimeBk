package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AppCasesSearch struct{
    app.AppCases
    request.PageInfo
}
