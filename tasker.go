package main

import (
    _ "github.com/lib/pq"
    "net/http"
    "html/template"
    "github.com/jmoiron/sqlx"
    "time"
    "log"
)

var tasks = []Task{}

type Page struct {
    Title string
    Text string
}

type Task struct {
    taskID int
    taskName string
    taskCreated time.Time
    taskTimeSpent int64 

}

func taskerHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("tasker.html")
    t.Execute(w, tasks)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/tasker", http.StatusFound)
}


func main() {
    db, err := sqlx.Connect("postgres", "user=postgres password=mario64 host=localhost port=5432 dbname=tasker sslmode=disable")
   
    if err != nil {
        log.Fatalln(err)
    }

    db.Select(&tasks, "SELECT * FROM task")

    http.HandleFunc("/tasker/", taskerHandler)
    http.HandleFunc("/", rootHandler)
    http.ListenAndServe(":8080", nil)
}
