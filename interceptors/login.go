package interceptors

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
	"muxin.io/chronos/models"
	"muxin.io/chronos/consts"
)

func LoginInterceptor() gin.HandlerFunc {
	return func(context *gin.Context) {
		session := sessions.Default(context)
		username := session.Get("username")
		if username == nil {
			context.JSON(http.StatusUnauthorized, models.Error{ErrCode: 1, ErrMsg: "请登录"})
		} else {
			session.Options(sessions.Options{MaxAge: consts.SessionMaxAge})
			session.Save()
			context.Next()
		}
	}
}
