package Rest_API

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)



func Start() {
	
	fmt.Println("Implementing simple rest API")

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/{param1}/{param2}", paramsModify)
	
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page")
}

func paramsModify(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["param1"]
	age := params["param2"]
	fmt.Println("These are the two parameters", name, age)
	
}