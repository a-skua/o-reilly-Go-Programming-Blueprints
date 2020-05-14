package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

// templは1つのテンプレートを表します
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTPはHTTPリクエストを処理します
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		// 最初に実行されるときにコンパイルされる
		// ただし、エラーが発生居売る場合には初期化時に実行される方が好ましい場合もある
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	err := t.templ.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	r := newRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// チャットルームを開始します
	go r.run()
	// Webサーバーを開始します
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
