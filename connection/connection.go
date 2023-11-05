package connection

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func GetConnection() *sql.DB {
	if db != nil {
		return db
	}

	var err error
	db, err = sql.Open("sqlite3", "data.sqlite")
	if err != nil {
		panic(err)
	}
	return db
}

func CreateDatabase() error {
    db := GetConnection()

	/*
    _, err := db.Exec("ATTACH DATABASE 'dbNotes' AS db_notes")
    if err != nil {
        return err
    }

    // Cerramos la conexión después de ejecutar la consulta.
    defer db.Close()

    return nil
	*/

	//defer db.Close()

	q := `CREATE TABLE IF NOT EXISTS notes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(64) NULL,
			description VARCHAR(200) NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL
	);`

	_, err := db.Exec(q)
	if err != nil {
		return err
	}

	fmt.Println("Base de datos y tabla creadas con exito!")
	return nil
}

func MakeMigrations() error {
	db := GetConnection()
	q := `CREATE TABLE IF NOT EXISTS notes (
	        id INTEGER PRIMARY KEY AUTOINCREMENT,
       		title VARCHAR(64) NULL,
       		description VARCHAR(200) NULL,
	        created_at TIMESTAMP DEFAULT DATETIME,
	        updated_at TIMESTAMP NOT NULL
	      );`

	_, err := db.Exec(q)
	if err != nil {
		return err
	}
	return nil
}