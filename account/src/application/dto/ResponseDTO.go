package dto

type ResponseAuthDTO struct {
	ResponseRequestDTO
	Token string
}

type ResponseRequestDTO struct {
	Mensagem string
	Ok       bool
	ID       int64
	Token    string
}

type ApiResponse struct {
	StatusCode int
	Message    string
}
