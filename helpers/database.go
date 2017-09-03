package helpers

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"log"
	"golang.org/x/crypto/bcrypt"
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

func (d *DatabaseHelper) Init() error {
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
			numbers		TEXT,
			bingo_at	INTEGER,
			updated_at	INTEGER,
			created_at	INTEGER
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
		return err
	}

	return nil
}

func (d *DatabaseHelper) Connect() *sql.DB {
	db, err := sql.Open("sqlite3", databasePath);
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func (d *DatabaseHelper) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (d *DatabaseHelper) VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}