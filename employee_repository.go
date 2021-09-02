package main

import (
	"gorm.io/gorm"
)

// EmployeeRepo is the actual "object" that does the work.
type EmployeeRepo struct {
	db *gorm.DB
}

type IEmployeeRepo interface {
	FindAll() ([]Employee, error)
	FindById() (Employee, error)
	Save() error
	DeleteById() error
}

func NewEmployeeRepo(db *gorm.DB) *EmployeeRepo {
	return &EmployeeRepo{
		db: db,
	}
}

func (e EmployeeRepo) FindAll() ([]Employee, error) {
	var employees []Employee

	result := e.db.Find(&employees)
	if result.Error != nil {
		return []Employee{}, nil
	}

	return employees, nil
}

func (e EmployeeRepo) FindById(id int) (Employee, error) {
	var employee Employee

	result := e.db.First(&employee, id)
	if result.Error != nil {
		return Employee{}, result.Error
	}

	return employee, nil
}

func (e EmployeeRepo) Save(employee Employee) error {
	var gemployee Employee
	mode := "create"
	if employee.ID != 0 {
		mode = "update"
	}

	gemployee = Employee{}
	if mode == "update" {
		e.db.First(&gemployee, employee.ID)
	}

	gemployee.Name = employee.Name
	gemployee.Role = employee.Role

	var result *gorm.DB
	if mode == "create" {
		result = e.db.Create(&gemployee)
	} else if mode == "update" {
		result = e.db.Save(&gemployee)
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (e EmployeeRepo) DeleteById(id uint) error {
	result := e.db.Delete(&Employee{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
