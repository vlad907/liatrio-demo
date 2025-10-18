package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	app := buildApp()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	body = bytes.TrimSpace(body)
	var payload map[string]any
	if err := json.Unmarshal(body, &payload); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if payload["message"] != "My name is "+candidateName {
		t.Fatalf("unexpected message: %v", payload["message"])
	}

	ts, ok := payload["timestamp"].(float64)
	if !ok || ts <= 0 {
		t.Fatalf("timestamp is not a positive number: %v", payload["timestamp"])
	}

	expected, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("failed to re-marshal payload: %v", err)
	}

	if !bytes.Equal(body, expected) {
		t.Fatalf("response body is not minified JSON. got=%s want=%s", body, expected)
	}
}
