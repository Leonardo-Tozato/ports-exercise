package portsrepo

import (
	"maps"
	"ports-exercise/m/internal/ports"
	"ports-exercise/m/internal/ports/domain"
)

// In memory key value database
type memDB struct {
	ports domain.PortsData
}

func NewMemDB() ports.Repository {
	return &memDB{ports: make(domain.PortsData)}
}

// Upsert the database ports collection
func (repo *memDB) Upsert(ports domain.PortsData) {
	maps.Copy(repo.ports, ports)
}
