package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/webfortune", webfortune)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}

}

func webfortune(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(5)
	if randomInt == 0 {
		fmt.Fprintln(w, "大吉")
	} else if randomInt == 1 {
		fmt.Fprintln(w, "中吉")
	} else if randomInt == 2 {
		fmt.Fprintln(w, "小吉")
	} else if randomInt == 3 {
		fmt.Fprintln(w, "末吉")
	} else if randomInt == 4 {
		fmt.Fprintln(w, "凶")
	}
}
