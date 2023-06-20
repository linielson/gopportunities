package schemas

import "gorm.io/gorm"

type Opening struct {
	gorm.Modules
	Roles    string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}
