package main

import (
    "fmt"
    "log"
    "flag"
    "runtime"
    "net/http"
    "database/sql"
    "route"
    "db_handler"
    "service_handler"
)

var(
    nwork = flag.Int("worker", 20, "number of worker")
)

func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())

	var err error
	db_handler.Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/piara")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
    flag.Parse()
    fmt.Println("Starting the dispatcher")
    service_handler.StartDispatcher(*nwork)

    router := route.NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}