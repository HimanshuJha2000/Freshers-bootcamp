//Routes/Routes.go
package Routes
import (
	"first-api/Controllers"
	"github.com/gin-gonic/gin"
)
//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/student-api")
	{
		grp1.GET("students", Controllers.GetStudents)
		grp1.POST("create", Controllers.CreateStudent)
		grp1.GET("getstudent/:id", Controllers.GetStudentByID)
		grp1.PUT("update/:id", Controllers.UpdateStudent)
		grp1.DELETE("delete/:id", Controllers.DeleteStudent)
	}
	return r
}