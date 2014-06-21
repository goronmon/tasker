package main

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"time"
)

var tasks = []Task{}

type Task struct {
	TaskName      sql.NullString `db:"task_name"`
	TaskCreated   time.Time      `db:"task_created"`
	TaskTimeSpent int64          `db:"task_time_spent"`
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

	err = db.Select(&tasks, "SELECT task_name, task_created, task_time_spent FROM task")

	if err != nil {
		log.Print(err)
	}

	// Debug section
	task1 := Task{}
	if len(tasks) > 0 {
		task1 = tasks[0]
	}
	fmt.Printf("%#v\n%#v", task1, tasks)

	http.HandleFunc("/tasker/", taskerHandler)
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}
