package ports

import "ports-exercise/m/internal/ports/domain"

type Service interface {
	Upsert(ports domain.PortsData)
}

type service struct {
	repository Repository
}

func New(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Upsert(ports domain.PortsData) {
	s.repository.Upsert(ports)
}
