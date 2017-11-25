package main

import (
	// "net/http"
	"database/sql"
	"go-echo-vue/handlers"
    "github.com/labstack/echo"
    "github.com/go-sql-driver/mysql"
    "fmt"
)

func main() {

	db := initDB()
    migrate(db)

    // Create a new instance of Echo
    e := echo.New()

    // e.File("/", "dist/index.html")
    e.Static("/", "dist")
    e.GET("/tasks", handlers.GetTasks(db))
    e.PUT("/tasks", handlers.PutTask(db))
    e.DELETE("/tasks/:id", handlers.DeleteTask(db))
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

func migrate(db *sql.DB) {
    stmt, err := db.Prepare(`
    CREATE TABLE IF NOT EXISTS tasks(
        id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
        name VARCHAR(25) NOT NULL
    );
    `)
    
    if err != nil {
        panic(err)
    }

    res, err := stmt.Exec()
    
    if err != nil {
        panic(err)
    } else {
    fmt.Println(res)
    }
}