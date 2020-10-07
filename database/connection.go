package database

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/lib/pq"
	
)

const (
	// Initialize connection constants.
	HOST     = "ec2-50-17-197-184.compute-1.amazonaws.com"
	DATABASE = "da6dlmoflefm1t"	
	USER     = "kncmwkmpxwqdfm"
	PASSWORD = "956ca3e0ce1b1f91fbb18925e549ee9598c811f0601f3f69f9525272e6080d7c"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetConnection() *sql.DB {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require", HOST, USER, PASSWORD, DATABASE)
	db, err := sql.Open("postgres", connectionString)


	if err != nil {
		log.Fatal(err)
		/*  */
	}

 	/* if err != nil {		
		panic(err)	
	} */
	
	return db
}
