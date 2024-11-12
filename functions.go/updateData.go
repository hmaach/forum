package forum

import (
	"database/sql"
	"fmt"
)

func UpdateField(db *sql.DB, table string, id int, field, data string) error {
	query := "UPDATE " + table + " SET " + field + " = " + field + " || ? WHERE id = ?"
	_, err := db.Exec(query, ", "+data, id)
	if err == nil {
		fmt.Println(field, " Updated success!")
	}
	return err
}

func EmptyField(db *sql.DB, table string, id int, field string) error {
	query := "UPDATE " + table + " SET " + field + " = ? WHERE id = ?"
	_, err := db.Exec(query, "", id)
	if err == nil {
		fmt.Println(field, " Updated success!")
	}
	return err
}
