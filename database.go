package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Trackt struct {
	Id        int       `json:"id"`
	UserId    int       `json:"userid"`
	ActionId  int       `json:"actionid"`
	ProjectId int       `json:"projectid"`
	TaskId    int       `json:"taskid"`
	Descr     string    `json:"descr"`
	CreatedAt time.Time `json:"createdat"`
}

type TrackRen struct {
	Project   string    `json:"project"`
	Task      string    `json:"task"`
	Action    string    `json:"action"`
	User      string    `json:"user"`
	Descr     string    `json:"descr"`
	Timestamp time.Time `json:"timestamp"`
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Function string `json:"function"`
}

type Project struct {
	Id          int
	Name        string `sql:"not null"`
	Description string
}

type Task struct {
	Id          int
	Name        string `sql:"not null"`
	Description string
}

type Action struct {
	Id   int
	Name string
}

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

type session struct {
	un           string
	lastActivity time.Time
}

type DropdownItem struct {
	A []actions
	P []project
	T []task
}

type actions = map[string]interface{}
type project = map[string]interface{}
type task = map[string]interface{}

var actionslice []actions
var projectslice []project
var taskslice []task

// Article - Our struct for all articles :::: this is an example
type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func createDdl() ([]map[string]interface{}, error) {

	statement := `SELECT * FROM actions ORDER BY Name ASC`
	rows, _ := db.Query(statement)
	cols, _ := rows.Columns()

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		// Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...]
		fmt.Println(m)
		actionslice = append(actionslice, m)
	}
	return actionslice, nil
}

func pdProject() ([]map[string]interface{}, error) {

	statement := `SELECT * FROM project ORDER BY Name ASC`
	rows, _ := db.Query(statement)
	cols, _ := rows.Columns()

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		// Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...]
		fmt.Println(m)
		projectslice = append(projectslice, m)
	}
	return projectslice, nil
}

func pdTask() ([]map[string]interface{}, error) {

	statement := `SELECT * FROM task ORDER BY Name ASC`
	rows, _ := db.Query(statement)
	cols, _ := rows.Columns()

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}

		// Outputs: map[columnName:value columnName2:value2 columnName3:value3 ...]
		fmt.Println(m)
		taskslice = append(taskslice, m)
	}
	return taskslice, nil
}

// Create a new tracktime record
func (track *Trackt) Create() (err error) {
	err = db.QueryRow("insert into trackt (usid, project_id, action_id, task_id, description, createdat) values ($1, $2, $3,$4,$5, $6) returning id", track.UserId, track.ProjectId, track.ActionId, track.TaskId, track.Descr, track.CreatedAt).Scan(&track.Id)
	if err != nil {
		log.Fatalf("Unable to connect to the DB: %v", err)
	}
	return
}

func SelectRecs(uid int) ([]*TrackRen, error) {

	statement := `SELECT actions.name, project.name, task.name, users.name, trackt.description, trackt.createdat
	FROM actions, project, task, users, trackt
	WHERE actions.id=trackt.action_id
	AND project.id=trackt.project_id
	AND task.id= trackt.task_id
	AND users.id= $1`

	rows, err := db.Query(statement, uid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ts := make([]*TrackRen, 0)
	for rows.Next() {
		t := new(TrackRen)
		err := rows.Scan(&t.Action, &t.Project, &t.Task, &t.User, &t.Descr, &t.Timestamp)

		if err != nil {
			log.Fatal(err)
		}

		ts = append(ts, t)
	}

	for _, t := range ts {
		fmt.Printf("%s, %s, %s, %s,%s\n", t.Action, t.Project, t.Task, t.User, t.Descr, t.Timestamp)
	}
	return ts, nil
}

func SelectUsers(uid int) ([]*User, error) {
	statement := `SELECT name FROM users WHERE id=$1`

	rows, err := db.Query(statement, uid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	ns := make([]*User, 0)
	for rows.Next() {
		n := new(User)
		err := rows.Scan(&n.Name)

		if err != nil {
			log.Fatal(err)
		}

		ns = append(ns, n)
	}

	fmt.Println(ns)
	return ns, nil

}
