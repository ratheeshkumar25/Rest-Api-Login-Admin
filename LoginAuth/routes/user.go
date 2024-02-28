package routes

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"main.go/database"
	"main.go/jwt"
	"main.go/model"
)

var Error string
var Fetch model.UserModel
var UpdateUser model.UserModel

const RoleUser = "user"

func Login(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RoleUser)
	if check == nil {
		c.HTML(200, "login.html", Error)
		Error = ""
	} else {
		c.Redirect(http.StatusSeeOther, "/home")
	}
}

func Postlogin(c *gin.Context) {
	Fetch = model.UserModel{}
	database.DB.First(&Fetch, "email=?", c.Request.FormValue("Email"))

	password := c.Request.FormValue("password")
	err := bcrypt.CompareHashAndPassword([]byte(Fetch.Password), []byte(password))

	if err != nil {
		Error = "Invalid Username or password"
		c.Redirect(http.StatusSeeOther, "/")
	} else {
		if Fetch.Status == "Blocked" {
			Error = "Blocked User"
			c.Redirect(http.StatusSeeOther, "/")
		} else {
			jwt.JwtToken(c, Fetch.Email, RoleUser)
			c.Redirect(http.StatusSeeOther, "/home")
		}
	}
}

func UserHome(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RoleUser)
	if check != nil {
		c.HTML(200, "user.html", gin.H{
			"Name":  Fetch.Name,
			"Email": Fetch.Email,
		})
	} else {
		c.Redirect(http.StatusSeeOther, "/")
	}
}

func Signup(c *gin.Context) {

	  // Setting Cache-Control header to ensure no caching
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	
	 // Retrieving session
	session := sessions.Default(c)
	check := session.Get(RoleUser)

	 // If user is logged in, redirect to the home page; otherwise, render the signup page
	if check != nil {
		c.Redirect(http.StatusSeeOther, "/home")
	} else {
		c.HTML(200, "Signup.html", Error)
		Error = ""
	}
}

func Postsignup(c *gin.Context) {
	//Hash password 
	 hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.Request.PostFormValue("password")), 10)
	 if err != nil {
		 // Handle hashing error
		 c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		 return
	 }
 
	 // Create a new user model
	 newUser := model.UserModel{
		 Name:     c.Request.PostFormValue("username"),
		 Email:    c.Request.PostFormValue("email"),
		 Password: string(hashedPassword),
		 Status:   "Block",
	 }
 
	 // Attempt to create the user in the database
	 if err := database.DB.Create(&newUser).Error; err != nil {
		 // Handle database error
		 Error = "User already exists"
		 c.Redirect(http.StatusSeeOther, "/signup")
		 return
	 }
 
	 // User successfully created
	 Error = "Successfully signed up"
	 c.Redirect(http.StatusSeeOther, "/")
}
func Logout(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	
	 // Retrieving session and deleting user's session
	session := sessions.Default(c)
	session.Delete(RoleUser)
	session.Save()

	 // Clearing Fetch variable and setting success message
	Fetch = model.UserModel{}
	Error = "Successfully Logged out"

	 // Clearing Fetch variable and setting success message
	c.Redirect(http.StatusSeeOther, "/")
}

