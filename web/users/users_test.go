package users_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"ta/pkg/users"
	usersHandler "ta/web/users"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateUser(t *testing.T) {
	u := users.NewUser(1, "test_user_01", "test1", "test1")
	user, err := json.Marshal(&u)
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}

	request := httptest.NewRequest("POST", "http://localhost:8080/user/1", bytes.NewBuffer(user))
	request.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", usersHandler.CreateUserHandler).Methods("POST")
	router.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusCreated, response.Code)
	}
}

func TestGetUser(t *testing.T) {
	request, err := http.NewRequest("GET", "http://localhost:8080/user/1", nil)
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}

	response := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", usersHandler.GetUserHandler).Methods("GET")
	router.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestUpdateUser(t *testing.T) {
	u := users.NewUser(1, "Oleg2", "Djur2", "Golang2")

	arr, err := json.Marshal(&u)
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest("PUT", "/user/1", bytes.NewBuffer(arr))
	req.Header.Set("Content-Type", "application/json")

	r := mux.NewRouter()

	r.HandleFunc("/user/{id}", usersHandler.UpdateUserHandler).Methods("PUT")

	r.ServeHTTP(rr, req)

	user := &users.User{}
	if err = json.Unmarshal(rr.Body.Bytes(), &user); err != nil {
		t.Errorf("An error occurred. %v", err)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}

func TestDeleteProduct(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	r := mux.NewRouter()

	r.HandleFunc("/user/{id}", usersHandler.DeleteUserHandler).Methods("DELETE")

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d.\n Got %d", http.StatusOK, status)
	}
}
