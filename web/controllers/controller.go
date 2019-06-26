package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hyperledger/firstproject/blockchain"
)

// Application struct
type Application struct {
	Fabric *blockchain.FabricSetup
}

func renderTemplate(w http.ResponseWriter, r *http.Request, templateName string, data interface{}) {
	lp := filepath.Join("web", "templates", "layout.html")
	tp := filepath.Join("web", "templates", templateName)

	fmt.Println("\n Value of lp is: ", lp)
	fmt.Println("\n Value of tp is: ", tp)
	// Return a 404 if the template doesn't exist
	info, err := os.Stat(tp)
	//fmt.Println("Value of info is: ", reflect.TypeOf(info))
	//fmt.Printf("Value of info is: %s", info)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			fmt.Println("\n Inside error condition")
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	resultTemplate, err := template.ParseFiles(tp, lp)
	//fmt.Printf("\n Value of resultTemplate is: %s", resultTemplate)
	if err != nil {
		// Log the detailed error
		fmt.Println("\n ######### Inside resultTemplate error")
		fmt.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}
	if err := resultTemplate.ExecuteTemplate(w, "layout", data); err != nil {
		fmt.Println(err.Error())
		fmt.Println("\n ##### inside executeTemplate error in controller.go")
		http.Error(w, http.StatusText(500), 500)
	}
}
