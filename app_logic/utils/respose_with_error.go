package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Функция для отправки пользователю данных об ошибках

func ResponseWithError(w http.ResponseWriter, r *http.Request, e error, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	response, _ := json.Marshal(map[string]string{"error": fmt.Sprintf("%v %v: %v", statusCode, http.StatusText(statusCode), e)})
	http.Error(
		w,
		string(response),
		statusCode,
	)
	log.Printf("\033[31;1mError in request %v %v: %v\033[0m", r.Method, r.RequestURI, e)
}
