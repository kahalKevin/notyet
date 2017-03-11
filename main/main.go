package main

import (
    "log"
    "net/http"
    "database/sql"
    "route"
    "db_handler"
)

func init() {
	var err error
	db_handler.Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/piara")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

    router := route.NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}