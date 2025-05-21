package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
)

var db *sql.DB

type Feedback struct {
    Rating int    `json:"rating"`
    Note   string `json:"note"`
}

func main() {
    // Initialize database connection
    var err error
    db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    // Create feedback table if it doesn't exist
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS feedback (
            id SERIAL PRIMARY KEY,
            rating INT NOT NULL CHECK (rating >1 AND rating < 6),
            note TEXT
        )
    `)
    if err != nil {
        log.Fatalf("Failed to create table: %v", err)
    }

    http.HandleFunc("/submit-feedback", submitFeedbackHandler)
    http.HandleFunc("/", healthCheckHandler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Starting server on port %s...", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func submitFeedbackHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var feedback Feedback
    if err := json.NewDecoder(r.Body).Decode(&feedback); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if feedback.Rating < 1 || feedback.Rating > 6 {
        http.Error(w, "Rating must be between 1 and 6", http.StatusBadRequest)
        return
    }

	_, err := db.Exec(fmt.Sprintf("INSERT INTO feedback (rating, note) VALUES (%d, '%s')", feedback.Rating, feedback.Note))
	if err != nil {
		http.Error(w, "Failed to save feedback", http.StatusInternalServerError)
		log.Printf("Failed to save feedback: %v", err)
		return
	}

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Feedback submitted successfully"))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
}