package ldb

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type UserEntity struct {
	Username string
	Password string
	Email string
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
