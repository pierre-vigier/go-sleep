package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/sleep", func(w http.ResponseWriter, req *http.Request) {
		log.Println("Start of sleep handler")
		defer log.Println("Sleep handler processing done")
		seconds, err := strconv.Atoi(req.FormValue("time"))
		if err != nil {
			log.Printf("Could not get sleep time: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("nok"))
			return
		}
		if seconds < 0 {
			seconds = 0
		}
		log.Printf("Will sleep for %d seconds", seconds)

		time.Sleep(time.Duration(seconds) * time.Second)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Println("Listen on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
