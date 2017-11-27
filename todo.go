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
    Email  string `json:"name"`
    Admin  bool   `json:"admin"`
    jwt.StandardClaims
}

func main() {

    // Create a new instance of Echo
    e := echo.New()

    db := initDB();

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{"*"},
    }))

    // Login route
    e.POST("/login", handlers.Login(db))
    e.POST("/Register", handlers.PostUser(db))

    // Restricted group
    r := e.Group("/api")
    config := middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: []byte("fa39ps8ndf#$@#0p8@sd9s08df0-na88fnasdfn-a2080"),
    }
    r.Use(middleware.JWTWithConfig(config))
    r.GET("", restricted)

    e.Static("/", "dist")
    r.GET("/api/users", handlers.GetUsers(db))
    r.POST("/api/users", handlers.PostUser(db))
    r.PUT("/api/users/:id", handlers.PutUser(db))
    r.DELETE("/api/users/:id", handlers.DeleteUser(db))

    r.GET("/api/projects", handlers.GetProjects(db))
    r.POST("/api/projecs", handlers.PostProject(db))
    r.PUT("/api/projects/:id", handlers.PutProject(db))
    r.DELETE("/api/projects/:id", handlers.DeleteProject(db))

    r.GET("/api/markers", handlers.GetMarkers(db))
    r.POST("/api/markers", handlers.PostMarker(db))
    r.PUT("/api/markers/:id", handlers.PutMarker(db))
    r.DELETE("/api/markers/:id", handlers.DeleteMarker(db))

    r.GET("/api/marker_types", handlers.GetMarkerTypes(db))
    r.POST("/api/marker_types", handlers.PostMarkerType(db))
    r.PUT("/api/marker_types/:id", handlers.PutMarkerType(db))
    r.DELETE("/api/marker_types/:id", handlers.DeleteMarkerType(db))

    r.GET("/api/user_types", handlers.GetUserTypes(db))
    r.POST("/api/user_types", handlers.PostUserType(db))
    r.PUT("/api/user_types/:id", handlers.PutUserType(db))
    r.DELETE("/api/user_types/:id", handlers.DeleteUserType(db))

    r.GET("/api/contractors", handlers.GetContractors(db))
    r.POST("/api/contractors", handlers.PostContractor(db))
    r.PUT("/api/contractors/:id", handlers.PutContractor(db))
    r.DELETE("/api/contractors/:id", handlers.DeleteContractor(db))

    r.GET("/api/project_managers", handlers.GetProjectManagers(db))
    r.POST("/api/project_managers", handlers.PostProjectManager(db))
    r.DELETE("/api/project_managers/:id", handlers.DeleteProjectManager(db))

    
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
    email := claims.Email
    return c.String(http.StatusOK, "Welcome "+email+"!")
}

func checkErr(err error) {
        if err != nil {
            panic(err)
        }
    }