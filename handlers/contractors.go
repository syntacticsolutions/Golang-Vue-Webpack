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

func GetContractors(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Fetch tasks using our new model
        return c.JSON(http.StatusOK, models.GetContractors(db))
    }
}

func PostContractor(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {

        var contractor = models.Contractor{}

        c.Bind(&contractor)
        // // Add a task using our new model
        id, err := models.PostContractor(db, contractor)
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

func PutContractor(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        contractor_id, _ := strconv.Atoi(c.Param("id"))
        var contractor = models.Contractor{}

        c.Bind(&contractor)

        id, err := models.PutContractor(db, contractor, contractor_id)

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

// DeleteContractor endpoint
func DeleteContractor(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, _ := strconv.Atoi(c.Param("id"))
        // Use our new model to delete a task
        _, err := models.DeleteContractor(db, id)
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