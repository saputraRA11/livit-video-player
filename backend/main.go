package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// Video structure simulating video data
type Video struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	URL       string `json:"url"`
	Source    string `json:"source"` // "youtube", "tiktok", or "direct"
	Thumbnail string `json:"thumbnail"`
}

// Mock dataset of videos
var videos = []Video{
	{ID: "1", Title: "sample", Author: "petlover", URL: "https://youtu.be/qYNweeDHiyU?si=g9Kbz2IRkE5M4DG5", Source: "youtube", Thumbnail: "https://example.com/cat-thumbnail.jpg"},
	{ID: "2", Title: "sample 2", Author: "dancer123", URL: "https://www.tiktok.com/@wahyuomdriverbali/video/7375568576914164998?is_from_webapp=1&sender_device=pc", Source: "tiktok", Thumbnail: "https://example.com/dance-thumbnail.jpg"},
	{ID: "3", Title: "Cooking Tips", Author: "chefdelight", URL: "https://example.com/sample-video.mp4", Source: "direct", Thumbnail: "https://example.com/cook-thumbnail.jpg"},
}

// CORS middleware to allow cross-origin requests
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Handler for searching videos based on keywords
func searchVideos(w http.ResponseWriter, r *http.Request) {
	keyword := r.URL.Query().Get("keyword")
	if keyword == "" {
		http.Error(w, "Keyword parameter is required", http.StatusBadRequest)
		return
	}

	var results []Video
	for _, video := range videos {
		// Search in title, author, and description fields
		if strings.Contains(strings.ToLower(video.Title), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(video.Author), strings.ToLower(keyword)) {
			results = append(results, video)
		}
	}

	// Set JSON response content type
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	http.Handle("/search", enableCORS(http.HandlerFunc(searchVideos)))
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
