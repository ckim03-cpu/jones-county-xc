package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"jones-xc-backend/db"

	_ "github.com/mattn/go-sqlite3"
)

var queries *db.Queries

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func athletesHandler(w http.ResponseWriter, r *http.Request) {
	athletes, err := queries.GetAllAthletes(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(athletes)
}

func athleteByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid athlete id", http.StatusBadRequest)
		return
	}

	athlete, err := queries.GetAthleteByID(r.Context(), id)
	if err == sql.ErrNoRows {
		http.Error(w, "athlete not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(athlete)
}

func meetsHandler(w http.ResponseWriter, r *http.Request) {
	meets, err := queries.GetAllMeets(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(meets)
}

func resultsByMeetHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid meet id", http.StatusBadRequest)
		return
	}

	results, err := queries.GetResultsByMeet(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func createResultHandler(w http.ResponseWriter, r *http.Request) {
	var params db.CreateResultParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	result, err := queries.CreateResult(r.Context(), params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

func topTimesHandler(w http.ResponseWriter, r *http.Request) {
	times, err := queries.GetTopTimes(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(times)
}

func initDB(dbPath string) *sql.DB {
	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Fatal("Could not read schema.sql: ", err)
	}
	if _, err := conn.Exec(string(schema)); err != nil {
		log.Fatal("Could not apply schema: ", err)
	}

	return conn
}

func main() {
	dbPath := "jones-xc.db"
	if envPath := os.Getenv("DB_PATH"); envPath != "" {
		dbPath = envPath
	}

	conn := initDB(dbPath)
	defer conn.Close()
	queries = db.New(conn)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)
	mux.HandleFunc("GET /api/athletes", corsMiddleware(athletesHandler))
	mux.HandleFunc("GET /api/athletes/{id}", corsMiddleware(athleteByIDHandler))
	mux.HandleFunc("GET /api/meets", corsMiddleware(meetsHandler))
	mux.HandleFunc("GET /api/meets/{id}/results", corsMiddleware(resultsByMeetHandler))
	mux.HandleFunc("POST /api/results", corsMiddleware(createResultHandler))
	mux.HandleFunc("GET /api/top-times", corsMiddleware(topTimesHandler))

	log.Println("Backend server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
