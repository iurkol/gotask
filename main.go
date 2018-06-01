package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	LOGFILE    = "/tmp/requests.txt"
	SEPLINE    = "------------"
	LOGRECORD  = "IP: %s\nHostname: %s\nPath: %s\nTimestamp: %s\n%s\n"
	TIMELAYOUT = "2006-01-02 15:04:05 +0000"
)

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	logfile, err := os.OpenFile(LOGFILE, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic("Can't open log file.")
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf(LOGRECORD,
			strings.Split(r.RemoteAddr, ":")[0],
			strings.Split(r.Host, ":")[0],
			r.URL.Path,
			time.Now().Format(TIMELAYOUT),
			SEPLINE)
		w.Write([]byte("Hello World!"))
	})

	http.ListenAndServe(":9494", nil)
}
