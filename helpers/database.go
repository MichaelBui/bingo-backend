package helpers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"log"
)

type DatabaseHelper struct {

}

var (
	database *DatabaseHelper;
	databasePath string = "./db.sqlite";
)

func Database() *DatabaseHelper {
	if database == nil {
		database = &DatabaseHelper{}
	}
	return database;
}

func (d *DatabaseHelper) Init() bool {
	err := os.Remove(databasePath);
	if err != nil {
		log.Fatal(err);
	}

	db := d.Connect();
	defer db.Close();

	sqlQuery := `
		CREATE TABLE users (
			id 			INTEGER NOT NULL PRIMARY KEY,
			display 	TEXT,
			email 		TEXT,
			password	TEXT,
			score		INTEGER,
			number		TEXT,
			timestamp	INTEGER
		);
		DELETE FROM users;
		CREATE TABLE numbers (
			id 			INTEGER NOT NULL PRIMARY KEY,
			value 		INTEGER,
			timestamp	INTEGER
		);
		DELETE FROM numbers;
	`
	_, err = db.Exec(sqlQuery)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlQuery)
		return false
	}

	return true
}

func (d *DatabaseHelper) Connect() *sql.DB {
	db, err := sql.Open("sqlite3", databasePath);
	if err != nil {
		log.Fatal(err)
	}

	return db
}