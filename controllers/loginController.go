package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
    "github.com/ivancc666/tkp-register-go/models"
    "github.com/gin-gonic/contrib/sessions"

)

func LoginGet(c *gin.Context) {
    
    c.HTML(http.StatusOK, "login.html", gin.H{"title": "login"})
}

func LoginPost(c *gin.Context) {
    
    email := c.PostForm("username")
    password := c.PostForm("password")
    fmt.Println("username:", email, ",password:", password)
    id := models.QueryUserWithEmail(email)
    fmt.Println("id:", id)
    if id > 0 {
    	session := sessions.Default(c)
    	session.Set("loginuser", id)
    	session.Save()
    	c.Redirect(http.StatusMovedPermanently, "/administrator/home")
    } else {
        c.JSON(http.StatusOK, gin.H{"code": 0, "message": "fail"})
    }
}

func ExitGet(c *gin.Context)  {

    session := sessions.Default(c)
    session.Delete("loginuser")
    session.Save()

    fmt.Println("delete session...",session.Get("loginuser"))
    c.Redirect(http.StatusMovedPermanently,"/")
}