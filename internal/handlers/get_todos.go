package handlers

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"github.com/kalelmartinho/goapi/api"
	"github.com/kalelmartinho/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var params = api.TodoParams{}
	var decoder *schema.Decoder = schema.NewDecoder()

	var err error
	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.TodosDetails
	tokenDetails = (*database).GetTodos(params.Username)

	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var (
		response = api.TodoResponse{
			TodosDetails: *tokenDetails,
			Code:         http.StatusOK,
		}
	)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}
