// models/student.go
package Models

import "github.com/jinzhu/gorm"

// Student represents the student model
type Student struct {
	gorm.Model
	FirstName string
	LastName  string
	DOB       string
	Address   string
	Subject   string
	Marks     int32
}
