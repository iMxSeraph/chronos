package interceptors

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
	"muxin.io/chronos/consts"
	"net/url"
)

func LoginInterceptor() gin.HandlerFunc {
	return func(context *gin.Context) {
		session := sessions.Default(context)
		username := session.Get("username")
		if username == nil {
			context.Redirect(http.StatusFound, "/login?redirect=" + url.PathEscape(context.Request.RequestURI))
			context.Abort()
		} else {
			session.Options(sessions.Options{MaxAge: consts.SessionMaxAge})
			session.Save()
			context.Next()
		}
	}
}
