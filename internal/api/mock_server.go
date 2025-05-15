package api

import (
"net/http"
"net/http/httptest"
"os"
)

func startMockAPIServer() *httptest.Server {
return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// Validate auth header in CI environment
if os.Getenv("CI") == "true" {
if auth := r.Header.Get("Authorization"); auth != "Bearer test_key_ci" {
w.WriteHeader(http.StatusUnauthorized)
return
}
}

w.Header().Set("Content-Type", "application/json")
w.Write([]byte(`{
"choices": [{
"message": {
"content": "Mocked Chinese phishing email content",
"role": "assistant"
}
}],
"usage": {
"prompt_tokens": 10,
"completion_tokens": 20
}
}`))
}))
}
