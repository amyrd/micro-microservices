package main

import "fmt"
import {
	"database/sql"
	"fmt"
	"log"
	"google.golang.org/grpc"
}
import "database/sql"
import "log"

// hardcoding credentials just for development purposes
const (
	dbDriver   = "mysql"
	dbUser     = "root"
	dbPassword = "Admin123"
	dbName     = "user"
)

var db *sql.DB

func main() {
	var err error

	// open a db connection
	dsn := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", dbUser, dbPassword, dbName)
	db, err = sql.Open(dbDriver, dsn)

	// if theres error it will exit the app totally
	if err != nil {
		log.Fatal(err)
	}

	// in the even of a shutdown this will run right before app closes officially to clean up stuff
	defer func() {
		// if error closing db, print error
		if err = db.Close(); err != nil {
			log.Printf("Error closing db: %s", err)
		}
	}()

	// ping db to check conenction
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer(

}
