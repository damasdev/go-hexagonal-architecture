package user

import "fmt"

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Hello(message []byte) error {
	fmt.Println("hello:", string(message))
	return nil
}

func (h *Handler) World(message []byte) error {
	fmt.Println("world:", string(message))
	return nil
}
