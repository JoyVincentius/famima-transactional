package main

import (
	"database/sql"
	"log"

	"famima-transactional/controllers"
	"famima-transactional/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/karyawan_db")
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	// Ping database untuk memastikan koneksi berhasil
	if err := db.Ping(); err != nil {
		log.Fatal("Database is unreachable: ", err)
	}

	// Set database ke controller
	controllers.SetDB(db)

	// Inisialisasi Gin
	r := gin.Default()

	// Mendaftarkan rute
	routes.RegisterRoutes(r)

	log.Println("Server is running on port 8000")
	if err := r.Run(":8000"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
