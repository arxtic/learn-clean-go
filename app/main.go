package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

/*
	Init function it gonna execute to prepare all the main function needs,
	like a "constructor"
*/
func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbSSL := os.Getenv("DB_SSL")

	/*
		Connection database defined by url
	*/
	connection := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbName, dbSSL)
	dbConn, err := sql.Open(`postgres`, connection)

	/*
		logging the connection info
	*/
	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	/*
		Preparing the server
	*/
	e := echo.New()
	e.GET("/test", testServer)

	log.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}

func testServer(c echo.Context) error {
	return c.String(http.StatusOK, "Ready to go 🚀")
}
