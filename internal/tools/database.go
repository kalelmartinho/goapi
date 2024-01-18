package tools

import log "github.com/sirupsen/logrus"

type LoginDetails struct {
	Username  string
	AuthToken string
}

type TodoDetails struct {
	Title       string
	Description string
	Username    string
}

type TodosDetails struct {
	Todos []TodoDetails
}

type DatabaseInterface interface {
	GetLoginDetails(username string) *LoginDetails
	GetTodos(username string) *TodosDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDatabase{}
	var err error = database.SetupDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
