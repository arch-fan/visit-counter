package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"
	"visit-counter/pkg/db"
	"visit-counter/pkg/dotenv"
	"visit-counter/pkg/routes"
)

//go:embed assets/*.gif
var imagesFS embed.FS

func main() {
	dotenv.LoadEnv()

	vars := []string{"TOKEN", "SQLITE_FILE_PATH"}
	if missingVars := dotenv.GetMissingVars(vars...); len(missingVars) > 0 {
		fmt.Printf("Missing environment variables: %v\n", missingVars)
		os.Exit(1)
	}

	db, err := db.GetDB()
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		os.Exit(1)
	}

	http.HandleFunc("/create", routes.Create(db))
	http.HandleFunc("/counter.svg", routes.GetCount(db))

	address := dotenv.GetEnv("SERVER", "0.0.0.0")
	port := "8080"

	fmt.Printf("Server is running on http://%s:%s\n", address, port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), nil)
}
