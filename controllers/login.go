package controllers

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ldap.v2"
	"log"
	"muxin.io/chronos/models"
	"net/http"
	"muxin.io/chronos/consts"
)

const (
	ldapDn = "cn=%s,dc=benzhi,dc=io"
)

func DoLogin(ctx *gin.Context) {
	var login models.Login
	if err := ctx.BindJSON(&login); err == nil {
		result := checkLdap(login)
		if result {
			session := sessions.Default(ctx)
			session.Options(sessions.Options{MaxAge: consts.SessionMaxAge})
			session.Set("username", login.Username)
			session.Save()
			ctx.JSON(http.StatusOK, consts.Success)
		} else {
			ctx.JSON(http.StatusOK, consts.WrongInput)
		}
	} else {
		ctx.JSON(http.StatusBadRequest, consts.WrongInput)
	}
}

func LogoutEndpoint(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
}

func LoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

func checkLdap(login models.Login) bool {
	l, err := ldap.Dial("tcp", "ldap.benzhi.io:389")
	if err != nil {
		log.Println("error connet to ladp")
		return false
	}
	defer l.Close()

	err = l.Bind(fmt.Sprintf(ldapDn, login.Username), login.Password)
	return err == nil
}
