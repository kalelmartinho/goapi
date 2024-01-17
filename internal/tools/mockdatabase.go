package tools

import "time"

type mockDatabase struct {
}

var mockLoginDetails = map[string]LoginDetails{
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

var mockTodoDetails = map[string][]TodoDetails{
	"Kalel": {
		{
			Title:       "Kalel's Morning Routine",
			Description: "Wake up, eat breakfast, go to work",
		},
		{
			Title:       "Kalel's Evening Routine",
			Description: "Go home, eat dinner, go to sleep",
		},
	},
	"Leonardo": {
		{
			Title:       "Leonardo's Morning Routine",
			Description: "Wake up, eat breakfast, go to school",
		},
		{
			Title:       "Leonardo's Evening Routine",
			Description: "Go home, eat dinner, go to sleep",
		},
	},
	"Martinho": {
		{
			Title:       "Martinho's Morning Routine",
			Description: "Wake up, eat breakfast, go to gym",
		},
		{
			Title:       "Martinho's Evening Routine",
			Description: "Go home, eat dinner, go to sleep",
		},
	},
}

func (db *mockDatabase) GetLoginDetails(username string) *LoginDetails {
	time.Sleep(1 * time.Second)
	var userData = LoginDetails{}
	userData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &userData
}

func (db *mockDatabase) GetTodos(username string) *[]TodoDetails {
	time.Sleep(1 * time.Second)
	var todos []TodoDetails
	todos, ok := mockTodoDetails[username]
	if !ok {
		return nil
	}
	return &todos
}

func (db *mockDatabase) SetupDatabase() error {
	return nil
}
