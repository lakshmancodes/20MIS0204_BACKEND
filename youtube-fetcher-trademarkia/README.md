1. The YouTube Fetcher Trademarkia project is a Go-based application designed to fetch and manage YouTube video data. 
2. The project includes a MySQL database for storing video details and a set of API endpoints for interacting with the data. 
3. The application is containerized using Docker and Docker Compose.
4. I have used Postman to test the API endpoints
5. Used sql workbench ide to store the DB in sqldb

So the project structure goes like
youtube-fetcher-trademarkia/
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go
├── .env
├── internal/
    ├── sqldatabase/
    │   └── database.go
    └── server/
    │   ├── server.go
    │   └── handlers.go
    └── youtube.go

Dockerfile: Defines the Docker image for the Go application

docker-compose.yml: Defines the services for the application, including the Go application and MySQL database.

go.mod & go.sum: Go modules files for dependency management according to the needs

main.go: The entry point for the Go application, yes there can be only one main

.env: Environment variables configuration file.

internal/sqldatabase/database.go: Contains database connection and related functions.

internal/server/server.go: Contains server setup and middleware

internal/server/handlers.go: Contains API handlers for different endpoints

youtube.go: Contains functions for interacting with the YouTube API.


Create and Configure the .env File
# MySQL Environment Variables
MYSQL_CONN="root:xxx@tcp(localhost:3306)/youtube_videos"
DB_HOST=db
DB_PORT=3307
DB_USER=root
DB_PASSWORD=xxx
DB_NAME=youtube_videos

# YouTube API Key
API_KEY=AIzaSyBhrDtIxAOBKHKRquclq0j6c3XUEdhr2QQ

# Fetch Interval (in minutes)
FETCH_INTERVAL=10

Build and Start the Docker Containers

docker-compose build
docker-compose up

Once the docker starts you can monitor the sessions in the docker desktop
This will build the Docker images and start the containers for both the application and the MySQL database.

API Endpoints:

GET http://localhost:8080/api/videos?page=1&limit=10
this end point Fetches a paginated list of videos

GET http://localhost:8080/api/videos/search?q=your-search-query
this end point fetches the Searches videos by title, description, and other filters
