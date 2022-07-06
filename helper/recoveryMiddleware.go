package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				log.Printf("RecoveryMiddleware : ")
				fmt.Println(err)
				jsonData, jsonErr := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})
				if jsonErr != nil {
					log.Printf("RecoveryMiddleware : Error in json")
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonData)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
