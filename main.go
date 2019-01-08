package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var err error
var db *sql.DB
var tpl *template.Template
var dbUsers = map[string]user{}       // user ID, user
var dbSessions = map[string]session{} // session ID, session
var dbSessionsCleaned time.Time
var d DropdownItem

const sessionLength int = 30

// connect to the Db
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	dbSessionsCleaned = time.Now()

	//connect to the database and manage the errors
	connStr := "user=trackt dbname=trackt password=trackt sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	projects, _ := pdProject()
	tasks, _ := pdTask()
	actions, _ := createDdl()

	d = DropdownItem{
		P: projects,
		T: tasks,
		A: actions,
	}

}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:10000",
	}
	// application logic handlers
	http.HandleFunc("/rec", rectime)
	http.HandleFunc("/sel", listRecs)
	http.HandleFunc("/entry", postime)
	http.HandleFunc("/users", listUsers)
	http.HandleFunc("/articles", returnAllArticles)

	// session logic handlers
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	server.ListenAndServe()
}
