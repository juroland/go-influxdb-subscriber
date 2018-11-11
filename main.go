package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/influxdata/influxdb/models"
)

func printPoint(p models.Point) {
	fmt.Printf("Measurement: %v\n", string(p.Name()))
	fields, err := p.Fields()
	if err == nil {
		fmt.Printf("Fields: %v\n", fields)
	}
	tags := p.Tags()
	if err == nil {
		fmt.Printf("Tags: %v\n", tags)
	}
}

func writePointsHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	points, _ := models.ParsePoints(body)
	for _, p := range points {
		printPoint(p)
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Println("subscriber started")
	http.HandleFunc("/write", writePointsHandler)
	log.Fatal(http.ListenAndServe(Config.Address, nil))
}
