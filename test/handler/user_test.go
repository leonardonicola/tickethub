package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/leonardonicola/tickethub/internal/dto"
	"github.com/leonardonicola/tickethub/internal/router"
)

func TestUserHandler(t *testing.T) {
	router, err := router.InitRoutes()

	if err != nil {
		t.Fatalf("%v", err)
	}
	t.Run("register new user", func(t *testing.T) {
		w := httptest.NewRecorder()
		user := dto.CreateUserInputDTO{
			Name:    "Leonardo",
			Email:   "email@hotmail.com",
			CPF:     "12345678910",
			Surname: "Leonardo",
			Address: "Av Paulista",
		}
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(&user); err != nil {
			t.Fatal("Error while encoding body!")
		}
		req, _ := http.NewRequest("POST", "/api/v1/register", &buf)
		router.ServeHTTP(w, req)

		var createdUser dto.CreateUserOutputDTO
		err := json.NewDecoder(w.Body).Decode(&createdUser)
		if err != nil {
			t.Errorf("Error while decoding response body: %v", err)
		}

		if reflect.DeepEqual(createdUser, user) {
			t.Errorf("WANTED: %s\n HAD: %s\n", createdUser.Email, user.Email)
		}
	})
}
