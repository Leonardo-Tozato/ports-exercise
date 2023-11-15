package portsrepo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"ports-exercise/m/internal/port/domain"
	"testing"
)

// Unit tests for repository upsert
func TestMemDB_Insert(t *testing.T) {
	jsonData := `{"ZWUTA":{"name":"Mutare","city":"Mutare","country":"Zimbabwe","alias":[],"regions":[],"coordinates":[32.650351,-18.9757714],"province":"Manicaland","timezone":"Africa/Harare","unlocs":["ZWUTA"]}}`

	portData := make(domain.PortData)

	err := json.Unmarshal([]byte(jsonData), &portData)
	if err != nil {
		assert.Error(t, err)
		return
	}

	repository := NewMemDB()
	assert.Equal(t, repository.Upsert(portData), portData)
}

func TestMemDB_Update(t *testing.T) {
	jsonData := `{"ZWUTA":{"name":"Mutare","city":"Mutare","country":"Zimbabwe","alias":[],"regions":[],"coordinates":[32.650351,-18.9757714],"province":"Manicaland","timezone":"Africa/Harare","unlocs":["ZWUTA"]}}`

	portData := make(domain.PortData)

	err := json.Unmarshal([]byte(jsonData), &portData)
	if err != nil {
		assert.Error(t, err)
		return
	}

	repository := NewMemDB()
	repository.Upsert(portData)

	if entry, ok := portData["ZWUTA"]; ok {
		entry.Code = "TEST UPDATE"
		portData["ZWUTA"] = entry
	}
	result := repository.Upsert(portData)

	assert.Equal(t, result["ZWUTA"].Code, "TEST UPDATE")
}
