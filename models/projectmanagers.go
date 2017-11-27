package models

import (
    "database/sql"
    // "fmt"
)

type ProjectManager struct {
    ID   int    `json:"id"`
    User_id int `json:"user_id"`
}

// TaskCollection is collection of Tasks
type ProjectManagerCollection struct {
    ProjectManagers []ProjectManager `json:"project_managers"`
}

func GetProjectManagers(db *sql.DB) ProjectManagerCollection {

    sql := "SELECT * FROM project_managers"
    rows, err := db.Query(sql)
    // Exit if the SQL doesn't work for some reason
    if err != nil {
        panic(err)
    }
    // make sure to cleanup when the program exits
    defer rows.Close()

    result := ProjectManagerCollection{}

    for rows.Next() {
        project_manager := ProjectManager{}
        err2 := rows.Scan(
            &project_manager.ID,
            &project_manager.User_id)

        if err2 != nil {
            panic(err2)
        }

        result.ProjectManagers = append(result.ProjectManagers, project_manager)
    }
    // fmt.Printf("%+v", result)
    return result
}

func PostProjectManager(db *sql.DB, project_manager ProjectManager) (int64, error) {
    sql := `
    INSERT INTO project_managers(user_id) 
    VALUES(?)
    `
    // Create a prepared SQL statement
    stmt, err := db.Prepare(sql)
    // Exit if we get an error
    checkErr(err)
    // Make sure to cleanup after the program exits
    defer stmt.Close()
    // Replace the '?' in our prepared statement with 'name'
    result, err2 := stmt.Exec(project_manager.User_id)
    // Exit if we get an error
    checkErr(err2);

    return result.LastInsertId()
}

func DeleteProjectManager(db *sql.DB, id int) (int64, error) {
    sql := "DELETE FROM project_managers WHERE id = ?"

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
