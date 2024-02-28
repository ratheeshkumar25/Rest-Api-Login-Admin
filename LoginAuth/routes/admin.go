package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/jwt"
	"main.go/model"
)

//Global variables are declared to store error messages, admin verification details, and user information.
var Err string
var Verify model.AdminModel
var UserTable []model.UserModel

//A constant RoleAdmin is declared with the value-admin
const RoleAdmin = "admin"

func Admin(c *gin.Context) {
	// Setting Cache-Control header to ensure no caching
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	
	// Retrieving session
	session := sessions.Default(c)
	check := session.Get(RoleAdmin)

	 // If admin is not logged in, render the admin page with any error message
	if check == nil {
		c.HTML(200, "Admin.html", Err)
	    Err = ""
	} else {
		c.Redirect(http.StatusSeeOther, "/valadmin")
	}
}

func PostAdmin(c *gin.Context) {
	Verify = model.AdminModel{}
	database.DB.First(&Verify, "email=?", c.Request.PostFormValue("AEmail"))
	if Verify.Password == c.Request.FormValue("Apassword") {
		jwt.JwtToken(c, Verify.Email, RoleAdmin)
		c.Redirect(http.StatusSeeOther, "/valadmin")
	} else {
		Err = "invalid email or password"
		c.Redirect(http.StatusSeeOther, "/admin")
	}
}

func Valadmin(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RoleAdmin)
	if check != nil {
		database.DB.Find(&UserTable)
	c.HTML(200, "Adminhome.html", gin.H{
		"Name":  Verify.Name,
		"Users": UserTable,
		"Error": Err,
	})
	Error = ""
	Err = ""
	} else {
		c.Redirect(http.StatusSeeOther, "/admin")
	}
}

func Adminlogout(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	session.Delete(RoleAdmin)
	session.Save()
	Err = "successfully logged out"
	c.Redirect(http.StatusSeeOther, "/admin")
}

func Delete(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RoleAdmin)
	if check != nil {
	user := c.Param("ID")
	database.DB.First(&UpdateUser, "ID=?", user)
	database.DB.Delete(&UpdateUser)
	UpdateUser = model.UserModel{}
	c.Redirect(http.StatusSeeOther, "/valadmin")
	} else {
		c.Redirect(http.StatusSeeOther, "/admin")
	}
}

func Update(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RoleAdmin)
	if check != nil {
		user := c.Param("ID")
	c.HTML(http.StatusSeeOther, "update.html", user)
	} else {
		c.Redirect(http.StatusSeeOther, "/admin")
	}
}

func Updateuser(c *gin.Context) {
	user := c.Param("ID")
	database.DB.First(&UpdateUser, "ID=?", user)
	UpdateUser.Name = c.Request.FormValue("name")
	UpdateUser.Email = c.Request.FormValue("email")
	database.DB.Save(&UpdateUser)
	Err = "User details updated successfully"
	UpdateUser = model.UserModel{}
	c.Redirect(http.StatusSeeOther, "/valadmin")
}

func Block(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RoleAdmin)
	  if check != nil {
	user := c.Param("ID")
	database.DB.First(&UpdateUser, "ID=?", user)
	if UpdateUser.Status == "Active" {
		UpdateUser.Status = "Blocked"
		database.DB.Save(&UpdateUser)
		UpdateUser = model.UserModel{}
		c.Redirect(http.StatusSeeOther, "/valadmin")
	} else {
		UpdateUser.Status = "Active"
		database.DB.Save(&UpdateUser)
		UpdateUser = model.UserModel{}
		c.Redirect(http.StatusSeeOther, "/valadmin")
	}
 } else {
	c.Redirect(http.StatusSeeOther, "/admin")
 }
}  
