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

	transaction, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = createTable(transaction)
	if err != nil {
		transaction.Rollback()
		fmt.Println(err.Error())
		return
	}
	_, err = insertValues(transaction)
	if err != nil {
		transaction.Rollback()
		fmt.Println(err.Error())
		return
	}

	transaction.Commit()

	rows, _ := db.Query(
		"SELECT name, age, admin FROM Users;")

	for rows.Next() {
		rows.Scan(&name, &age, &admin)
	}
	defer rows.Close()


	fmt.Printf("Name: %s, Age: %d, Admin:%t", name, age, admin)
}
func insertValues(tx *sql.Tx) (sql.Result, error) {
	return tx.Exec(
		"INSERT INTO Users " +
			"(name, age, admin) VALUES " +
			"('Dino Omanovic', 27, true);")
}
func createTable(tx *sql.Tx) (sql.Result, error) {
	return tx.Exec(
		"CREATE TABLE Users (" +
			"id INT PRIMARY KEY AUTOINCREMENT, " +
			"name VARCHAR, " +
			"age int, " +
			"admin BOOLEAN" +
			");")
}
