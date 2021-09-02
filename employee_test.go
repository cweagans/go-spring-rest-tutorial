package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployeeString(t *testing.T) {
	e := NewEmployee("John Doe", "Tester")
	assert.Equal(t, "Employee{id=0, name='John Doe', role='Tester'}", e.String())
}
