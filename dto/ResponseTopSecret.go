package dto

type ResponseTopSecret struct {

	Position Position `json:"position"`
	Message string `json:"message"`
}

type Position struct{
	X float32 `json:"x"`
	Y float32 `json:"y"`
}
