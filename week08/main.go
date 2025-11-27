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
	http.HandleFunc("/fdump", fdump)
		http.HandleFunc("/ave",avehandler)

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

func avehandler(w http.ResponseWriter, r *http.Request) {
	var sum, tt, i, j int
	var ave float32
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	tokuten := strings.Split(r.FormValue("dd"), ",")
	fmt.Println(tokuten)
	var dosuu = [10]int{0,0,0,0,0,0,0,0,0,0}
	for i := range tokuten {
		tt, _ = strconv.Atoi(tokuten[i])
		sum += tt
		if tt >= 90 {
			dosuu[9]++
		} else if tt >= 80 {
			dosuu[8]++
		} else if tt >= 70 {
			dosuu[7]++
		} else if tt >= 60 {
			dosuu[6]++
		} else if tt >= 50 {
			dosuu[5]++
		} else if tt >= 40 {
			dosuu[4]++
		} else if tt >= 30 {
			dosuu[3]++
		} else if tt >= 20 {
			dosuu[2]++
		} else if tt >= 10 {
			dosuu[1]++
		} else {
			dosuu[0]++
		}
	}
	ave = float32(sum) / float32(len(tokuten))
	fmt.Fprint(w, "平均値は ")
	fmt.Fprintln(w, ave)
	fmt.Println(ave)
	for i = 9; i >= 0; i-- {
	    fmt.Fprint(w, i * 10, "点以上:")
	    for j = 0; j < dosuu[i]; j++ {
	        fmt.Fprint(w, "*")
	    }
	    fmt.Fprintln(w, "")
	}
}
