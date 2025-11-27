package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/info", info)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func info(w http.ResponseWriter, r *http.Request) {
currentTime := time.Now().Format("15:04")
burauza := r.Header.Get("User-Agent")
fmt.Fprintln(w, "今の時刻は",currentTime, "で、利用しているブラウザは",burauza, "ですね")
}
