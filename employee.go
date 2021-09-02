package main

import (
	"fmt"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name string `json:"name"`
	Role string `json:"role"`
}

func (e Employee) String() string {
	return fmt.Sprintf("Employee{id=%d, name='%s', role='%s'}", e.ID, e.Name, e.Role)
}

func NewEmployee(name string, role string) *Employee {
	return &Employee{
		Name: name,
		Role: role,
	}
}
