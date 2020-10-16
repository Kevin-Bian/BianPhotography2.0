package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Kevin-Bian/BianPhotography2.0/src/helpers"
	"github.com/Kevin-Bian/BianPhotography2.0/src/models"
	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// Greet Default greet message
func Greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	res := "Hello There! Welcome to Bian Photography API V2."
	json.NewEncoder(w).Encode(res)
}

// createConnection Connects to our PSQL DB
func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}

// CreatePhoto Uploads a photo to collage
func CreatePhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var photo models.Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	insertID := helpers.InsertPhoto(photo)
	res := response{
		ID:      insertID,
		Message: "Photo created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

// GetPhoto Gets a photo given id
func GetPhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	user, err := helpers.GetPhoto(int64(id))
	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	json.NewEncoder(w).Encode(user)
}

// GetCollage Gets a collage given id
func GetCollage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	photos, err := helpers.GetCollage(params["id"])
	if err != nil {
		log.Fatalf("Unable to get all photo. %v", err)
	}

	json.NewEncoder(w).Encode(photos)
}

// GetAllPhoto Gets all photos uploaded
func GetAllPhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	photos, err := helpers.GetAllPhotos()
	if err != nil {
		log.Fatalf("Unable to get all photo. %v", err)
	}

	json.NewEncoder(w).Encode(photos)
}

// DeletePhoto Deletes a photo by id
func DeletePhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	status, err := helpers.DeletePhoto(int64(id))
	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	res := response{
		Message: status,
	}

	json.NewEncoder(w).Encode(res)
}
