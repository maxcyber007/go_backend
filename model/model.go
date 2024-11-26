package model

import (
	"gorm.io/gorm"
)

type Tbl_employee struct {
	gorm.Model
	Emp_id int
	Emp_firstname string
	Emp_lastname string
}