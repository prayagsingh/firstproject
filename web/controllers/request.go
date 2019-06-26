package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// RequestHandler func
func (app *Application) RequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n Inside RequestHandler function")
	data := &struct {
		TransactionID string
		Success       bool
		Response      bool
	}{
		TransactionID: "",
		Success:       false,
		Response:      false,
	}
	fmt.Println("\n Value of r.formValue is: ", r.FormValue("submitted"))
	if r.FormValue("submitted") == "true" {
		helloValue := r.FormValue("hello")
		txid, err := app.Fabric.InvokeHello(helloValue)
		if err != nil {
			http.Error(w, "Unable to invoke hello in the blockchain", 500)
		}
		data.TransactionID = txid
		data.Success = true
		data.Response = true
	}
	renderTemplate(w, r, "request.html", data)
}

// ChangeValue for changing and sending the values to blockchain
func (app *Application) ChangeValue(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\n Inside ChangeValue func")
	data := &struct {
		TransactionID string
		Success       bool
		Response      bool
	}{
		TransactionID: "",
		Success:       false,
		Response:      false,
	}

	// StoreBody struct for storing the Body value
	type StoreBody struct {
		Value string
	}
	var storebody StoreBody

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	log.Println("Getting body is: ", body)
	if err != nil {
		log.Fatalln("Error in ChangeValue", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = r.Body.Close(); err != nil {
		log.Fatalln("Error in Closing Body", err)
	}

	if err = json.Unmarshal(body, &storebody); err != nil {
		log.Fatalln("Error in ChangeValue when unmarshalling data", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Print("\n Value of storebody.value is: ", storebody.Value)

	txid, err := app.Fabric.InvokeHello(storebody.Value)
	if err != nil {
		http.Error(w, "Unable to invoke hello in the blockchain", 500)
	}
	data.TransactionID = txid
	data.Success = true
	data.Response = true

	// Sending response to client
	dataJSON, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	w.Write(dataJSON)
	return
}
