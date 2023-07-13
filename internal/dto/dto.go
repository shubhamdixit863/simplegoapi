package dto

type AddRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserResponseDto struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
