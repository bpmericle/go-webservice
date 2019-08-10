package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	type args struct {
		w          *statusCodeResponseWriter
		r          *http.Request
		statusCode int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Returns200",
			args: args{
				w:          getResponseWriter(),
				r:          getRequest(),
				statusCode: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Health(tt.args.w, tt.args.r)
			if tt.args.statusCode != tt.args.w.statusCode {
				t.Fatalf("expected [%d] != actual [%d]", tt.args.statusCode, tt.args.w.statusCode)
			}
		})
	}
}

func getResponseWriter() *statusCodeResponseWriter {
	return newStatusCodeResponseWriter(httptest.NewRecorder())
}

func getRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/health", nil)
	return request
}

type statusCodeResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newStatusCodeResponseWriter(w http.ResponseWriter) *statusCodeResponseWriter {
	return &statusCodeResponseWriter{w, http.StatusInternalServerError}
}

func (lrw *statusCodeResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
