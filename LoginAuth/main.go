package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/helper"
	"main.go/routes"
)

func init() {
	helper.LoadEnv()
	database.DBconnect()
}

func main() {
	//Default Engine 
	r := gin.Default()

	//Loading HTML files
	r.LoadHTMLGlob("templates/*")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	//user
	r.GET("/", routes.Login)
	r.POST("/", routes.Postlogin)
	r.GET("/home", routes.UserHome)
	r.GET("/signup", routes.Signup)
	r.POST("/signup", routes.Postsignup)
	r.GET("/logout", routes.Logout)

	//admin
	r.GET("/admin", routes.Admin)
	r.POST("/admin", routes.PostAdmin)
	r.GET("/valadmin", routes.Valadmin)
	r.GET("/delete/:ID", routes.Delete)
	r.GET("/update/:ID", routes.Update)
	r.POST("/update/:ID", routes.Updateuser)
	r.GET("/block/:ID", routes.Block)
	r.GET("/adminlogout", routes.Adminlogout)
	

	r.Run(":3000")
}
