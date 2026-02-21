package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	
	"github.com/Bechan94/go-Notes-Api/internal/notes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil { log.Fatal("Error loading .env") }

	db, err := sql.Open("postgres",
		"host="+os.Getenv("DB_HOST")+
			" port="+os.Getenv("DB_PORT")+
			" user="+os.Getenv("DB_USER")+
			" password="+os.Getenv("DB_PASSWORD")+
			" dbname="+os.Getenv("DB_NAME")+
			" sslmode=disable")
	if err != nil { log.Fatal(err) }
	defer db.Close()

	repo := notes.NewRepository(db)
	service := notes.NewService(repo)
	handler := notes.NewHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/notes", handler.GetNotesHandler).Methods("GET")
	r.HandleFunc("/notes/{id}", handler.GetNoteHandler).Methods("GET")
	r.HandleFunc("/notes", handler.CreateNoteHandler).Methods("POST")
	r.HandleFunc("/notes/{id}", handler.UpdateNoteHandler).Methods("PUT")
	r.HandleFunc("/notes/{id}", handler.DeleteNoteHandler).Methods("DELETE")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}