package api

import (
	"file-structure/internal/model"
)

type Handler struct {
	srv SrvI
}

func New(srvs SrvI) *Handler { //конструктор
	return &Handler{
		srv: srvs,
	}
}

func (h *Handler) Handle(a, b int) int {
	data := model.Data{A: a, B: b}
	h.srv.Summator(&data) //Handler->srv->Summator
	return data.Sum
}
