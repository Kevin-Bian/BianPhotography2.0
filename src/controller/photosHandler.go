package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Kevin-Bian/BianPhotography2.0/src/models"
	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func Greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	res := "Hello There! Welcome to Bian Photography API V2."
	json.NewEncoder(w).Encode(res)
}

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

func CreatePhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var photo models.Photo
	err := json.NewDecoder(r.Body).Decode(&photo)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	insertID := insertPhoto(photo)
	res := response{
		ID:      insertID,
		Message: "Photo created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetPhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	user, err := getPhoto(int64(id))
	if err != nil {
		log.Fatalf("Unable to get user. %v", err)
	}

	json.NewEncoder(w).Encode(user)
}

func GetCollage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	photos, err := getCollage(params["id"])
	if err != nil {
		log.Fatalf("Unable to get all photo. %v", err)
	}

	json.NewEncoder(w).Encode(photos)
}

func GetAllPhoto(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	photos, err := getAllPhotos()
	if err != nil {
		log.Fatalf("Unable to get all photo. %v", err)
	}

	json.NewEncoder(w).Encode(photos)
}

func insertPhoto(photo models.Photo) int64 {

	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO photos (collageid, name, link, description) VALUES ($1, $2, $3, $4) RETURNING photoid`
	var id int64

	err := db.QueryRow(sqlStatement, photo.CollageID, photo.Name, photo.Link, photo.Description).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)
	return id
}

func getPhoto(id int64) (models.Photo, error) {
	// create the postgres db connection
	db := createConnection()
	defer db.Close()

	var photo models.Photo
	sqlStatement := `SELECT * FROM photos WHERE photoid=$1`

	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&photo.ID, &photo.CollageID, &photo.Name, &photo.Link, &photo.Description)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return photo, nil
	case nil:
		return photo, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return photo, err
}

func getAllPhotos() ([]models.Photo, error) {
	db := createConnection()
	defer db.Close()

	var photos []models.Photo
	sqlStatement := `SELECT * FROM photos`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var photo models.Photo
		err = rows.Scan(&photo.ID, &photo.CollageID, &photo.Name, &photo.Link, &photo.Description)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		photos = append(photos, photo)

	}
	return photos, err
}

func getCollage(collageid string) ([]models.Photo, error) {
	db := createConnection()
	defer db.Close()

	var photos []models.Photo

	sqlStatement := `SELECT * FROM photos where collageid=$1`

	rows, err := db.Query(sqlStatement, collageid)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var photo models.Photo
		err = rows.Scan(&photo.ID, &photo.CollageID, &photo.Name, &photo.Link, &photo.Description)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		photos = append(photos, photo)

	}
	return photos, err
}
