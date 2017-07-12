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

	createTable(db)

	statement, _ := db.Prepare(
		"INSERT INTO Users " +
		"(name, age, admin) VALUES " +
		"(?, ?, ?);")

	statement.Exec("Dino Omanovic", 27, true)
	rows, _:= db.Query(
		"SELECT name, age, admin FROM Users;")

	for rows.Next() {
		rows.Scan(&name, &age, &admin)
	}
	defer rows.Close()


	fmt.Printf("Name: %s, Age: %d, Admin:%t", name, age, admin)
}

func createTable(db *sql.DB) (sql.Result, error) {
	return db.Exec(
		"CREATE TABLE Users (" +
			"id INT PRIMARY KEY AUTOINCREMENT, " +
			"name VARCHAR, " +
			"age int, " +
			"admin BOOLEAN" +
			");")
}
