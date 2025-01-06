package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func testMiddleware(t *testing.T) {
	// 建立一個 HTTP 測試伺服器
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// 建立一個 HTTP 客戶端
	client := server.Client()

	// 調用 HTTP 客戶端發送請求
	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fatalf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 驗證回應的 HTTP 狀態碼
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

}
