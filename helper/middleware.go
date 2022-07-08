package helper

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"todo/models"
)

func Middleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		sessToken := r.Header.Get("SessionID")
		if sessToken == "" {
			log.Printf("MiddleWare : session Token is empty")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		isExpired, sessTokenErr := SessionExist(sessToken)
		if sessTokenErr != nil {
			if sessTokenErr == sql.ErrNoRows {
				log.Printf("middleware : no row exist")
				w.WriteHeader(http.StatusBadRequest)
				return
			}
		}

		if isExpired {
			w.WriteHeader(http.StatusUnauthorized)
		}

		sessTokenRef := RefreshSessToken(sessToken)
		if sessTokenRef != nil {
			log.Printf("middleware : error in refreshing the token")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		//todo multiple queries not needed
		userID, getIdErr := GetID(sessToken)
		if getIdErr != nil {
			log.Printf("MiddleWare : Error in retrieving the user id from session table")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//user, err := GetUserDetails(userID)
		//if err != nil {
		//	log.Printf("Middleware : Error in getting user detail")
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}

		ctxMap := models.ContextMap{CtxMap: map[string]string{
			"UserID": userID,
			"SessID": sessToken,
		}}
		c := context.Background()
		ctx := context.WithValue(c, "ID", ctxMap)
		//ctx = context.WithValue(r.Context(), "UserID", userID)

		next.ServeHTTP(w, r.WithContext(ctx))

	})

}
