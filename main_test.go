package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuccessHealthz(t *testing.T) {
	req, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleHealthz)

	handler.ServeHTTP(responseRecorder, req)

	h := &HealthZResponse{}
	json.Unmarshal(responseRecorder.Body.Bytes(), h)
	expected := &HealthZResponse{"OK"}

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("HandleHealthz returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if *h != *expected {
		t.Errorf("HandleHealthz returned unexpected body: got %v want %v",
			*h, *expected)
	}
}

func TestMethodNotAllowedHealthz(t *testing.T) {
	req, err := http.NewRequest("POST", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleHealthz)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("HandleHealthz returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestMethodNotAllowedUsers(t *testing.T) {
	req, err := http.NewRequest("PATCH", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUsers)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("HandleUsers returned wrong status code: got %v want %v",
			status, http.StatusMethodNotAllowed)
	}
}

func TestManyPathParamsUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUsers)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusNotFound {
		t.Errorf("HandleUsers returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestSucessGetUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUsers)

	handler.ServeHTTP(responseRecorder, req)

	h := &[]User{}
	json.Unmarshal(responseRecorder.Body.Bytes(), h)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("HandleUsers returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if len(*h) == 0 {
		t.Errorf("HandleUsers returned empty body %v", *h)
	}
}

func TestSucessGetUsersWithQueryParams(t *testing.T) {
	req, err := http.NewRequest("GET", "/users?firstName=Lucas", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUsers)

	handler.ServeHTTP(responseRecorder, req)

	h := &[]User{}
	json.Unmarshal(responseRecorder.Body.Bytes(), h)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("HandleUsers returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if len(*h) == 0 {
		t.Errorf("HandleUsers returned empty body %v", *h)
	}
}

func TestSucessGetUserById(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUsers)

	handler.ServeHTTP(responseRecorder, req)

	h := &User{}
	json.Unmarshal(responseRecorder.Body.Bytes(), h)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("HandleUsers returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if *h == (User{}) {
		t.Errorf("HandleUsers returned empty body %v", *h)
	}
}

func TestSucessInsertUser(t *testing.T) {

	user, _ := json.Marshal(User{"456", "Lucas", "Machado", 40})

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(user))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUsers)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusCreated {
		t.Errorf("HandleUsers returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestErrorInsertUser(t *testing.T) {

	user, _ := json.Marshal(User{"456", "Lucas", "Machado", 40})

	req, err := http.NewRequest("POST", "/users/222", bytes.NewBuffer(user))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUsers)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusNotFound {
		t.Errorf("HandleUsers returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestSucessUpdateUser(t *testing.T) {

	user, _ := json.Marshal(User{"123", "Lucas", "Machado", 41})

	req, err := http.NewRequest("PUT", "/users/123", bytes.NewBuffer(user))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUsers)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusNoContent {
		t.Errorf("HandleUsers returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}

func TestErrorUpdateUser(t *testing.T) {

	user, _ := json.Marshal(User{"123", "Lucas", "Machado", 41})

	req, err := http.NewRequest("PUT", "/users", bytes.NewBuffer(user))
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUsers)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("HandleUsers returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestSucessDeleteUser(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/users/123", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUsers)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusNoContent {
		t.Errorf("HandleUsers returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}

func TestErrorDeleteUser(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleUsers)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("HandleUsers returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
