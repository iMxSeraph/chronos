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

func LoginEndpoint(context *gin.Context) {
	var login models.Login
	if err := context.BindJSON(&login); err == nil {
		result := checkLdap(login)
		if result {
			session := sessions.Default(context)
			session.Options(sessions.Options{MaxAge: consts.SessionMaxAge})
			session.Set("username", login.Username)
			session.Save()
		}
		context.JSON(http.StatusOK, gin.H{"result": result})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func LogoutEndpoint(context *gin.Context) {
	session := sessions.Default(context)
	session.Clear()
	session.Save()
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
