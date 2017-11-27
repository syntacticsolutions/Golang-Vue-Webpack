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

func GetProjects(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Fetch tasks using our new model
        return c.JSON(http.StatusOK, models.GetProjects(db))
    }
}

func PostProject(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {

        var project = models.Project{}

        c.Bind(&project)
        // // Add a task using our new model
        id, err := models.PostProject(db, project)
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

func PutProject(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        project_id, _ := strconv.Atoi(c.Param("id"))
        var project = models.Project{}

        c.Bind(&project)

        id, err := models.PutProject(db, project, project_id)

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

// DeleteProject endpoint
func DeleteProject(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, _ := strconv.Atoi(c.Param("id"))
        // Use our new model to delete a task
        _, err := models.DeleteProject(db, id)
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