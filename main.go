package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

// Employee is a struct for employee info
type Employee struct {
	Id     int    `json:"id" required`
	Gender string `json:"gender" required`
}

func main() {
	r := setupRouter()
	// Check for port environment variable or run on 8080 as a default
	port := os.Getenv("BEA_PORT")
	if port == "" {
		port = "8080"
	}
	// Run on localhost explicitly so others on network aren't able to access.
	r.Run("localhost:" + port)
}

func openDBConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./employees.db")
	if err != nil {
		return nil, fmt.Errorf("error opening databases %v", err)
	}
	return db, nil
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	// Gin-gonic trusts all proxies by default so here establishing an empty allowlist.
	r.SetTrustedProxies([]string{})

	r.GET("/employees", func(c *gin.Context) {
		result, err := getEmployees()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, result)
	})

	return r
}

func getEmployees() ([]Employee, error) {
	db, err := openDBConnection()
	if db != nil {
		defer db.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("establishing db connection: %v", err)
	}

	rows, err := db.Query("SELECT * FROM employees")
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("executing query: %v", err)
	}

	result := make([]Employee, 0)

	for rows.Next() {
		currentRecord := Employee{}
		if err := rows.Scan(&currentRecord.Id, &currentRecord.Gender); err != nil {
			return nil, fmt.Errorf("scan: %v", err)
		}

		result = append(result, currentRecord)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterating: %v", err)
	}
	return result, nil
}
