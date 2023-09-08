package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Slack struct {
	SlackName     string `json:"slack_name"`
	CurrentDay    string `json:"current_day"`
	UtcTime       string `json:"utc_time"`
	Track         string `json:"track"`
	GithubFileUrl string `json:"github_file_url"`
	GithubRepoUrl string `json:"github_repo_url"`
	StatusCode    int    `json:"status_code"`
}

func getSlack(w http.ResponseWriter, r *http.Request) {
	var slack Slack
	slackName := r.URL.Query().Get("slack_name")
	track := r.URL.Query().Get("track")

	slack.SlackName = slackName
	slack.Track = track
	slack.CurrentDay = time.Now().UTC().Format("Monday")

	currentTime := time.Now().UTC()
	currentTime = currentTime.Add(time.Duration(rand.Intn(5)-2) * time.Minute)
	slack.UtcTime = currentTime.Format("2006-01-02T15:04:05Z")

	slack.GithubFileUrl = "https://github.com/dastardlyjockey/Hngx-1/main.go"
	slack.GithubRepoUrl = "https://github.com/dastardlyjockey/Hngx-1"
	slack.StatusCode = 200

	data, err := json.Marshal(&slack)
	if err != nil {
		log.Println("Failed to marshal slack data")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(data)
}

func main() {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	router.Get("/api", getSlack)

	port := "3001"
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Println("Starting server on port: ", port)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
