package handlers

import (
	"database/sql"
    "net/http"
    "strconv"
    "github.com/dgrijalva/jwt-go"
    "go-echo-vue/models"
    "github.com/labstack/echo"
    "encoding/json"
    "time"
    // "fmt"
)

type H map[string]interface{}

type jwtCustomClaims struct {
    Email  string `json:"name"`
    Admin  bool   `json:"admin"`
    jwt.StandardClaims
}

// type User struct {
//     ID   int    `json:"id"`
//     User_type_id int `json:"user_type_id"`
//     First_name string `json:"first_name"`
//     Last_name string `json:"last_name"`
//     Email string `json:"email"`
//     Password string `json:"email"`
//     Phone string `json:"email"`
//     Cell string `json:"email"`
// }


func GetUsers(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Fetch tasks using our new model
        return c.JSON(http.StatusOK, models.GetUsers(db))
    }
}

func Login(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {

        json_map := make(map[string]interface{})
        err := json.NewDecoder(c.Request().Body).Decode(&json_map)
        
        checkErr(err)

        email := json_map["email"]
        password := json_map["password"]
        
        var uid int
        var useremail string
        var userpass string
        //TODO BCrypt

        err = db.QueryRow(`
            SELECT id, email, password FROM users WHERE email = ? AND password = ?
            `, email, password).Scan(&uid, &useremail, &userpass)

        checkErr(err)

        if email == useremail && password == userpass {

            // Set custom claims
            claims := &jwtCustomClaims{
                useremail,
                true, // TODO map admin to this variable
                jwt.StandardClaims{
                    ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
                },
            }

            // Create token with claims
            token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

            // Generate encoded token and send it as response.
            t, err := token.SignedString([]byte("fa39ps8ndf#$@#0p8@sd9s08df0-na88fnasdfn-a2080"))
            if err != nil {
                return err
            }
            return c.JSON(http.StatusOK, echo.Map{
                "token": t,
            })
        }

    return echo.ErrUnauthorized

    }
    
}

// PutTask endpoint
func PostUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var user = models.User{}

		c.Bind(&user)
        // // Add a task using our new model
        id, err := models.PostUser(db, user)
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

func PutUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user_id, _ := strconv.Atoi(c.Param("id"))
		var user = models.User{}

		c.Bind(&user)

		id, err := models.PutUser(db, user, user_id)

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

// DeleteTask endpoint
func DeleteUser(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, _ := strconv.Atoi(c.Param("id"))
        // Use our new model to delete a task
        _, err := models.DeleteUser(db, id)
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