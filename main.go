package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	TIMELAYOUT = "2006-01-02 15:04:05 +0000"

	INSERTQUERY = `INSERT INTO requests (ip, path, host, requested_at)
 VALUES ($1, $2, $3, $4);`

	GETQUERY = `SELECT ip, path, host, requested_at
	FROM requests
	ORDER BY id DESC
	LIMIT 25;`
)

type RequestInfo struct {
	Ip        string `db:"ip"`
	Path      string `db:"path"`
	Host      string `db:"host"`
	TimeStamp string `db:"requested_at"`
}

var Connection *sqlx.DB
var err error

func init() {
	dbUser := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	dbPasswd := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")

	postgresDataSource := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",
		dbUser, dbName, dbPasswd, dbHost)

	Connection, err = sqlx.Connect("postgres", postgresDataSource)
	if err != nil {
		log.Panicln("Error connecting to db:", err)
	}
}

func main() {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // todo: it is bulky; split to separate functions
		timeStamp := time.Now().Format(TIMELAYOUT)
		ip := strings.Split(r.RemoteAddr, ":")[0]
		host := strings.Split(r.Host, ":")[0]
		path := r.URL.Path

		Connection.MustExec(INSERTQUERY, ip, path, host, timeStamp)

		var result []RequestInfo
		err = Connection.Select(&result, GETQUERY)

		tmpl.Execute(w, result)
	})

	http.ListenAndServe(":9494", nil)
}
