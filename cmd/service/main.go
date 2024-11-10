package main

import (
	"file-structure/internal/api"
	"file-structure/internal/repository"
	"file-structure/internal/service"
	"fmt"
)

// компилируемые (go, c, c++, java)
// интерпритируемые (python, js)

func main() {
	repo := repository.New()

	srv := service.New(repo)

	handler := api.New(srv)

	res := handler.Handle(10, 2)
	fmt.Println(res)
}
