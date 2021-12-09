package model

type TodoRequest struct {
	Todo string `json:"todo"`
}

type TodoResponse struct {
	Id   int    `json:"id"`
	Todo string `json:"todo"`
}

type TodoUpdateRequest struct {
	Id   int    `json:"id"`
	Todo string `json:"todo"`
}
