package main

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/codegangsta/martini"
	"fmt"
	"net/http"
	"github.com/mrcoleman/go-learning/ldb"
)
func Get_Home(db *sql.DB, r *http.Request, rw http.ResponseWriter){
	var user = ldb.GetFirstUser(db)
	fmt.Fprintf(rw,"%s",user.Username)
}
func main(){
	m := martini.Classic()
	m.Map(ldb.SetupDB())
	m.Get("/",Get_Home)
	m.Run()
}
