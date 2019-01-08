package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func rectime(w http.ResponseWriter, r *http.Request) {

	queryValues := r.URL.Query()
	ac, _ := strconv.Atoi(queryValues.Get("a"))
	pj, _ := strconv.Atoi(queryValues.Get("p"))
	ta, _ := strconv.Atoi(queryValues.Get("t"))

	fmt.Printf("registration of %s done!\n", queryValues.Get("a"))

	rt := time.Now()
	track := Trackt{UserId: 1, ActionId: ac, ProjectId: pj, TaskId: ta, CreatedAt: rt}
	track.Create()

	et := tpl.ExecuteTemplate(w, "tmpl.html", track)

	if et != nil {
		http.Error(w, et.Error(), http.StatusInternalServerError)
		return
	}
}

func postime(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {

		e1 := tpl.ExecuteTemplate(w, "rectime.html", d)
		if e1 != nil {
			http.Error(w, e1.Error(), http.StatusInternalServerError)
			return
		}

	} else {
		// logic part of log in

		pj, _ := strconv.Atoi(r.FormValue("project"))
		tk, _ := strconv.Atoi(r.FormValue("task"))
		at, _ := strconv.Atoi(r.FormValue("action"))
		ds := r.FormValue("description")
		us := 1
		rt := time.Now()

		fmt.Printf("Action: %d\n", at)
		fmt.Printf("Project: %d\n", pj)
		fmt.Printf("Task: %d\n", tk)
		fmt.Printf("Description: %s\n", ds)
		fmt.Printf("record time: %s\n", rt)

		track := Trackt{UserId: us, ActionId: at, ProjectId: pj, TaskId: tk, Descr: ds, CreatedAt: rt}
		track.Create()

		//redirect to listrecs

		http.Redirect(w, r, "/sel", http.StatusSeeOther)
	}
}

func listRecs(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	//uid, _ := strconv.Atoi(queryValues.Get("u"))
	uid := 1 //user=1, me

	ts, _ := SelectRecs(uid)
	json.NewEncoder(w).Encode(ts)
	fmt.Printf("returned array of trackt=%d\n", ts)
}

func listUsers(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	uid := 1 //user=1, me
	ns, _ := SelectUsers(uid)
	fmt.Printf("returned array of trackt=%d\n", ns)
}

func getUsers(c *gin.Context) {

}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	fmt.Println("Endpoint Hit: returnAllArticles")

	json.NewEncoder(w).Encode(articles)
}
