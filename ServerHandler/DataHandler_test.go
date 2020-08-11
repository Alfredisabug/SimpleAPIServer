package servehandler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerGet(t *testing.T) {
	req, err := http.NewRequest("GET", "/data", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handleGet(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	t.Log("Get data: ", rr.Body.String())
}