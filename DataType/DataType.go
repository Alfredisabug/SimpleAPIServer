package datatype

import "time"

type Data struct {
	ID       string
	Location struct {
		Lat  float32
		Long float32
	}
	DateAdded time.Time
}
