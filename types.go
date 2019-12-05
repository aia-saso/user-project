package main

import "time"

//IndexPage represents the content of the index page, available on "/"
//The index page shows a list of all books stored on db
type IndexPage struct {
	AllUsers []User
}

//userPage represents the content of the user page, available on "/user.html"
//The book page shows info about a given book
type UserPage struct {
	TargetUser User
}

//user represents a book object
type User struct {
	ID          int       `json:"id.omitempty" bson:"id"`
	Name        string    `json:"fname.omitempty" bson:"fname"`
	Username    string    `json:"username.omitempty" bson:"username"`
	Password    string    `json:"password.omitempty" bson:"password"`
	CreatedDate time.Time `json:"created_time.omitempty" bson:"createdTime"`
	UpdatedDate time.Time `json:"updated_time.omitempty" bson:"updatedTime"`
	Active      bool      `json:"active.omitempty" bson:"active"`
}

//PublicationDateStr returns a sanitized Publication Date in the format YYYY-MM-DD
func (b User) CreatedDateStr() string {
	return b.CreatedDate.Format("2006-01-02")
}

//ErrorPage represents shows an error message, available on "/book.html"
type ErrorPage struct {
	ErrorMsg string
}
