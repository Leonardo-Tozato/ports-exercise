package port

import "ports-exercise/m/internal/port/domain"

type Repository interface {
	Upsert(port domain.PortData)
	FindAll() domain.PortData
}
