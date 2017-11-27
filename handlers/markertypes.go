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

func GetMarkerTypes(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Fetch tasks using our new model
        return c.JSON(http.StatusOK, models.GetMarkerTypes(db))
    }
}

func PostMarkerType(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {

        var marker_type = models.MarkerType{}

        c.Bind(&marker_type)
        // // Add a task using our new model
        id, err := models.PostMarkerType(db, marker_type)
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

func PutMarkerType(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        marker_type_id, _ := strconv.Atoi(c.Param("id"))
        var marker_type = models.MarkerType{}

        c.Bind(&marker_type)

        id, err := models.PutMarkerType(db, marker_type, marker_type_id)

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

// Deletemarker_type endpoint
func DeleteMarkerType(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, _ := strconv.Atoi(c.Param("id"))
        // Use our new model to delete a task
        _, err := models.DeleteMarkerType(db, id)
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