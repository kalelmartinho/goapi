package api

type TodoParams struct {
	Username string
}

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoResponse struct {
	Todos []Todo
}

type Error struct {
	Code    int
	Message string
}
