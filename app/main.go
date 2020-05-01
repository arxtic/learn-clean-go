package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

/*
	Init function it gonna execute to prepare all the main function needs,
	like a "constructor"
*/
func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on debug mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	dbSSL := viper.GetString(`database.sslmode`)

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

	log.Fatal(e.Start(viper.GetString("server.address")))
}

func testServer(c echo.Context) error {
	return c.String(http.StatusOK, "Ready to go 🚀")
}
