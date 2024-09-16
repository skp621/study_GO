//説明
//http.HandleFunc("/", handler)は、ルートパス（/）にアクセスされた際にhandler関数を呼び出す設定です。
//handler関数内でfmt.Fprintf(w, "HELLO WORLD")により、HTTPレスポンスとして「HELLO WORLD」がクライアントに返されます。
//http.ListenAndServe(":8080", nil)は、localhost:8080でHTTPサーバーを起動します。

package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO WORLD")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
