package service

import (
	"file-structure/internal/model"
)

type Service struct {
	repo Repo
}

func New(rep Repo) *Service {
	return &Service{
		repo: rep,
	}
}

func (s *Service) Summator(data *model.Data) {
	data.Sum = data.A + data.B

	s.repo.Save(data.Sum)
}
