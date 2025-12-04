package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"net/http"
	"os"
	"runtime"
	"time"
)

const logFile = "public/logs.json" // データの保存先 --- (*1)

// Log 掲示板に保存するデータを構造体で定義 --- (*2)
type Log struct {
	ID    string `json:"id"`//stringに変更
	Name  string `json:"name"`
	Body  string `json:"body"`
	CTime int64  `json:"ctime"`
}

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.Handle("/", http.FileServer(http.Dir("public/")))
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/bbs", showHandler)
	http.HandleFunc("/write", writeHandler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}


}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

// 時間に基づいて64進数のIDを生成する関数
func createTimeBasedID() string {
	timestampNano := time.Now().UnixNano()
	timestampStr := fmt.Sprintf("%d", timestampNano)

	encodedID := base64.URLEncoding.EncodeToString([]byte(timestampStr))
	for len(encodedID) > 0 && encodedID[len(encodedID)-1] == '=' {
		encodedID = encodedID[:len(encodedID)-1]
	}

	return encodedID
}

// 書き込みログを画面に表示する --- (*6)
func showHandler(w http.ResponseWriter, r *http.Request) {
	// ログを読み出してHTMLを生成 --- (*7)
	htmlLog := ""
	logs := loadLogs() // データを読み出す

	for index, i := range logs {
		number := index + 1

		htmlLog += fmt.Sprintf(
			"<p>(%d) <span>%s [%s]</span>: %s --- %s</p>",
			number,
			html.EscapeString(i.Name),
			i.ID,
			html.EscapeString(i.Body),
			time.Unix(i.CTime, 0).Format("2006/1/2 15:04"))
	}

	// HTML全体を出力 --- (*8)
	htmlBody := "<html><head><style>" +
		"p { border: 1px solid silver; padding: 1em;} " +
		"span { background-color: #eef; } " +
		"</style></head><body><h1>BBS</h1>" +
		getForm() + htmlLog + "</body></html>"
	 if _, err := w.Write([]byte(htmlBody)); err !=nil {
		fmt.Printf("Error writing response body in showHandler: %v", err)
	 }
}

// フォームから送信された内容を書き込み --- (*9)
func writeHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data.", http.StatusBadRequest)
		fmt.Printf("Error parsing form in writeHandler: %v", err)
		return
	} // フォームを解析 --- (*10)
	var log Log
	log.Name = r.Form["name"][0]
	log.Body = r.Form["body"][0]
	if log.Name == "" {
		log.Name = "名無し"
	}
	logs := loadLogs() // 既存のデータを読み出し --- (*11)

	log.ID = createTimeBasedID()
	log.CTime = time.Now().Unix()

	logs = append(logs, log)          // 追記 --- (*12)
	saveLogs(logs)                    // 保存
	http.Redirect(w, r, "/bbs", 302) // リダイレクト --- (*13)
}

// 書き込みフォームを返す --- (*14)
func getForm() string {
	return "<div><form action='/write' method='get'>" +
		"名前: <input type='text' name='name'><br>" +
		"本文: <input type='text' name='body' style='width:30em;'><br>" +
		"<input type='submit' value='書込'>" +
		"</form></div><hr>"
}

// ファイルからログファイルの読み込み --- (*15)
func loadLogs() []Log {
	text, err := os.ReadFile(logFile)
	if err != nil {
		return make([]Log, 0)
	}
	// JSONをパース --- (*16)
	var logs []Log
	if err := json.Unmarshal([]byte(text), &logs); err !=nil {
		fmt.Printf("error unmarshaling logs: %v", err)
	}
	return logs
}

// ログファイルの書き込み --- (*17)
func saveLogs(logs []Log) {
	//JSONにエンコード
	bytes, _ := json.Marshal(logs)
	//ファイルへ書き込む
	if err := os.WriteFile(logFile, bytes, 0644); err != nil {
		fmt.Printf("error writing logs to file: %v", err)
	}
}
