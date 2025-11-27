package main
import (
	"fmt"
	"net/http"
	"runtime"
  "strconv"
)

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/fdump", fdump)
		http.HandleFunc("/bmi",bmihandler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

func fdump(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	// フォームはマップとして利用でき以下で内容を確認できる．
	for k, v := range r.Form {
		fmt.Printf("%v : %v\n", k, v)
	}
}

func bmihandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	weight,_ := strconv.Atoi(r.FormValue("weight"))
	height,_ := strconv.Atoi(r.FormValue("height"))
	bmi := float32(weight) / ((float32(height) / 100.0) * (float32(height) / 100.0))
	fmt.Fprintln(w, bmi)
}
