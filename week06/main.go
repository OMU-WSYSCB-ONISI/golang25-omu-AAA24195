package main
import (
	"fmt"
	"net/http"
	"runtime"
  "strconv"
  "strings"
)

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/hello", hellohandler)
        http.HandleFunc("/enq", enqhandler)
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

func enqhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	// r.FormValue("name")として，フォーム中name欄の値を得る
	fmt.Fprintln(w, r.FormValue("name")+"さん，ご協力ありがとうございます.\n年齢は"+r.FormValue("age")+"で，性別は"+r.FormValue("gend")+"で，出身地は"+r.FormValue("birthplace")+"ですね")
}

func cal00handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	price, _ := strconv.Atoi(r.FormValue("price"))
	num, _ := strconv.Atoi(r.FormValue("num"))
	fmt.Fprint(w, "合計金額は ")
	fmt.Fprintln(w, price*num)
}

func bmihandler(w http.ResponseWriter, r *http.Reques) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}

}
