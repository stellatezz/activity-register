package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/ivancc666/tkp-register-go/models"
    "net/http"
    "strconv"
    "time"
    "fmt"
)

func StudentIndex(c *gin.Context) {

	classes := models.ShowAllClass()
	c.HTML(http.StatusOK, "studentView.html", gin.H{"classes":classes})
}

func StudentEroll(c *gin.Context) {

	classes := models.ShowAllClass()
	c.HTML(http.StatusOK, "studentEroll.html", gin.H{"classes":classes})
}

func StudentErollProcess(c *gin.Context) {
	
	session := sessions.Default(c)
	sidTmp := session.Get("loginuser")
	sid := sidTmp.(int)

	cidTmp := c.PostForm("classid")
	cid, _ := strconv.Atoi(cidTmp)

	t := time.Now()
	enrolltime := t.Format("2006-01-02 15:04:05")

	fmt.Println(enrolltime)

	enroll := models.Enrollment{cid, sid, enrolltime, 1}

	_, err := models.AddEnroll(enroll)
	if err != nil {

	}
	
	c.Redirect(http.StatusMovedPermanently,"/student/enrollment")
}



