package port

import "ports-exercise/m/internal/port/domain"

type Repository interface {
	Upsert(portData domain.PortData) domain.PortData
}
