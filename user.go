package main

import (
	"fmt"
	"time"

	"github.com/lib/pq"
)

func getUser(userID int) (User, error) {
	//Retrieve
	res := User{}

	var id int
	var name string
	var username string
	var active bool
	var password string
	var createdDate pq.NullTime
	var updatedDate pq.NullTime

	err := db.QueryRow(`SELECT id, f_name, user_name, pass ,active, created_time, updated_time FROM users where id = $1`, userID).Scan(&id, &name, &username, &password, &active, &createdDate, &updatedDate)
	if err == nil {
		res = User{ID: id, Name: name, Password: password, Username: username, Active: active, CreatedDate: createdDate.Time, UpdatedDate: updatedDate.Time}
	}

	return res, err
}

func allUsers() ([]User, error) {
	//Retrieve
	users := []User{}

	rows, err := db.Query(`SELECT id, f_name, user_name, pass ,active, created_time, updated_time FROM users order by id`)
	defer rows.Close()
	if err == nil {
		for rows.Next() {
			var id int
			var name string
			var username string
			var password string
			var active bool
			var createdDate pq.NullTime
			var updatedDate pq.NullTime

			err = rows.Scan(&id, &name, &username, &password, &active, &createdDate, &updatedDate)
			if err == nil {
				currentUser := User{ID: id, Name: name, Username: username, Active: active, Password: password, CreatedDate: createdDate.Time, UpdatedDate: updatedDate.Time}
				if createdDate.Valid {
					currentUser.CreatedDate = createdDate.Time
				}
				if updatedDate.Valid {
					currentUser.UpdatedDate = updatedDate.Time
				}
				users = append(users, currentUser)
			} else {
				return users, err
			}
		}
	} else {
		return users, err
	}

	return users, err
}

func insertUser(name string, username string, password string, active bool, createdTime time.Time, updatedTime time.Time) (int, error) {
	//Create

	fmt.Printf("name  %v\n", name)
	fmt.Printf("username  %v\n", username)
	var userID int
	err := db.QueryRow(`INSERT INTO users(f_name, user_name, pass, active, created_time, updated_time) VALUES($1, $2, $3, $4, $5 ,$6) RETURNING id`, name, username, password, active, createdTime, updatedTime).Scan(&userID)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Last inserted ID: %v\n", userID)
	return userID, err
}

func updateUser(id int, name, username string, password string, active bool, createdTime time.Time, updatedTime time.Time) (int, error) {
	//Update
	res, err := db.Exec(`UPDATE users set f_name=$1, user_name =$2, pass = $3 ,active =$4, created_time =$5, updated_time =$6 where id=$7 RETURNING id`, name, username, password, active, createdTime, updatedTime, id)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
}

func removeUser(UserID int) (int, error) {
	//Delete
	res, err := db.Exec(`delete from users where id = $1`, UserID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}
