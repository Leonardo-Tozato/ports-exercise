package ports

import "ports-exercise/m/internal/ports/domain"

type Repository interface {
	Upsert(data domain.PortsData)
}
