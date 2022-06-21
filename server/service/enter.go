package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/app"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	AppServiceGroup     app.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
