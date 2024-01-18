package middleware

import (
	"errors"
	"github.com/kalelmartinho/goapi/api"
	"github.com/kalelmartinho/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var UnauthorizedError = errors.New("Invalid token or user")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")

		if token == "" || username == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err := tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetLoginDetails(username)

		if loginDetails == nil || token != (*loginDetails).AuthToken {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		next.ServeHTTP(w, r)

	})
}
