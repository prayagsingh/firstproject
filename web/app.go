package web

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hyperledger/firstproject/web/controllers"
)

// Serve controllers
func Serve(app *controllers.Application) {
	fmt.Println("\n ########## Inside Server function")
	r := mux.NewRouter()
	// Declare the static file directory and point it to the assets directory
	staticDirectory := http.Dir("web/assets")

	// Looking for files in assest directory
	fs := http.FileServer(staticDirectory)
	fmt.Println("Value of fs is: ", fs)
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	r.HandleFunc("/home.html", app.HomeHandler).Methods("GET")
	r.HandleFunc("/request.html", app.RequestHandler).Methods("POST")
	// exposing API's
	r.HandleFunc("/Index", app.GetValue).Methods("GET")
	r.HandleFunc("/changevalue", app.ChangeValue).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home.html", http.StatusTemporaryRedirect)
	})

	fmt.Println("Listening (http://localhost:3000/) ...")
	http.ListenAndServe(":3000", r)
}
