package main

import (
	"net/http"
	"database/sql"
	"go-echo-vue/handlers"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/dgrijalva/jwt-go"
    _ "github.com/go-sql-driver/mysql"
    // "time"
    // "fmt"
    "io/ioutil"
    "strings"
    // "encoding/json"
)

type jwtCustomClaims struct {
    email  string `json:"name"`
    admin  bool   `json:"admin"`
    jwt.StandardClaims
}

func main() {

    // Create a new instance of Echo
    e := echo.New()

    db := initDB();

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Login route
    e.POST("/login", handlers.Login(db))
    e.POST("/Register", handlers.PostUser(db))

    // Restricted group
    r := e.Group("/restricted")
    r.Use(middleware.JWT([]byte("secret")))
    r.GET("", restricted)

    e.Static("/", "dist")
    e.GET("/api/users", handlers.GetUsers(db))
    e.POST("/api/users", handlers.PostUser(db))
    e.PUT("/api/users/:id", handlers.PutUser(db))
    e.DELETE("/api/users/:id", handlers.DeleteUser(db))

    e.GET("/api/projects", handlers.GetProjects(db))
    e.POST("/api/projecs", handlers.PostProject(db))
    e.PUT("/api/projects/:id", handlers.PutProject(db))
    e.DELETE("/api/projects/:id", handlers.DeleteProject(db))

    e.GET("/api/markers", handlers.GetMarkers(db))
    e.POST("/api/markers", handlers.PostMarker(db))
    e.PUT("/api/markers/:id", handlers.PutMarker(db))
    e.DELETE("/api/markers/:id", handlers.DeleteMarker(db))

    e.GET("/api/marker_types", handlers.GetMarkerTypes(db))
    e.POST("/api/marker_types", handlers.PostMarkerType(db))
    e.PUT("/api/marker_types/:id", handlers.PutMarkerType(db))
    e.DELETE("/api/marker_types/:id", handlers.DeleteMarkerType(db))

    e.GET("/api/user_types", handlers.GetUserTypes(db))
    e.POST("/api/user_types", handlers.PostUserType(db))
    e.PUT("/api/user_types/:id", handlers.PutUserType(db))
    e.DELETE("/api/user_types/:id", handlers.DeleteUserType(db))

    e.GET("/api/contractors", handlers.GetContractors(db))
    e.POST("/api/contractors", handlers.PostContractor(db))
    e.PUT("/api/contractors/:id", handlers.PutContractor(db))
    e.DELETE("/api/contractors/:id", handlers.DeleteContractor(db))

    e.GET("/api/project_managers", handlers.GetProjectManagers(db))
    e.POST("/api/project_managers", handlers.PostProjectManager(db))
    e.DELETE("/api/project_managers/:id", handlers.DeleteProjectManager(db))

    
    // e.POST("/api/projects", handlers.PostProject(db))
    // e.GET("/init", migrate(db))
    // Start as a web server
    e.Logger.Fatal(e.Start(":1323"))
}

func initDB() *sql.DB { // add support for null

    db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang_vue")


    // Here we check for any db errors then exit
    if err != nil {
        panic(err)
    }

    // If we don't get any errors but somehow still don't get a db connection
    // we exit as well
    if db == nil {
        panic("db nil")
    }

    return db
}

func migrate(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {

        file, err := ioutil.ReadFile("./migrations/migrations.sql")

        if err != nil {
            return err
        }

        requests := strings.Split(string(file), ";")

        for _, request := range requests {
            _, err := db.Query(request + ";")
            if err != nil {
                return err
            }
        }


        return c.String(http.StatusOK, "Successfully Migrated")

    }

}


func accessible(c echo.Context) error {
    return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
    user := c.Get("user").(*jwt.Token)
    claims := user.Claims.(*jwtCustomClaims)
    email := claims.email
    return c.String(http.StatusOK, "Welcome "+email+"!")
}

func checkErr(err error) {
        if err != nil {
            panic(err)
        }
    }