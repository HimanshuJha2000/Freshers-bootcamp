//Controllers/User.go
package Controllers
import (
	"first-api/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
//GetUsers ... Get all Students
func GetStudents(c *gin.Context) {
	var user []Models.Student
	err := Models.GetStudents(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//CreateUser ... Create User
func CreateStudent(c *gin.Context) {
	var user Models.Student
	c.BindJSON(&user)
	err := Models.CreateStudent(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//GetUserByID ... Get the user by id
func GetStudentByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.Student
	err := Models.GetStudentByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//UpdateUser ... Update the user information
func UpdateStudent(c *gin.Context) {
	var user Models.Student
	id := c.Params.ByName("id")
	err := Models.GetStudentByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateStudent(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
//DeleteUser ... Delete the user
func DeleteStudent(c *gin.Context) {
	var user Models.Student
	id := c.Params.ByName("id")
	err := Models.DeleteStudent(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}