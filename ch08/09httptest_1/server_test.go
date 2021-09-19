package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleGet(t *testing.T) {
	// マルチプレクサを生成
	mux := http.NewServeMux()

	// テスト対象のハンドラhendleRequestをマルチプレクサに付与
	mux.HandleFunc("/post/", handleRequest)

	// レコーダを生成
	writer := httptest.NewRecorder()

	// リクエストを作成
	request, _ := http.NewRequest("GET", "/post/1", nil)

	// テスト対象のハンドラにリクエストを送信し、レコーダに記録
	mux.ServeHTTP(writer, request)

	// レコーダにより結果をチェック
	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Errorf("Cannot retrieve JSON post")
	}
}

func TestHandlePut(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
	request, _ := http.NewRequest("PUT", "/post/1", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
}
