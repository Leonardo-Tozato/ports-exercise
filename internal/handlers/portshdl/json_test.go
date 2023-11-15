package portshdl

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"ports-exercise/m/internal/port"
	"ports-exercise/m/internal/port/domain"
	"ports-exercise/m/internal/repositories/portsrepo"
	"testing"
)

// Integration tests covering all the ports upsert through JSON files
func TestJsonHandler_HandleUpsertStream(t *testing.T) {
	tests := []struct {
		name   string
		source string
	}{
		{
			name:   "should successfully process all data from ports.json",
			source: "ports.json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := fmt.Sprintf("./testdata/%s", tt.source)
			file, err := os.Open(filePath)
			if err != nil {
				t.Error(err)
				return
			}

			content, err := io.ReadAll(file)
			if err != nil {
				t.Error(err)
				return
			}

			var expectedOutput domain.PortData
			if err := json.Unmarshal(content, &expectedOutput); err != nil {
				t.Errorf("failed to read testdata file: %s - %v", tt.source, err)
				return
			}

			portsRepository := portsrepo.NewMemDB()
			portsService := port.NewService(portsRepository)
			portsHandler := NewJsonHandler(portsService)

			result, err := portsHandler.HandleUpsertStream(filePath)
			if err != nil {
				t.Error(err)
				return
			}

			assert.Equal(t, expectedOutput, result)
		})
	}
}
