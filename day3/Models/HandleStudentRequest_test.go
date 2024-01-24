package Models

import (
	"day3/Config"
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"

	"github.com/stretchr/testify/assert"
)

var err error

func init() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
}
func TestGetAllUser(t *testing.T) {

	//defer Config.DB.Close()
	students := []Student{}
	err := GetAllStudents(&students)
	assert.Nil(t, err, "We are not expecting error")
}

func TestGetStudentByIDWithValidID(t *testing.T) {
	student := Student{}
	var id string
	id = "2"
	err := GetStudentByID(&student, id)
	assert.Nil(t, err, "We are not expecting error")
	assert.EqualValues(t, "Sartha", student.FirstName, "we are expecting id as Sartha")
}

func TestGetStudentByIDWithInValidID(t *testing.T) {
	student := Student{}
	var id string
	id = "3"
	err := GetStudentByID(&student, id)
	assert.NotNil(t, err, "We are expecting error")
	assert.EqualValues(t, "record not found", err.Error(), "expecting record not found")
}
