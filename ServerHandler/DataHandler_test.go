package servehandler

import (
	datatype "SimpleAPIServer/DataType"
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

func TestHandlerPost(t *testing.T) {
	rand.Seed(int64(time.Now().UnixNano()))

	reqData := datatype.Data{
		Location: struct {
			Lat  float32
			Long float32
		}{
			Lat:  rand.Float32(),
			Long: rand.Float32(),
		},
	}
	reqBody := new(bytes.Buffer)
	json.NewEncoder(reqBody).Encode(reqData)
	req := httptest.NewRequest(
		http.MethodPost,
		"/Data",
		reqBody,
	)
	rr := httptest.NewRecorder()
	handlePost(rr, req)
}
