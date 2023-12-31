package forum

import (
	"database/sql"
	"log"
	_"github.com/mattn/go-sqlite3"
)


func CreateTables() {
	Database, err := sql.Open("sqlite3", "./forum.db") // The Database must be opened and assigned to a variable
	if err != nil {
		log.Fatal(err)
	}
	_, errpragma := Database.Exec("PRAGMA foreign_keys = ON") // not sure if we need this but foreign keys are disabled by default
	if errpragma != nil {
		log.Fatal(errpragma)
	}
	usertable, err2 := Database.Prepare("CREATE TABLE IF NOT EXISTS Users (id INTEGER PRIMARY KEY, username TEXT NOT NULL, Password TEXT NOT NULL, email TEXT)")
	if err2 != nil {      // table for users is created if it does not exist in line 24 
		log.Fatal(err2)
	}
	usertable.Exec() // Exec executes the query which was Prepared in line 20
	postsTable, err4 := Database.Prepare("CREATE TABLE IF NOT EXISTS Posts (id INTEGER PRIMARY KEY, Title TEXT, body TEXT, user_id INTEGER)")
	if err4 != nil { // table for posts is created if it doesnot exist
		log.Fatal(err4)
	}
	postsTable.Exec() // Exec executes query on line 25
	defer Database.Close()
}