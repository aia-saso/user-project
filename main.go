package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "go"
)

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	tmpDB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db = tmpDB
}

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	r := mux.NewRouter()
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("www/assets"))))

	r.HandleFunc("/users", handleListUsers).Methods("GET")
	r.HandleFunc("/user", handleViewUser).Methods("GET")
	r.HandleFunc("/save", handleSaveUser).Methods("POST")
	r.HandleFunc("/delete", handleDeleteUser).Methods("DELETE")

	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
func handleListUsersJSON(w http.ResponseWriter, req *http.Request) {
	b, err := json.Marshal(allUsers)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprint(w, "hello 22")
	json.NewEncoder(w).Encode(b)
}
func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go Web Server")
	w.WriteHeader(200)
}
