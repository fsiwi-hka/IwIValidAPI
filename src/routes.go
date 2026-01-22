package src

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/fsiwi-hka/IwIValidAPI/src/reader"
)

type BadResponse struct {
	Error string `json:"error"`
}

func CustomInternalError(w http.ResponseWriter, err string) {
	resBody := BadResponse{Error: err}
	jsonResBody, _ := json.Marshal(resBody)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(jsonResBody)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == os.Getenv("API_TOKEN") {
			next.ServeHTTP(w, r)

			return
		}
		CustomInternalError(w, "No valid Token in Authorization header")
	})
}

type GetValidUserResbody struct {
	Users []string `json:"users"`
}

func HandlerGetValidUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := reader.ReadUsersFromFile()
		if err != nil {
			CustomInternalError(w, "Error when Reading Current Valid Users")
			return
		}
		resBody := GetValidUserResbody{Users: users}
		jsonResBody, err := json.Marshal(resBody)
		if err != nil {
			CustomInternalError(w, "Marshal Error")
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResBody)
	}
}
