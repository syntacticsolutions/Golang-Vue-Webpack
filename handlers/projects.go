package handlers

import (
	"database/sql"
    "net/http"
    // "strconv"
    "go-echo-vue/models"
    "github.com/labstack/echo"
    // "encoding/json"
    // "time"
)

// type User struct {
//     id   int    `json:"id"`
//     user_type_id int `json:"user_type_id"`
//     first_name string `json:"first_name"`
//     last_name string `json:"last_name"`
//     email string `json:"email"`
//     password string `json:"email"`
//     phone string `json:"email"`
//     cell string `json:"email"`
// }

// var (
// 	users = map[int]*User{}
// 	seq   = 1
// )

func GetProjects(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Fetch tasks using our new model
        return c.JSON(http.StatusOK, models.GetProjects(db))
    }
}

// PutTask endpoint
func PostProject(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Instantiate a new task

  //       u := &User{
		// 	id: seq,
		// }

		// if err := c.Bind(u); err != nil {
		// 	return err
		// }

		// users[u.id] = u
		// seq++

		return c.JSON(http.StatusCreated, H{"created": "1", })

        // checkErr(err)
        // // Map imcoming JSON body to the new Task
        // // c.Bind(&id, &user_type_id, &first_name, &last_name, &email, &password, &phone, &cell)
        // // Add a task using our new model
        // id, err := models.PostUser(db, user)
        // // Return a JSON response if successful
        // if err == nil {
        //     return c.JSON(http.StatusCreated, H{
        //         "created": id,
        //     })
        // // Handle any errors
        // } else {
        //     return err
        // }
    }
}

// DeleteTask endpoint
// func DeleteUser(db *sql.DB) echo.HandlerFunc {
//     return func(c echo.Context) error {
//         id, _ := strconv.Atoi(c.Param("id"))
//         // Use our new model to delete a task
//         _, err := models.DeleteUser(db, id)
//         // Return a JSON response on success
//         if err == nil {
//             return c.JSON(http.StatusOK, H{
//                 "deleted": id,
//             })
//         // Handle errors
//         } else {
//             return err
//         }
//     }
// }

// func checkErr(err error){
// 	if err != nil {
// 		panic(err)
// 	}