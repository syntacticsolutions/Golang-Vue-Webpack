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

func GetProjectManagers(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Fetch tasks uProjectManagersing our new model
        return c.JSON(http.StatusOK, models.GetProjectManagers(db))
    }
}

func PostProjectManager(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {

        var project_manager = models.ProjectManager{}

        c.Bind(&project_manager)
        // // Add a task using our new model
        id, err := models.PostProjectManager(db, project_manager)
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

// DeleteProject endpoint
func DeleteProjectManager(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, _ := strconv.Atoi(c.Param("id"))
        // Use our new model to delete a task
        _, err := models.DeleteProjectManager(db, id)
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