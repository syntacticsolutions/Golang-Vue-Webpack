package models

import (
	"database/sql"
	// "fmt"
)

type Contractor struct {
	ID         int    `json:"id"`
	User_id    int    `json:"user_id"`
	Color      string `json:"color"`
	Contractor string `json:"contractor"`
	Phone      string `json:"contractor_phone"`
	Cell       string `json:"contractor_cell"`
	Email      string `json:"contractor_email"`
}

// TaskCollection is collection of Tasks
type ContractorCollection struct {
	Contractors []Contractor `json:"contractors"`
}

func GetContractors(db *sql.DB) ContractorCollection {

	sql := `SELECT contractors.id, user_id, color,
    CONCAT(first_name, ' ', last_name) as contractor,
    phone, cell, email
    FROM contractors
    LEFT JOIN users on user_id = users.id`
	rows, err := db.Query(sql)
	// Exit if the SQL doesn't work for some reason
	if err != nil {
		panic(err)
	}
	// make sure to cleanup when the program exits
	defer rows.Close()

	result := ContractorCollection{}

	for rows.Next() {
		contractor := Contractor{}
		err2 := rows.Scan(
			&contractor.ID,
			&contractor.User_id,
			&contractor.Color,
			&contractor.Contractor,
			&contractor.Phone,
			&contractor.Cell,
			&contractor.Email)

		if err2 != nil {
			panic(err2)
		}

		result.Contractors = append(result.Contractors, contractor)
	}
	// fmt.Printf("%+v", result)
	return result
}

func PostContractor(db *sql.DB, contractor Contractor) (int64, error) {
	sql := `
    INSERT INTO contractors(user_id, color) 
    VALUES(?, ?)
    `
	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	checkErr(err)
	// Make sure to cleanup after the program exits
	defer stmt.Close()
	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(contractor.User_id, contractor.Color)
	// Exit if we get an error
	checkErr(err2)

	return result.LastInsertId()
}

func PutContractor(db *sql.DB, contractor Contractor, id int) (int64, error) {
	sql := `
    UPDATE contractors SET user_id = ?, color = ? where id = ?
    `
	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	checkErr(err)
	// Make sure to cleanup after the program exits
	defer stmt.Close()
	// Replace the '?' in our prepared statement with 'name'
	result, err2 := stmt.Exec(contractor.User_id, contractor.Color, id)
	// Exit if we get an error
	checkErr(err2)

	return result.LastInsertId()
}

func DeleteContractor(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM contractors WHERE id = ?"

	// Create a prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	checkErr(err)

	// Replace the '?' in our prepared statement with 'id'
	result, err2 := stmt.Exec(id)
	// Exit if we get an error
	checkErr(err2)

	return result.RowsAffected()
}
