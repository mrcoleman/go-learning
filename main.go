package main

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/codegangsta/martini"
	"fmt"
	"net/http"
)

type UserEntity struct {
	username string
	password string
	email string
}

func SetupDB() *sql.DB{
	db,err := sql.Open("sqlite3","./temp.db")
	if(err != nil){
		panic(err)
	}
	return db
}

func GetFirstUser(db *sql.DB) *UserEntity {
	rows, err := db.Query("select username, password, email from users;")
	if(err != nil){
		panic(err)
	}
	defer rows.Close()
	rows.Next()
	var username string
	var password string
	var email string
	rows.Scan(&username, &password, &email)
	return &UserEntity {username, password, email}
}

func main(){
	m := martini.Classic()
	m.Map(SetupDB())
	m.Get("/",func(db *sql.DB, r *http.Request, rw http.ResponseWriter) {
		var user = GetFirstUser(db)
		fmt.Fprintf(rw,"%s",user.username)
	})
	m.Run()
}
