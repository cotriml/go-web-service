package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

const address = "0.0.0.0"
const port = "8080"
const readTimeout = 10 * time.Second
const writeTimeout = 10 * time.Second
const maxHeaderBytes = 1 << 20

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", HandleHealthz)
	mux.HandleFunc("/users/", HandleUsers)
	mux.HandleFunc("/users", HandleUsers)
	s := &http.Server{
		Addr:           address + ":" + port,
		Handler:        mux,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes, // 1MB
	}
	println("Server Running at", address, "Port", port)
	log.Fatal(s.ListenAndServe())
}

type HealthZResponse struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

func HandleHealthz(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		res := &HealthZResponse{Status: "OK"}
		b, _ := json.Marshal(res)
		w.Write(b)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	const userPath = "/users"

	id := strings.TrimPrefix(r.URL.Path, userPath+"/")

	if len(strings.Split(id, "/")) > 1 && id != userPath {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 page not found"))
		return
	}
	queryParams := make(map[string]string)

	params := r.URL.Query()
	for i, p := range params {
		for _, v := range p {
			if len(v) > 0 {
				queryParams[i] = v
				println(i, "=>", v)
			}
		}
	}

	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")

		if len(id) > 0 && id != userPath {
			user, _ := getUserById(id)
			w.WriteHeader(http.StatusOK)
			w.Write(user)
			return
		}

		users, _ := getAllUsers()
		w.WriteHeader(http.StatusOK)
		w.Write(users)

	case http.MethodPost:
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json")

		if len(id) > 0 && id != userPath {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		user := &User{}
		json.NewDecoder(r.Body).Decode(user)

		InsertUser(user)
		w.WriteHeader(http.StatusCreated)

	case http.MethodPut:
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json")

		user := &User{}
		json.NewDecoder(r.Body).Decode(user)

		if len(id) > 0 && id != userPath {
			UpdateUser(id, user)
			w.WriteHeader(http.StatusNoContent)
		} else {

			res, _ := json.Marshal(ErrorResponse{"Missing User ID from Path"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)
			return
		}

	case http.MethodDelete:
		w.Header().Set("Content-Type", "application/json")

		if len(id) > 0 && id != userPath {
			DeleteUser(id)
			w.WriteHeader(http.StatusNoContent)
		} else {
			res, _ := json.Marshal(ErrorResponse{"Missing User ID from Path"})
			w.WriteHeader(http.StatusBadRequest)
			w.Write(res)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getAllUsers() ([]byte, error) {

	res := &[]User{
		{"123", "Lucas", "Cotrim", 28},
		{"456", "Lucas", "Machado", 40},
	}

	b, _ := json.Marshal(res)

	return b, nil
}

func getUserById(id string) ([]byte, error) {
	res := &User{"123", "Lucas", "Cotrim", 28}
	b, _ := json.Marshal(res)
	return b, nil
}

func InsertUser(user *User) error {
	_, err := json.Marshal(user)
	return err
}

func UpdateUser(id string, user *User) error {
	_, err := json.Marshal(user)
	return err
}

func DeleteUser(id string) error {
	return nil
}
