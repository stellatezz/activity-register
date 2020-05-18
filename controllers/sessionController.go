package controllers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
)

func GetSession(c *gin.Context) bool {

	session := sessions.Default(c)
	loginuser := session.Get("loginuser")

	fmt.Println("loginuser:", loginuser)

	if loginuser != nil {
		return true
	} else {
		return false
	}

}

func AdminAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		v := session.Get("loginuser")
		fmt.Println("loginuser:", v)
		if v != 1 {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{})
			c.Abort()
		}
		c.Next()
	}
}

func StdAuthorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		v := session.Get("loginuser")
		fmt.Println("loginuser:", v)
		if v == nil {
			c.HTML(http.StatusUnauthorized, "login.html", gin.H{})
			c.Abort()
		}
		c.Next()
	}
}

