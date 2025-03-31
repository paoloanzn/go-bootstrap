package parsing

import (
	"fmt"
	"log"
	"encoding/json"
	"os"
	
)

type JSONTemplate struct {
	Project interface{} `json:"project"`
	Config map[string]interface{} `json:"config"`
}

func ParseTemplate(filePath string) (*JSONTemplate, error) {
	var pJsonTemplate *JSONTemplate = &JSONTemplate{}

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v\n", err)
	}

	err = json.Unmarshal(data, pJsonTemplate)
	if err != nil {
		return pJsonTemplate, fmt.Errorf("Unable to parse json file at %s: %v\n", filePath, err)
	}

	return pJsonTemplate, nil
}