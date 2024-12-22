package main

import (
	"fmt"
	"net/http"
	"os"
	"visit-counter/pkg/db"
	"visit-counter/pkg/dotenv"
	"visit-counter/pkg/routes"
)

func main() {
    dotenv.LoadEnv()

    db, err := db.GetDB(); if err != nil {
        fmt.Printf("Error connecting to database: %v\n", err)
        os.Exit(1)
    }

    http.HandleFunc("/create", routes.Create(db))
    http.HandleFunc("/count.svg", routes.GetCount(db))

    address := dotenv.GetEnv("ADDRESS", "0.0.0.0")
    port := dotenv.GetEnv("PORT", "8080")

    http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), nil)
    fmt.Printf("Server is running on http://%s:%s\n", address, port)
}