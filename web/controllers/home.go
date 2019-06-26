package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HomeHandler func
func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("\n Inside HomeHandler Function")
	helloValue, err := app.Fabric.QueryHello()
	fmt.Print("\n Value of hello Value is: ", helloValue, "\n")
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	data := &struct {
		Hello string
	}{
		Hello: helloValue,
	}
	renderTemplate(w, r, "home.html", data)
}

// GetValue method for exposing the API using curl command
func (app *Application) GetValue(w http.ResponseWriter, r *http.Request) {
	fmt.Print("\n Inside GetValue Handler function")
	helloValue, err := app.Fabric.QueryHello()
	fmt.Print("\n Value of hello Value is: ", helloValue, "\n")
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}
	data := &struct {
		Hello string
	}{
		Hello: helloValue,
	}
	dataJSON, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(dataJSON)
	return
}
