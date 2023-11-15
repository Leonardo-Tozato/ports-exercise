package portshdl

import (
	"encoding/json"
	"fmt"
	"os"
	"ports-exercise/m/internal/port"
	"ports-exercise/m/internal/port/domain"
)

type JsonHandler struct {
	service port.Service
}

func NewJsonHandler(service port.Service) *JsonHandler {
	return &JsonHandler{
		service: service,
	}
}

// HandleUpsertStream HandleUpsert process the JSON file line by line
func (h *JsonHandler) HandleUpsertStream(filePath string) (domain.PortData, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v\n", err)

	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	decoder := json.NewDecoder(file)

	// Read the opening brace {
	token, err := decoder.Token()
	fmt.Println(token)

	if err != nil {
		return nil, fmt.Errorf("Error reading JSON opening brace: %v\n", err)
	}

	var result domain.PortData
	for decoder.More() {
		keyToken, err := decoder.Token()
		if err != nil {
			return nil, fmt.Errorf("Error reading port key: %v\n", err)
		}
		key := keyToken.(string)

		var value domain.Port
		if err := decoder.Decode(&value); err != nil {
			return nil, fmt.Errorf("Error decoding port value: %v\n", err)
		}

		portData := domain.PortData{
			key: value,
		}

		result = h.service.Upsert(portData)
	}

	return result, nil
}
