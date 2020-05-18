package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	//"strconv"
	//"fmt"
)

func IndexGet(c *gin.Context) {
	//islogin := GetSession(c)
	c.HTML(http.StatusOK, "login.html", gin.H{})
}