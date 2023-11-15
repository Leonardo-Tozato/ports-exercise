package portsrepo

import (
	"maps"
	"ports-exercise/m/internal/port"
	"ports-exercise/m/internal/port/domain"
)

// In memory key value database
type memDB struct {
	ports domain.PortData
}

func NewMemDB() port.Repository {
	return &memDB{ports: make(domain.PortData)}
}

// Upsert the database ports collection
func (repo *memDB) Upsert(ports domain.PortData) {
	maps.Copy(repo.ports, ports)
}

func (repo *memDB) FindAll() domain.PortData {
	return repo.ports
}
