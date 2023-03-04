package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/gorilla/mux"
)

const PORT int = 3000

type Customer struct {
	ID        uuid.UUID
	Name      string
	Role      string
	Email     string
	Phone     uint64
	Contacted bool
}

var customers = map[uuid.UUID]Customer{}

var initialCustomers = []Customer{
	{
		ID:        uuid.New(),
		Name:      "Masatoshi Tanida",
		Role:      "Free-tier Customer",
		Email:     "masatoshi.tanida@gmail.com",
		Phone:     5555550000,
		Contacted: false,
	},
	{
		ID:        uuid.New(),
		Name:      "Atsuko Tanida",
		Role:      "Basic Customer",
		Email:     "atsuko.tanida@gmail.com",
		Phone:     5555550001,
		Contacted: false,
	},
	{
		ID:        uuid.New(),
		Name:      "Kaito Nakamura",
		Role:      "Premium Customer",
		Email:     "atsuko.tanida@gmail.com",
		Phone:     5555550002,
		Contacted: true,
	},
	{
		ID:        uuid.New(),
		Name:      "Yuto Tanaka",
		Role:      "Premium Customer",
		Email:     "atsuko.tanida@gmail.com",
		Phone:     5555550003,
		Contacted: false,
	},
	{
		ID:        uuid.New(),
		Name:      "Ayumi Takahashi",
		Role:      "Premium Customer",
		Email:     "atsuko.tanida@gmail.com",
		Phone:     5555550004,
		Contacted: false,
	},
}

func initializeCustomers() {
	for _, customer := range initialCustomers {
		customers[customer.ID] = customer
	}
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get path parameter
	s := mux.Vars(r)["id"]
	id, err := uuid.Parse(s)

	// Check if path parameter is a valid UUID
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		errorMsg := map[string]string{"error": "Invalid UUID"}
		json.NewEncoder(w).Encode(errorMsg)
		return
	}

	if _, ok := customers[id]; ok {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers[id])
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(struct{}{})
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Read request body
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		errorMsg := map[string]string{"error": "Invalid request body"}
		json.NewEncoder(w).Encode(errorMsg)
		return
	}

	// Unmarshal request body into Customer struct
	var newCustomer Customer
	err2 := json.Unmarshal(requestBody, &newCustomer)
	if err2 != nil {
		fmt.Println(err2)
		w.WriteHeader(http.StatusBadRequest)
		errorMsg := map[string]string{"error": "Invalid Customer data"}
		json.NewEncoder(w).Encode(errorMsg)
		return
	}

	// Add new Customer
	w.WriteHeader(http.StatusCreated)
	newCustomer.ID = uuid.New()
	customers[newCustomer.ID] = newCustomer
	json.NewEncoder(w).Encode(newCustomer)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	s := mux.Vars(r)["id"]
	id, err := uuid.Parse(s)

	// Check if path parameter is a valid UUID
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		errorMsg := map[string]string{"error": "Invalid UUID"}
		json.NewEncoder(w).Encode(errorMsg)
		return
	}

	if _, ok := customers[id]; ok {

		// Read request body
		requestBody, err := ioutil.ReadAll(r.Body)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			errorMsg := map[string]string{"error": "Invalid request body"}
			json.NewEncoder(w).Encode(errorMsg)
			return
		}

		// Unmarshal request body into Customer struct
		var updatedCustomer Customer
		err2 := json.Unmarshal(requestBody, &updatedCustomer)

		if err2 != nil {
			fmt.Println(err2)
			w.WriteHeader(http.StatusBadRequest)
			errorMsg := map[string]string{"error": "Invalid Customer data"}
			json.NewEncoder(w).Encode(errorMsg)
			return
		}

		// Update Customer data
		w.WriteHeader(http.StatusOK)
		updatedCustomer.ID = id
		customers[id] = updatedCustomer
		json.NewEncoder(w).Encode(updatedCustomer)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(struct{}{})
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get path parameter
	s := mux.Vars(r)["id"]
	id, err := uuid.Parse(s)

	// Check if path parameter is a valid UUID
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMsg := map[string]string{"error": "Invalid UUID"}
		json.NewEncoder(w).Encode(errorMsg)
		return
	}

	if _, ok := customers[id]; ok {
		delete(customers, id)
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(struct{}{})
	}
}

func main() {
	// Seed the database
	initializeCustomers()

	// Set up route handlers
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	router.HandleFunc("/customers", addCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	// Set up static web page
	index := http.FileServer(http.Dir("html"))
	router.Handle("/", index)

	// Run server
	fmt.Printf("Starting server at port %d...\n", PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", PORT), router))
}
