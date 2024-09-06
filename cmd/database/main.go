package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "polina"
	password = "123"
	dbname   = "mydatabase"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	name := "Kaluga"
	location := "64, 64"

	sqlStatement := `DELETE from cities WHERE name = $1`
	_, err = db.Exec(sqlStatement, name)
	if err != nil {
		panic(err)
	}

	sqlStatement = `INSERT INTO cities (name, location) VALUES ($1, $2)`
	_, err = db.Exec(sqlStatement, name, location)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM cities;")
	defer rows.Close()
	for rows.Next() {
		var city, location string
		if err := rows.Scan(&city, &location); err != nil {
			log.Fatal(err)
		}
		fmt.Println(city, location)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}
