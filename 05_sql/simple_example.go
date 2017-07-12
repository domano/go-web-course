package main

import (
	"database/sql"
	_ "github.com/proullon/ramsql/driver"

	"fmt"
)

var (
	name  string
	age   int
	admin bool
)

func main() {
	db, _ := sql.Open("ramsql", "MyFirstGoDB")
	defer db.Close()

	_, err := db.Exec(
		"CREATE TABLE Users (" +
			"id INT PRIMARY KEY AUTOINCREMENT, " +
			"name VARCHAR, " +
			"age int, " +
			"admin BOOLEAN" +
			");")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = db.Exec(
		"INSERT INTO Users " +
		"(name, age, admin) VALUES " +
		"('Dino Omanovic', 27, true);")
	if err != nil {
		fmt.Println(err.Error())
	}

	rows, err := db.Query(
		"SELECT name, age, admin FROM Users;")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&name, &age, &admin); err != nil {
			fmt.Println(err.Error())
		}
	}

	fmt.Printf("Name: %s, Age: %d, Admin:%t", name, age, admin)
}
