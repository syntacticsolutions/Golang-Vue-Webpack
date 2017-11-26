package models

import (
    "database/sql"
    "fmt"
)

type Project struct {
    ID   int    `json:"id"`
    Primary_marker_id int `json:"primary_marker_id"`
    Contractor_id int `json:"contractor_id"`
    Pm_id int `json:"pm_id"`
    Title string `json:"title"`
    Start_date sql.NullString `json:"start_date"`
    End_date sql.NullString `json:"end_date"`
}

// TaskCollection is collection of Tasks
type ProjectCollection struct {
    Projects []Project `json:"projects"`
}

func GetProjects(db *sql.DB) ProjectCollection {

    sql := "SELECT * FROM projects"
    rows, err := db.Query(sql)
    // Exit if the SQL doesn't work for some reason
    if err != nil {
        panic(err)
    }
    // make sure to cleanup when the program exits
    defer rows.Close()

    result := ProjectCollection{}

    for rows.Next() {
        project := Project{}
        err2 := rows.Scan(
            &project.ID,
            &project.Primary_marker_id,
            &project.Contractor_id,
            &project.Pm_id,
            &project.Title,
            &project.Start_date,
            &project.End_date)

        // fmt.Printf(user.first_name)
        // fmt.Printf("%+v", user)
        // Exit if we get an error
        if err2 != nil {
            panic(err2)
        }

        result.Projects = append(result.Projects, project)
    }
    fmt.Printf("%+v", result)
    return result
}



// func PostProject(db *sql.DB, user User) (int64, error) {
//     sql := `
//     INSERT INTO users(user_type_id, first_name, last_name, email, password, phone, cell) 
//     VALUES(?, ?, ?, ?, ?, ?, ?)
//     `
//     // Create a prepared SQL statement
//     stmt, err := db.Prepare(sql)
//     // Exit if we get an error
//     checkErr(err)
//     // Make sure to cleanup after the program exits
//     defer stmt.Close()
//     // Replace the '?' in our prepared statement with 'name'
//     result, err2 := stmt.Exec(user.User_type_id, user.First_name, user.Last_name, user.Email, user.Password, user.Phone, user.Cell)
//     // Exit if we get an error
//     checkErr(err2);

//     return result.LastInsertId()
// }

// func DeleteProject(db *sql.DB, id int) (int64, error) {
//     sql := "DELETE FROM users WHERE id = ?"

//     // Create a prepared SQL statement
//     stmt, err := db.Prepare(sql)
//     // Exit if we get an error
//     checkErr(err)

//     // Replace the '?' in our prepared statement with 'id'
//     result, err2 := stmt.Exec(id)
//     // Exit if we get an error
//     checkErr(err2)

//     return result.RowsAffected()
// }

// func checkErr(err error){
//     if err != nil {
//         panic(err)
//     }
// }