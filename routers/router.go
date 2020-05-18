package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/ivancc666/tkp-register-go/controllers"
    "github.com/gin-gonic/contrib/sessions"
)

func InitRouter() *gin.Engine {

    router := gin.Default()
    router.LoadHTMLGlob("views/*")
    
    store := sessions.NewCookieStore([]byte("loginuser"))
    router.Use(sessions.Sessions("mysession", store))

    {
    	router.GET("/", controllers.IndexGet)

    	router.GET("/login", controllers.LoginGet)
    	router.POST("/login/process", controllers.LoginPost)

    	router.GET("/exit", controllers.ExitGet)

    	v1 := router.Group("/administrator")
		v1.Use(controllers.AdminAuthorize())
		{	
			v1.GET("/home", controllers.AdminIndex)
			v1.GET("/exit", controllers.ExitGet)
			v1.GET("/create", controllers.AddClassGet)
    		v1.POST("/create/process", controllers.AddClassPost)
    		v1.GET("/showdetail", controllers.ShowClassDetail)
		}

		v2 := router.Group("/student")
		v2.Use(controllers.StdAuthorize())
		{
			v2.GET("/exit", controllers.ExitGet)
			v2.GET("/home", controllers.StudentIndex)
			v2.GET("/enrollment", controllers.StudentEroll)
			v2.POST("/enrollment/process", controllers.StudentErollProcess)
		}
		
    }
	
    return router
}