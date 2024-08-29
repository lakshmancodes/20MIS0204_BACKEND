package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"youtube-fetcher-trademarkia/internal/server"
	database "youtube-fetcher-trademarkia/internal/sqldatabase"
	youtube "youtube-fetcher-trademarkia/internal/youtube"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a specified .env file
func LoadEnv(envFilePath string) map[string]string {
	// Load the .env file
	err := godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Create a map to store environment variables
	envVars := make(map[string]string)

	// List of keys to retrieve
	keys := []string{"MYSQL_CONN", "DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "API_KEY", "FETCH_INTERVAL"}

	// Retrieve values from environment variables
	for _, key := range keys {
		value := os.Getenv(key)
		envVars[key] = value
	}

	return envVars
}

func main() {
	// Load environment variables
	envVars := LoadEnv(".env")

	// Access and print values
	fmt.Println("MySQL Connection:", envVars["MYSQL_CONN"])
	fmt.Println("Database Host:", envVars["DB_HOST"])
	fmt.Println("Database Port:", envVars["DB_PORT"])
	fmt.Println("Database User:", envVars["DB_USER"])
	fmt.Println("Database Password:", envVars["DB_PASSWORD"])
	fmt.Println("Database Name:", envVars["DB_NAME"])
	fmt.Println("API Key:", envVars["API_KEY"])
	fmt.Println("Fetch Interval:", envVars["FETCH_INTERVAL"])

	// Initialize the database (no arguments needed)
	database.InitDB()

	//connect the api and fetch the results from API
	go youtube.StartFetching(envVars["API_KEY"], "your-search-query", database.DB)

	// Initialize the router
	r := mux.NewRouter()

	// Define the API routes
	r.HandleFunc("/api/videos", server.GetPaginatedVideosHandler).Methods("GET")
	r.HandleFunc("/api/videos/search", server.SearchVideosHandler).Methods("GET")

	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", r)
}
