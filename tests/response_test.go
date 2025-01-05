package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/q10242/MiniBasicBackendGolang/lib"
	// 引用被測試的套件
)

// TestJsonResponse tests the JsonResponse function
func TestJsonResponse(t *testing.T) {
	// 建立一個 HTTP 測試回應寫入器
	w := httptest.NewRecorder()

	// 調用 JsonResponse 函式
	lib.JsonResponse(w, http.StatusOK, "Test Message", "SUCCESS")

	// 驗證回應的 HTTP 狀態碼
	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// 驗證回應的 Content-Type
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected Content-Type application/json, got %s", w.Header().Get("Content-Type"))
	}

	// 驗證回應的 JSON 內容
	var response lib.GenericResponse
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Message != "Test Message" {
		t.Errorf("expected message 'Test Message', got '%s'", response.Message)
	}
	if response.Status != "SUCCESS" {
		t.Errorf("expected status 'SUCCESS', got '%s'", response.Status)
	}
}
