package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func handleSaveUser(w http.ResponseWriter, r *http.Request) {
	var id = 0
	var err error

	r.ParseForm()
	// params := r.PostForm
	// params := r.URL.Query()
	query := r.URL.Query()

	fmt.Println("GET parameters string : ", query)
	idStr := query.Get("id")

	//////////////////

	// first := query.Get("name")
	// second := query.Get("id")
	// w.Write([]byte("First value : " + first + "\n"))
	// w.Write([]byte("Second value : " + second + "\n"))
	// // because query is a map, we can use it like a hash table
	// // map[first:[1] second:[2]]
	// // query by the key "first" example
	// firstvalue := query["first"]
	// fmt.Println(firstvalue)
	////////////////

	fmt.Printf("mmmmm idStr  :%v\n", idStr)
	if len(idStr) > 0 {
		id, err = strconv.Atoi(idStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}

	name := query.Get("name")
	username := query.Get("userName")
	active := query.Get("Active")
	password := query.Get("Password")
	createdDateStr := query.Get("createdDate")
	updatedDateStr := query.Get("updatedDate")

	fmt.Printf("name  %v\n", name)
	fmt.Printf("username  %v\n", username)
	activev := false

	if active == "true" {
		activev = true
	}

	var createdDate time.Time
	if len(createdDateStr) > 0 {
		createdDate, err = time.Parse("2006-01-02", createdDateStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}
	var updatedDate time.Time
	if len(updatedDateStr) > 0 {
		updatedDate, err = time.Parse("2006-01-02", updatedDateStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}
	opname := ""
	if id == 0 {
		createdDate = time.Now()
		_, err = insertUser(name, username, password, activev, createdDate, updatedDate)
		opname = "created"
	} else {
		updatedDate = time.Now()
		_, err = updateUser(id, name, username, password, activev, createdDate, updatedDate)
		opname = "updated"
	}

	if err != nil {
		renderErrorPage(w, err)
		return
	}
	io.WriteString(w, opname+" Succ ")
	// http.Redirect(w, r, "/", 302)
}

func handleListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := allUsers()
	if err != nil {
		renderErrorPage(w, err)
		return
	}
	e, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}
	io.WriteString(w, "USERS "+string(e))
	// buf, err := ioutil.ReadFile("templates/index.html")
	// if err != nil {
	// 	renderErrorPage(w, err)
	// 	return
	// }

	// var page = IndexPage{AllUsers: users}
	// indexPage := string(buf)
	// t := template.Must(template.New("indexPage").Parse(indexPage))
	// t.Execute(w, page)
}

func handleViewUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")

	var currentUser = User{}
	currentUser.CreatedDate = time.Now()

	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}

		currentUser, err = getUser(id)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
	}

	e, err := json.Marshal(currentUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	io.WriteString(w, "USER "+string(e))
	// buf, err := ioutil.ReadFile("templates/user.html")
	// if err != nil {
	// 	renderErrorPage(w, err)
	// 	return
	// }

	// var page = UserPage{TargetUser: currentUser}
	// userPage := string(buf)
	// t := template.Must(template.New("userPage").Parse(userPage))
	// err = t.Execute(w, page)
	// if err != nil {
	// 	renderErrorPage(w, err)
	// 	return
	// }
}

func handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")
	fmt.Println(idStr)
	if len(idStr) > 0 {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			renderErrorPage(w, err)
			return
		}

		n, err := removeUser(id)
		if err != nil {
			renderErrorPage(w, err)
			return
		}
		io.WriteString(w, "Rows removed: "+strconv.Itoa(n))
		// fmt.Printf("Rows removed: %v\n", n)
	}
	// http.Redirect(w, r, "/", 302)
}

func renderErrorPage(w http.ResponseWriter, errorMsg error) {
	buf, err := ioutil.ReadFile("templates/error.html")
	if err != nil {
		log.Printf("%v\n", err)
		fmt.Fprintf(w, "%v\n", err)
		io.WriteString(w, "Error ")
		return
	}

	var page = ErrorPage{ErrorMsg: errorMsg.Error()}
	errorPage := string(buf)
	t := template.Must(template.New("errorPage").Parse(errorPage))
	t.Execute(w, page)
}
