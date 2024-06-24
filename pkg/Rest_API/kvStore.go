// Copyright (c) avijit bhattacharjee 2024

package main

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// KeyValueStore is a simple in-memory key-value store
type KeyValueStore struct {
	store map[string]string
	mu    sync.RWMutex
}

// NewKeyValueStore creates a new key-value store
func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		store: make(map[string]string),
	}
}

// Set sets the value for a key in the store
func (kvs *KeyValueStore) Set(key, value string) {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()
	kvs.store[key] = value
}

// Get gets the value for a key from the store
func (kvs *KeyValueStore) Get(key string) (string, bool) {
	kvs.mu.RLock()
	defer kvs.mu.RUnlock()
	value, exists := kvs.store[key]
	return value, exists
}

// Delete deletes a key from the store
func (kvs *KeyValueStore) Delete(key string) {
	kvs.mu.Lock()
	defer kvs.mu.Unlock()
	delete(kvs.store, key)
}

// Handler functions
func (kvs *KeyValueStore) SetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	var request struct {
		Value string json:"value"
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	kvs.Set(key, request.Value)
	w.WriteHeader(http.StatusOK)
}

func (kvs *KeyValueStore) GetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	value, exists := kvs.Get(key)
	if !exists {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	response := map[string]string{"value": value}
	json.NewEncoder(w).Encode(response)
}

func (kvs *KeyValueStore) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	kvs.Delete(key)
	w.WriteHeader(http.StatusOK)
}

func main() {
	kvs := NewKeyValueStore()
	router := mux.NewRouter()

	// Define the routes
	router.HandleFunc("/kv/{key}", kvs.SetHandler).Methods("POST")
	router.HandleFunc("/kv/{key}", kvs.GetHandler).Methods("GET")
	router.HandleFunc("/kv/{key}", kvs.DeleteHandler).Methods("DELETE")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}