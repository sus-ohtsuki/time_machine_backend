package main
import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"

	"crypto/rand"
	"database/sql"
    "encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load("./env/dev/.env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

	db, err := sql.Open("postgres", "host=db user=test_admin password=password dbname=test_tmm_db sslmode=disable")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	} else {
		fmt.Println("Successfully connected to database")
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World4",
		})
	})

	r.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World4",
		})
	})
	r.Run(":8080")
}
