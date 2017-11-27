package models

import (
    "database/sql"
    // "fmt"
)

type Marker struct {
    ID          int             `json:"id"`
    Type_id     int             `json:"type_id"`
    Project_id  int             `json:"project_id"`
    Lat         float32         `json:"lat"`
    Lng         float32         `json:"lng"`
    Start_date  sql.NullString  `json:"start_date"`
    End_date    sql.NullString  `json:"end_date"`
    Info        string          `json:"info"`
}

// TaskCollection is collection of Tasks
type MarkerCollection struct {
    Markers []Marker `json:"markers"`
}

func GetMarkers(db *sql.DB) MarkerCollection {

    sql := "SELECT * FROM markers"
    rows, err := db.Query(sql)
    // Exit if the SQL doesn't work for some reason
    if err != nil {
        panic(err)
    }
    // make sure to cleanup when the program exits
    defer rows.Close()

    result := MarkerCollection{}

    for rows.Next() {
        marker := Marker{}
        err2 := rows.Scan(
            &marker.ID,
            &marker.Type_id,
            &marker.Project_id,
            &marker.Lat,
            &marker.Lng,
            &marker.Start_date,
            &marker.End_date,
            &marker.Info)

        if err2 != nil {
            panic(err2)
        }

        result.Markers = append(result.Markers, marker)
    }
    return result
}

func PostMarker(db *sql.DB, marker Marker) (int64, error) {
    sql := `
    INSERT INTO markers(type_id, project_id, lat, lng, start_date, end_date, info) 
    VALUES(?, ?, ?, ?, ?, ?)
    `
    // Create a prepared SQL statement
    stmt, err := db.Prepare(sql)
    // Exit if we get an error
    checkErr(err)
    // Make sure to cleanup after the program exits
    defer stmt.Close()
    // Replace the '?' in our prepared statement with 'name'
    result, err2 := stmt.Exec(marker.Type_id, marker.Project_id, marker.Lat, marker.Lng, marker.Start_date, marker.End_date, marker.Info)
    // Exit if we get an error
    checkErr(err2);

    return result.LastInsertId()
}

func PutMarker(db *sql.DB, marker Marker, id int) (int64, error) {
    sql := `
    UPDATE markers SET type_id = ?, project_id = ?, lat = ?, lng = ?, start_date = ?, end_date = ?, info = ? where id = ?
    `
    // Create a prepared SQL statement
    stmt, err := db.Prepare(sql)
    // Exit if we get an error
    checkErr(err)
    // Make sure to cleanup after the program exits
    defer stmt.Close()
    // Replace the '?' in our prepared statement with 'name'
    result, err2 := stmt.Exec(marker.Type_id, marker.Project_id, marker.Lat, marker.Lng, marker.Start_date, marker.End_date, marker.Info, id)
    // Exit if we get an error
    checkErr(err2);

    return result.LastInsertId()
}

func DeleteMarker(db *sql.DB, id int) (int64, error) {
    sql := "DELETE FROM markers WHERE id = ?"

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
