package SqlControl

import (
	"log"
	"testing"
	"time"
)

type Data struct {
	ID       string
	Location struct {
		Lat  float32
		Long float32
	}
	DateAdded time.Time
}

func TestOpenDB(t *testing.T) {
	db := NewDataBaseConnect("SimpleDBOwner", "Newpassword")
	db.Open()
	defer db.Close()

	db.Ping()
	return
}

func TestQueryDB(t *testing.T) {
	db := NewDataBaseConnect("SimpleDBOwner", "Newpassword")
	db.Open()
	defer db.Close()

	rows, err := db.Query("SELECT Data.Id, Data.DateAdded, Location.Lat, Location.Long FROM Data inner join Location On Data.Id=Location.Id;")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		_data := Data{}
		err = rows.Scan(&_data.ID, &_data.DateAdded, &_data.Location.Lat, &_data.Location.Long)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(_data)
	}
	return
}

func TestInsertDB(t *testing.T) {
	db := NewDataBaseConnect("SimpleDBOwner", "Newpassword")
	db.Open()
	defer db.Close()

	cmdStrings := []string{
		"INSERT INTO Data(DateAdded) VALUES(NOW());",
		"INSERT INTO Location() VALUES(LAST_INSERT_ID(), '100.0', '200.0')",
	}

	err := db.Insert(cmdStrings)
	if err != nil {
		log.Fatal(err)
	}
	return
}
