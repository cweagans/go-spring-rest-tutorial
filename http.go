package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetHttpRouter(repo *EmployeeRepo) (*gin.Engine, error) {
	router := gin.Default()

	// Add a custom middleware to make sure that the repo is avilable in the
	// request context.
	router.Use(provideEmployeeRepo(repo))

	router.GET("/_health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	router.GET("/some-invalid-endpoint", notImplemented)
	router.GET("/employees", getAllEmployees)
	router.POST("/employees", postEmployee)
	router.GET("/employees/:id", getEmployee) // this is still wrong. revisit tomorrow.
	router.PUT("/employees/:id", putEmployee)
	router.DELETE("/employees/:id", deleteEmployee)

	return router, nil
}

func provideEmployeeRepo(repo *EmployeeRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("employee_repo", repo)
	}
}

func notImplemented(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{})
}

func getAllEmployees(c *gin.Context) {
	repo := c.MustGet("employee_repo").(*EmployeeRepo)
	employees, err := repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{"employees": employees})
	}
}

func getEmployee(c *gin.Context) {
	id := getIdAsUintFromRequest(c)

	repo := c.MustGet("employee_repo").(*EmployeeRepo)

	employee, err := repo.FindById(int(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	} else {
		c.JSON(http.StatusOK, employee)
	}
}

func postEmployee(c *gin.Context) {
	e := getEmployeeFromRequest(c)
	saveEmployee(c, e)
}

func putEmployee(c *gin.Context) {
	e := getEmployeeFromRequest(c)
	id := getIdAsUintFromRequest(c)

	e.ID = id

	saveEmployee(c, e)
}

func deleteEmployee(c *gin.Context) {
	id := getIdAsUintFromRequest(c)
	repo := c.MustGet("employee_repo").(*EmployeeRepo)

	err := repo.DeleteById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}

func getEmployeeFromRequest(c *gin.Context) Employee {
	var e Employee

	ctype := c.ContentType()
	switch ctype {
	case "application/x-www-form-urlencoded":
		e.Name = c.PostForm("name")
		e.Role = c.PostForm("role")
	case "application/json":
		c.BindJSON(&e)
	default:
		c.AbortWithStatusJSON(http.StatusNotAcceptable, errors.New("unacceptable content type"))
		return Employee{}
	}

	return e
}

func saveEmployee(c *gin.Context, e Employee) {
	repo := c.MustGet("employee_repo").(*EmployeeRepo)

	err := repo.Save(e)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}

func getIdAsUintFromRequest(c *gin.Context) uint {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return 0
	}

	return uint(id)
}
