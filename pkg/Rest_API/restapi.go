// Copyright (c) avijit bhattacharjee 2024

package Rest_API

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Start() {
	fmt.Println("Implementing simple rest API")

	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/{param1}/{param2}", paramsModify)

	// Create a server with timeouts
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	// Run the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe(): %v\n", err)
		}
	}()

	// Wait for a signal or shutdown trigger
	// Here, simulate a shutdown after a delay for demonstration purposes
	time.Sleep(10 * time.Second)

	// Shutdown the server gracefully
	fmt.Println("Shutting down the server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Printf("Server Shutdown Failed:%+v", err)
	}

	fmt.Println("Server exited properly")
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
