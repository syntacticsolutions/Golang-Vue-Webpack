package models

import (
    "database/sql"
    // "fmt"
)

type User struct {
    ID   int    `json:"id"`
    User_type_id int `json:"user_type_id"`
    First_name string `json:"first_name"`
    Last_name string `json:"last_name"`
    Email string `json:"email"`
    Password string `json:"password"`
    Phone string `json:"phone"`
    Cell string `json:"cell"`
}

// TaskCollection is collection of Tasks
type UserCollection struct {
    Users []User `json:"users"`
}

func GetUsers(db *sql.DB) UserCollection {

    sql := "SELECT id, user_type_id, first_name, last_name, email, phone, cell FROM users"
    rows, err := db.Query(sql)
    // Exit if the SQL doesn't work for some reason
    if err != nil {
        panic(err)
    }
    // make sure to cleanup when the program exits
    defer rows.Close()

    result := UserCollection{}

    for rows.Next() {
        user := User{}
        err2 := rows.Scan(
            &user.ID,
            &user.User_type_id,
            &user.First_name,
            &user.Last_name,
            &user.Email,
            &user.Phone,
            &user.Cell)

        // fmt.Printf(user.first_name)
        // fmt.Printf("%+v", user)
        // Exit if we get an error
        if err2 != nil {
            panic(err2)
        }

        result.Users = append(result.Users, user)
    }
    // fmt.Printf("%+v", result)
    return result
}



func PostUser(db *sql.DB, user User) (int64, error) {
    sql := `
    INSERT INTO users(user_type_id, first_name, last_name, email, password, phone, cell) 
    VALUES(?, ?, ?, ?, ?, ?, ?)
    `
    // Create a prepared SQL statement
    stmt, err := db.Prepare(sql)
    // Exit if we get an error
    checkErr(err)
    // Make sure to cleanup after the program exits
    defer stmt.Close()
    // Replace the '?' in our prepared statement with 'name'
    result, err2 := stmt.Exec(user.User_type_id, user.First_name, user.Last_name, user.Email, user.Password, user.Phone, user.Cell)
    // Exit if we get an error
    checkErr(err2);

    return result.LastInsertId()
}

func PutUser(db *sql.DB, user User, id int) (int64, error) {
    sql := `
    UPDATE users SET user_type_id = ?, first_name = ?, last_name = ?, email = ?, password = ?, phone = ?, cell = ? where id = ?
    `
    // Create a prepared SQL statement
    stmt, err := db.Prepare(sql)
    // Exit if we get an error
    checkErr(err)
    // Make sure to cleanup after the program exits
    defer stmt.Close()
    // Replace the '?' in our prepared statement with 'name'
    result, err2 := stmt.Exec(user.User_type_id, user.First_name, user.Last_name, user.Email, user.Password, user.Phone, user.Cell, id)
    // Exit if we get an error
    checkErr(err2);

    return result.LastInsertId()
}

func DeleteUser(db *sql.DB, id int) (int64, error) {
    sql := "DELETE FROM users WHERE id = ?"

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
