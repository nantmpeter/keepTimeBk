package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
)

type AppUserResponse struct {
	User app.AppUsers `json:"user"`
}

type LoginResponse struct {
	User      app.AppUserShow `json:"user"`
	Token     string          `json:"token"`
	ExpiresAt int64           `json:"expiresAt"`
}
