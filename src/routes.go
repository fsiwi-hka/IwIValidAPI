package src

import (
	"bufio"
	"encoding/json"
	"net/http"
	"os"
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
		OpenFile, err := os.Open("users")
		if err != nil {
			CustomInternalError(w, "Error when Reading Current Valid Users")
		}
		defer OpenFile.Close()
		scanner := bufio.NewScanner(OpenFile)

		var users []string
		i := 0
		for scanner.Scan() {
			line := scanner.Text()
			users = append(users, line)
			i++
		}

		if err := scanner.Err(); err != nil {
			CustomInternalError(w, "Error when Reading Current Valid Users")
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
