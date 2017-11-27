package models

import (
    "database/sql"
    "fmt"
)

type UserType struct {
    ID   int    `json:"id"`
    User_type string `json:"user_type"`
}

// TaskCollection is collection of Tasks
type UserTypeCollection struct {
    UserTypes []UserType `json:"user_types"`
}

func GetUserTypes(db *sql.DB) UserTypeCollection {

    sql := "SELECT * FROM user_types"
    rows, err := db.Query(sql)
    // Exit if the SQL doesn't work for some reason
    if err != nil {
        panic(err)
    }
    // make sure to cleanup when the program exits
    defer rows.Close()

    result := UserTypeCollection{}

    for rows.Next() {
        user_type := UserType{}
        err2 := rows.Scan(
            &user_type.ID,
            &user_type.User_type)

        if err2 != nil {
            panic(err2)
        }

        result.UserTypes = append(result.UserTypes, user_type)
    }
    fmt.Printf("%+v", result)
    return result
}

func PostUserType(db *sql.DB, user_type UserType) (int64, error) {
    sql := `
    INSERT INTO user_types(user_type) 
    VALUES(?)
    `
    // Create a prepared SQL statement
    stmt, err := db.Prepare(sql)
    // Exit if we get an error
    checkErr(err)
    // Make sure to cleanup after the program exits
    defer stmt.Close()
    // Replace the '?' in our prepared statement with 'name'
    result, err2 := stmt.Exec(user_type.User_type)
    // Exit if we get an error
    checkErr(err2);

    return result.LastInsertId()
}

func PutUserType(db *sql.DB, user_type UserType, id int) (int64, error) {
    sql := `
    UPDATE user_types SET user_type = ? where id = ?
    `
    // Create a prepared SQL statement
    stmt, err := db.Prepare(sql)
    // Exit if we get an error
    checkErr(err)
    // Make sure to cleanup after the program exits
    defer stmt.Close()
    // Replace the '?' in our prepared statement with 'name'
    result, err2 := stmt.Exec(user_type.User_type, id)
    // Exit if we get an error
    checkErr(err2);

    return result.LastInsertId()
}

func DeleteUserType(db *sql.DB, id int) (int64, error) {
    sql := "DELETE FROM user_types WHERE id = ?"

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

func checkErr(err error){
    if err != nil {
        panic(err)
    }
}