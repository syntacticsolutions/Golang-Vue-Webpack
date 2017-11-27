package handlers

import (
	"database/sql"
    "net/http"
    "strconv"
    "go-echo-vue/models"
    "github.com/labstack/echo"
    // "encoding/json"
    // "time"
)

func GetUserTypes(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Fetch tasks using our new model
        return c.JSON(http.StatusOK, models.GetUserTypes(db))
    }
}

func PostUserType(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {

        var user_type = models.UserType{}

        c.Bind(&user_type)
        // // Add a task using our new model
        id, err := models.PostUserType(db, user_type)
        // Return a JSON response if successful
        if err == nil {
            return c.JSON(http.StatusCreated, H{
                "created": id,
            })
        // Handle any errors
        } else {
            return err
        }
    }
}

func PutUserType(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        user_type_id, _ := strconv.Atoi(c.Param("id"))
        var user_type = models.UserType{}

        c.Bind(&user_type)

        id, err := models.PutUserType(db, user_type, user_type_id)

        if err == nil {
            return c.JSON(http.StatusCreated, H{
                "created": id,
            })
        // Handle any errors
        } else {
            return err
        }
    }
}

// Deleteuser_type endpoint
func DeleteUserType(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, _ := strconv.Atoi(c.Param("id"))
        // Use our new model to delete a task
        _, err := models.DeleteUserType(db, id)
        // Return a JSON response on success
        if err == nil {
            return c.JSON(http.StatusOK, H{
                "deleted": id,
            })
        // Handle errors
        } else {
            return err
        }
    }
}

func checkErr(err error){
    if err != nil {
        panic(err)
    }
}