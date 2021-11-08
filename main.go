package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Users struct {
	gorm.Model
	Name          string
	Email         string `gorm:"typevarchar(100);unique_index"`
	UserAudioList []AudioFile
}

type AudioFile struct {
	gorm.Model
	Title       string
	Description string
	Category    string
	AudioFile   string //path of file from s3 bucket.
	UserID      int
}

var db *gorm.DB
var err error

func main() {

	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	dbpassword := os.Getenv("PASSWORD")

	awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	fmt.Println(awsAccessKeyID)
	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, dbpassword, dbPort)

	// Openning connection to database
	db, err = gorm.Open(dialect, dbURI)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database successfully")
	}

	// Close the databse connection when the main function closes
	defer db.Close()

	// Make migrations to the database if they haven't been made already
	db.AutoMigrate(&Users{})
	db.AutoMigrate(&AudioFile{})

	router := mux.NewRouter()

	router.HandleFunc("/audios", GetAudioList).Methods("GET")
	router.HandleFunc("/audio/{id}", GetAudio).Methods("GET")
	router.HandleFunc("/audio/upload", UploadAudio).Methods("POST")
	router.HandleFunc("/audio/delete/{id}", DeleteAudio).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

//get invidual audio
func GetAudio(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user Users
	var audio []AudioFile

	db.First(&user, params["id"])
	db.Model(&user).Related(&audio)

	user.UserAudioList = audio

	json.NewEncoder(w).Encode(&user)
}

//get all audio func
func GetAudioList(w http.ResponseWriter, r *http.Request) {
	var audio []AudioFile

	db.Find(&audio)

	json.NewEncoder(w).Encode(&audio)
}

//create audio func
func UploadAudio(w http.ResponseWriter, r *http.Request) {
	var audio AudioFile
	json.NewDecoder(r.Body).Decode(&audio)

	uploadedAudio := db.Create(&audio)
	err = uploadedAudio.Error
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(&uploadedAudio)
}

//delete audio func
func DeleteAudio(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var audio AudioFile

	db.First(&audio, params["id"])
	db.Delete(&audio)

	json.NewEncoder(w).Encode(&audio)
}
