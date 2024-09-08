package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// リクエストボディを読み取る
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}
		// 受け取ったデータを表示
		fmt.Printf("********************\n")
		fmt.Printf("Received POST data: %s\n", string(body))

		// リクエストヘッダーを取得して表示
		fmt.Printf("********************\n")
		for name, values := range r.Header {
			for _, value := range values {
				if name == "Sampleheader" {
					fmt.Printf("Sampleheader: %s\n", value)
				}
			}
		}

		// ヘッダー情報をレスポンスとして返す
		// fmt.Fprintf(w, "Received headers:\n")
		fmt.Fprintf(w, "********************\n")
		fmt.Fprintf(w, "[Request Header]\n")
		for name, values := range r.Header {
			for _, value := range values {
				fmt.Fprintf(w, "%s: %s\n", name, value)
			}
		}

		// クライアントにレスポンスを返す
		fmt.Fprintf(w, "********************\n")
		fmt.Fprintf(w, "[Request Body]\n")
		fmt.Fprintf(w, "%s", string(body))
	} else {
		// POSTリクエスト以外のリクエストに対する処理
		fmt.Printf("********************\n")
		fmt.Fprintf(w, "Only POST method is supported.")
	}
}

func handler_2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Good Afternoon, World!")
}

func main() {
	http.HandleFunc("/", handler)    // ハンドラ関数を設定
	http.HandleFunc("/2", handler_2) // ハンドラ関数を設定
	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
