package servehandler

import (
	SqlControl "SimpleAPIServer/DataBaseControl"
	datatype "SimpleAPIServer/DataType"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type DataHandler struct{}

func (h *DataHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	default:
		err = errors.New("Not Impelement")
	}
	if err != nil {
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	db := SqlControl.NewDataBaseConnect("SimpleDBOwner", "Newpassword")
	err = db.Open()
	if err != nil {
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT Data.Id, Data.DateAdded, Location.Lat, Location.Long FROM Data inner join Location On Data.Id=Location.Id;")
	datas := []datatype.Data{}
	for rows.Next() {
		_data := datatype.Data{}
		err = rows.Scan(&_data.ID, &_data.DateAdded, &_data.Location.Lat, &_data.Location.Long)
		if err != nil {
			return
		}
		datas = append(datas, _data)
	}
	output, err := json.Marshal(datas)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	db := SqlControl.NewDataBaseConnect("SimpleDBOwner", "Newpassword")
	err = db.Open()
	if err != nil {
		return
	}

	defer db.Close()
	decoder := json.NewDecoder(r.Body)
	var data datatype.Data
	err = decoder.Decode(&data)
	if err != nil {
		return
	}

	writeStrings := []string{
		"INSERT INTO Data(DateAdded) VALUES(NOW());",
		fmt.Sprintf("INSERT INTO Location() VALUES(LAST_INSERT_ID(), '%v', '%v')", data.Location.Lat, data.Location.Long),
	}

	err = db.Insert(writeStrings)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	return
}
