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

func GetMarkers(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Fetch tasks using our new model
        return c.JSON(http.StatusOK, models.GetMarkers(db))
    }
}

func PostMarker(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {

        var marker = models.Marker{}

        c.Bind(&marker)
        // // Add a task using our new model
        id, err := models.PostMarker(db, marker)
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

func PutMarker(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        marker_id, _ := strconv.Atoi(c.Param("id"))
        var marker = models.Marker{}

        c.Bind(&marker)

        id, err := models.PutMarker(db, marker, marker_id)

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

// DeleteMarker endpoint
func DeleteMarker(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, _ := strconv.Atoi(c.Param("id"))
        // Use our new model to delete a task
        _, err := models.DeleteMarker(db, id)
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