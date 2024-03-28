package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Metadata struct {
	LastUpdated time.Time `json:"lastUpdated"`
}

func writeMetadata(path string, data Metadata) error {
	var metadata bytes.Buffer

	err := json.NewEncoder(&metadata).Encode(&data)
	if err != nil {
		return err
	}

	metadataFilePath := fmt.Sprintf("%v/metadata.json", path)
	err = os.WriteFile(metadataFilePath, metadata.Bytes(), 0777)

	return err
}
