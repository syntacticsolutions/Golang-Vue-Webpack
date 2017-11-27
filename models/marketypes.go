package models

import (
    "database/sql"
    // "fmt"
)

type MarkerType struct {
    ID   int    `json:"id"`
    Marker_type string `json:"marker_type"`
    Svg string `json:"svg"`
}

// TaskCollection is collection of Tasks
type MarkerTypeCollection struct {
    MarkerTypes []MarkerType `json:"marker_types"`
}

func GetMarkerTypes(db *sql.DB) MarkerTypeCollection {

    sql := "SELECT * FROM marker_types"
    rows, err := db.Query(sql)
    // Exit if the SQL doesn't work for some reason
    if err != nil {
        panic(err)
    }
    // make sure to cleanup when the program exits
    defer rows.Close()

    result := MarkerTypeCollection{}

    for rows.Next() {
        marker_type := MarkerType{}
        err2 := rows.Scan(
            &marker_type.ID,
            &marker_type.Marker_type,
            &marker_type.Svg)

        if err2 != nil {
            panic(err2)
        }

        result.MarkerTypes = append(result.MarkerTypes, marker_type)
    }
    // fmt.Printf("%+v", result)
    return result
}

func PostMarkerType(db *sql.DB, marker_type MarkerType) (int64, error) {
    sql := `
    INSERT INTO marker_types(marker_type, svg) 
    VALUES(?, ?)
    `
    // Create a prepared SQL statement
    stmt, err := db.Prepare(sql)
    // Exit if we get an error
    checkErr(err)
    // Make sure to cleanup after the program exits
    defer stmt.Close()
    // Replace the '?' in our prepared statement with 'name'
    result, err2 := stmt.Exec(marker_type.Marker_type, marker_type.Svg)
    // Exit if we get an error
    checkErr(err2);

    return result.LastInsertId()
}

func PutMarkerType(db *sql.DB, marker_type MarkerType, id int) (int64, error) {
    sql := `
    UPDATE marker_types SET marker_type = ?, svg = ? where id = ?
    `
    // Create a prepared SQL statement
    stmt, err := db.Prepare(sql)
    // Exit if we get an error
    checkErr(err)
    // Make sure to cleanup after the program exits
    defer stmt.Close()
    // Replace the '?' in our prepared statement with 'name'
    result, err2 := stmt.Exec(marker_type.Marker_type, marker_type.Svg, id)
    // Exit if we get an error
    checkErr(err2);

    return result.LastInsertId()
}

func DeleteMarkerType(db *sql.DB, id int) (int64, error) {
    sql := "DELETE FROM marker_types WHERE id = ?"

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