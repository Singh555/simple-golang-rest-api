package main

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// User struct represents the user model
type User struct {
	ID       int    `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Email    string `db:"email" json:"email"`
}

var db *sqlx.DB

func initDB() {
	var err error

	// Replace your PostgreSQL connection details here
	connectionString := "user=postgres password=Jarvismark2@ dbname=fr sslmode=disable"
	db, err = sqlx.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("***********Connected to the database************")

	// Run migrations
	err = runMigrations(db)
	if err != nil {
		log.Fatal(err)
	}
}

func runMigrations(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // replace with your actual migrations path
		"postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	fmt.Println("*************Migrations ran successfully*************")
	return nil
}

func main() {
	// Initialize the database connection and run migrations
	initDB()

	// Create a new Gin router
	router := gin.Default()

	// Define a route to create a new user
	router.POST("/users", createUser)

	// Run the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := router.Run(":" + port)
	if err != nil {
		fmt.Println("error starting the server ", err)
		return
	}
}

func createUser(c *gin.Context) {
	var user User

	// Bind the JSON data from the request body to the user struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Insert the user into the database
	err := db.QueryRowx("INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id", user.Username, user.Email).Scan(&user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created user with the inserted ID
	c.JSON(http.StatusCreated, user)
}
