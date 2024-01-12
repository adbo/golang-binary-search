package handlers

import (
  "main/data"
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/sirupsen/logrus/hooks/test"
  "github.com/stretchr/testify/assert"
)

func TestHandleNumberIndexRequest(t *testing.T) {
  logger, _ := test.NewNullLogger()
  store := data.NewNumberStore(logger)
  store.Numbers = []int{10, 20, 30, 40, 50}

  handler := NewNumberIndexHandler(store, logger)

  tests := []struct {
    name           string
    requestPath    string
    expectedStatus int
    expectedBody   string
  }{
    {"ValidRequest", "/endpoint/30", http.StatusOK, "2\n"},
    {"InvalidRequest", "/endpoint/", http.StatusBadRequest, "Invalid request format\n"},
    {"NonIntegerValue", "/endpoint/abc", http.StatusBadRequest, "Value must be an integer\n"},
    {"NotFoundValue", "/endpoint/60", http.StatusNotFound, "60\n"},
  }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
      req, _ := http.NewRequest("GET", tt.requestPath, nil)
      rr := httptest.NewRecorder()

      http.HandlerFunc(handler.HandleNumberIndexRequest).ServeHTTP(rr, req)

      assert.Equal(t, tt.expectedStatus, rr.Code, "Expected and actual status codes do not match")
      assert.Equal(t, tt.expectedBody, rr.Body.String(), "Expected and actual response bodies do not match")
    })
  }
}