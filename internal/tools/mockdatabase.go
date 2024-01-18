package tools

import "time"

type mockDatabase struct {
}

var (
	mockLoginDetails = map[string]LoginDetails{
		"Kalel": {
			AuthToken: "1234",
			Username:  "Kalel",
		},
		"Leonardo": {
			AuthToken: "5678",
			Username:  "Leonardo",
		},
		"Martinho": {
			AuthToken: "9012",
			Username:  "Martinho",
		},
	}
	mockTodoDetails = map[string]TodosDetails{
		"Kalel": {
			Todos: []TodoDetails{
				{
					Title:       "Todo 1",
					Description: "Description 1",
					Username:    "Kalel",
				},
				{
					Title:       "Todo 2",
					Description: "Description 2",
					Username:    "Kalel",
				},
			},
		},
		"Leonardo": {
			Todos: []TodoDetails{
				{
					Title:       "Todo 1",
					Description: "Description 1",
					Username:    "Leonardo",
				},
				{
					Title:       "Todo 2",
					Description: "Description 2",
					Username:    "Leonardo",
				},
			},
		},
		"Martinho": {
			Todos: []TodoDetails{
				{
					Title:       "Todo 1",
					Description: "Description 1",
					Username:    "Martinho",
				},
				{
					Title:       "Todo 2",
					Description: "Description 2",
					Username:    "Martinho",
				},
			},
		},
	}
)

func (db *mockDatabase) GetLoginDetails(username string) *LoginDetails {
	time.Sleep(1 * time.Second)
	var userData = LoginDetails{}
	userData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &userData
}

func (db *mockDatabase) GetTodos(username string) *TodosDetails {
	time.Sleep(1 * time.Second)
	var todos TodosDetails
	todos, ok := mockTodoDetails[username]
	if !ok {
		return nil
	}
	return &todos
}

func (db *mockDatabase) SetupDatabase() error {
	return nil
}
