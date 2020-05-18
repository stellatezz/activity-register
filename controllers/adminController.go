package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ivancc666/tkp-register-go/models"
    "net/http"
    "strconv"
)

func AdminIndex(c *gin.Context) {
	classes := models.ShowAllClass()
	c.HTML(http.StatusOK, "admin", gin.H{"classes":classes})
}

func AddClassGet(c *gin.Context) {
	cats := models.ShowAllCategory()
	c.HTML(http.StatusOK, "create", gin.H{"cats":cats})
}

func AddClassPost(c *gin.Context) {

	name := c.PostForm("name")
	cat := c.PostForm("category")
	date := c.PostForm("date")
	time := c.PostForm("time")
	venue := c.PostForm("venue")
	tutor := c.PostForm("tutor")
	quo := c.PostForm("quota")
	fee := c.PostForm("fee")
	description := c.PostForm("description")

	quota, _ := strconv.Atoi(quo)
	category, _ := strconv.Atoi(cat)

	class := models.Class{0, name, date, time, venue, tutor, quota, fee, description, category}

	_, err := models.AddClass(class)
	if err != nil {

	}


	c.Redirect(http.StatusMovedPermanently,"/administrator/create")
}

func ShowClassDetail(c *gin.Context) {
	id := c.Query("classid")
	class := models.QueryClassWithID(id)
	enrolls := models.ShowEnrollmentWithCid(id)
	c.HTML(http.StatusOK, "detail", gin.H{"Name":class.Name, "Date":class.Date, "Time":class.Time, "Venue":class.Venue, "Tutor":class.Tutor, "Fee":class.Fee, "Quota":class.Quota, "Description":class.Description, "Enrolls":enrolls})
}