package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"triples/http_utils"
)

func TestDELETE(t *testing.T) {
	log.SetOutput(&nullWriter{})

	tests := []struct {
		name         string
		requestURL   string
		expectedBody string
		expectedCode int
	}{
		{
			name:       "DELETE request without session 1",
			requestURL: "/",
			expectedBody: "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
				"<Response>\n" +
				"  <Code>403</Code>\n" +
				"  <Messege>403 Forbidden</Messege>\n" +
				"</Response>\n",
			expectedCode: http.StatusForbidden,
		},
		{
			name:       "DELETE request without session 2",
			requestURL: "/////////////////////////////////",
			expectedBody: "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
				"<Response>\n" +
				"  <Code>403</Code>\n" +
				"  <Messege>403 Forbidden</Messege>\n" +
				"</Response>\n",
			expectedCode: http.StatusForbidden,
		},
		{
			name:       "DELETE request without session 3",
			requestURL: "/:/:/:/:/",
			expectedBody: "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
				"<Response>\n" +
				"  <Code>403</Code>\n" +
				"  <Messege>403 Forbidden</Messege>\n" +
				"</Response>\n",
			expectedCode: http.StatusForbidden,
		},
		{
			name:       "DELETE request without session 4",
			requestURL: "/123",
			expectedBody: "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
				"<Response>\n" +
				"  <Code>403</Code>\n" +
				"  <Messege>403 Forbidden</Messege>\n" +
				"</Response>\n",
			expectedCode: http.StatusForbidden,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", tt.requestURL, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(http_utils.Handler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedCode)
			}

			if rr.Body.String() != tt.expectedBody {
				t.Errorf("handler returned unexpected body: got\n %v\n want\n %v\n", rr.Body.String(), tt.expectedBody)
			}
		})
	}
}
