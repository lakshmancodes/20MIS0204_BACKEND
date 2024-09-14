package youtube

import (
    "context"
    "database/sql"
    "log"
    "time"
    "google.golang.org/api/option"
    "google.golang.org/api/youtube/v3"
)

// Video struct (assuming this struct is already defined in your codebase)
type Video struct {
    VideoID      string
    Title        string
    Description  string
    PublishedAt  time.Time
    ThumbnailURL string
}

// fetchYouTubeVideos fetches videos from the YouTube API and stores them in the database
func fetchYouTubeVideos(apiKey, query string, db *sql.DB) {
    service, err := youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
    if err != nil {
        log.Fatalf("Error creating YouTube client: %v", err)
    }

    call := service.Search.List([]string{"snippet"}).Q(query).Order("date").MaxResults(50)
    response, err := call.Do()
    if err != nil {
        log.Fatalf("Error making search API call: %v", err)
    }

    for _, item := range response.Items {
        // Parse the published date to time.Time
        publishedAt, err := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
        if err != nil {
            log.Printf("Error parsing PublishedAt for video %s: %v", item.Snippet.Title, err)
            continue
        }

        videoID := item.Id.VideoId
        title := item.Snippet.Title
        description := item.Snippet.Description
        thumbnailURL := item.Snippet.Thumbnails.Default.Url

        // Check if the table exists before inserting
        _, err = db.Exec(`CREATE TABLE IF NOT EXISTS videos (
            video_id VARCHAR(255) PRIMARY KEY,
            title VARCHAR(255),
            description TEXT,
            published_at DATETIME,
            thumbnail_url VARCHAR(255)
        )`)
        if err != nil {
            log.Printf("Error creating table if not exists: %v", err)
        }
        
        // Insert video into the database
        _, err = db.Exec(`
            INSERT INTO videos (video_id, title, description, published_at, thumbnail_url)
            VALUES (?, ?, ?, ?, ?)
            ON DUPLICATE KEY UPDATE 
                title = VALUES(title), 
                description = VALUES(description), 
                published_at = VALUES(published_at), 
                thumbnail_url = VALUES(thumbnail_url)
        `, videoID, title, description, publishedAt, thumbnailURL)

        if err != nil {
            log.Printf("Error inserting video %s into database: %v", title, err)
        }
    }
}

// startFetching starts a background process that fetches videos periodically
func StartFetching(apiKey, query string, db *sql.DB) {
    ticker := time.NewTicker(10 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ticker.C:
            fetchYouTubeVideos(apiKey, query, db)
        }
    }
}
