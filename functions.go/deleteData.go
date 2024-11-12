package forum

import (
	"database/sql"
	"fmt"
	"log"
)

func DeletRows(db *sql.DB, RowToDelete string) {
	stmt, err := db.Prepare("DELETE FROM " + RowToDelete)
	if err != nil {
		log.Fatal("Error when deleting posts rows -> ", err)
	}

	result, err := stmt.Exec()
	if err != nil {
		log.Fatal("Error when executing deleting posts rows -> ", err)
	}
	_, err = db.Exec("DELETE FROM sqlite_sequence WHERE name='" + RowToDelete + "'")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("VACUUM")
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Rows deleted : ", rowsAffected)
}
