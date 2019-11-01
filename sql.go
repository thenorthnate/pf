package main

/*
import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

const (
	createUsers        = "CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)"
	createBanks        = "CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)"
	createAccounts     = "CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)"
	createTransactions = "CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)"

	insertUsers = "INSERT INTO people (firstname, lastname) VALUES (?, ?)"

	selectUsers = "SELECT id, firstname, lastname FROM people"
)

func insertData() {
	database, _ := sql.Open("sqlite3", "./vault.db")
	statement, _ := database.Prepare(createUsers)
	statement.Exec()

	statement, _ = database.Prepare(insertUsers)
	statement.Exec("Hello", "World")


	rows, _ := database.Query(selectUsers)
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}
}
*/
