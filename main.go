package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	db, err := GetDatabaseConnection()
	if err != nil {
		panic("cannot open db")
	}

	conn, err := db.DB()
	if err != nil {
		panic("cannot get db connection handle")
	}
	defer conn.Close()

	repo := NewEmployeeRepo(db)

	router, err := GetHttpRouter(repo)
	if err != nil {
		panic("cannot create http router")
	}

	s := &http.Server{
		Addr:           ":9090",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println(s.ListenAndServe())
}
