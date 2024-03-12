package routes

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	c.Redirect(304, "/admin")
}

func Delete(c *gin.Context) {
	c.Header("Cache-Control", "no-cache,no-store,must-revalidate")
	session := sessions.Default(c)
	check := session.Get(RoleAdmin)
	if check != nil {
	user := c.Param("ID")
	database.DB.First(&UpdateUser, "ID=?", user)
	database.DB.Delete(&UpdateUser)
    Err = "User details deleted successfully"
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
		Err = "User access blocked"
		UpdateUser = model.UserModel{}
  

        UpdateUser = model.UserModel{}
		c.Redirect(http.StatusSeeOther, "/valadmin")
	} else {
		UpdateUser.Status = "Active"
		database.DB.Save(&UpdateUser)
		Err = "User access Activated"
		UpdateUser = model.UserModel{}
		c.Redirect(http.StatusSeeOther, "/valadmin")
	}
 } else {
	c.Redirect(http.StatusSeeOther, "/admin")
 }

 
}  

func Search(c *gin.Context) {
    // search query from the request parameters
    query := c.Query("query")

    // Perform the search operation 
    var searchResults []model.UserModel
    for _, user := range UserTable {
        // Check if the query matches any user's name or email
        if strings.Contains(strings.ToLower(user.Name), strings.ToLower(query)) ||
           strings.Contains(strings.ToLower(user.Email), strings.ToLower(query)) {
            searchResults = append(searchResults, user)
        }
    }

    // Render the search results template with the matching users
    c.HTML(http.StatusOK, "SearchResults.html", gin.H{
        "Results": searchResults,
        "Query":   query,
    })
}

func AddUser(c *gin.Context) {
    // Get user details from the form submission
    name := c.PostForm("name")
    email := c.PostForm("email")
	password := c.PostForm("password")

	    // Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			// Handle error (e.g., log it, return an error response)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

      // Create new user model with hashed password
    newUser := model.UserModel{
        Name:  name,
        Email: email,
		Password: string(hashedPassword),
      
    }
	
	Err = "User details added"
    // Add the new user to the database
    database.DB.Create(&newUser)
	

    // Redirect back to the admin page after adding the user
    c.Redirect(http.StatusSeeOther, "/valadmin")
}

