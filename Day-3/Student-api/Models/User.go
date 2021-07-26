//Models/User.go
package Models
import (
	"first-api/Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
//GetAllUsers Fetch all user data
func GetStudents(user *[]Student) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}
//CreateUser ... Insert New data
func CreateStudent(user *Student) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
//GetUserByID ... Fetch only one user by Id
func GetStudentByID(user *Student, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}
//UpdateUser ... Update user
func UpdateStudent(user *Student, id string) (err error) {
	fmt.Println(user)
	Config.DB.Save(user)
	return nil
}
//DeleteUser ... Delete user
func DeleteStudent(user *Student, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(user)
	return nil
}