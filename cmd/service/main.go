package main

import (
	"file-structure/internal/api"
	"file-structure/internal/repository"
	"file-structure/internal/service"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// компилируемые (go, c, c++, java)
// интерпритируемые (python, js)

func main() {
	repo := repository.New()

	srv := service.New(repo)

	handler := api.New(srv)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello GET!!"))
		case http.MethodPost:
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Hello POST!!"))
		}
	})

	http.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			_ = r.PathValue("id")

			a := r.URL.Query().Get("a")
			b := r.URL.Query().Get("b")
			log.Println(a, b)
			aInt, err := strconv.Atoi(a)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			bInt, err := strconv.Atoi(b)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			res := handler.Handle(aInt, bInt)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(fmt.Sprintf("%d", res)))
		}
	})

	//fmt.Println(res)

	err := http.ListenAndServe(":3232", nil)
	if err != nil {
		panic(err)
	}
}
