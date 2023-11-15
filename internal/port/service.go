package port

import "ports-exercise/m/internal/port/domain"

type Service interface {
	Upsert(ports domain.PortData)
	FindAll() domain.PortData
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Upsert(ports domain.PortData) {
	s.repository.Upsert(ports)
}

func (s *service) FindAll() domain.PortData {
	return s.repository.FindAll()
}
