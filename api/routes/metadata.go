package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Metadata struct {
	LastUpdated    time.Time `json:"lastUpdated"`
	StyleLanguage  string    `json:"styleLanguage"`
	ScriptLanguage string    `json:"scriptLanguage"`
}

func UpdatedMetaData(sessionId string) (Metadata, error) {
	var metadata Metadata

	filePath := fmt.Sprintf("./session_files/%s/metadata.json", sessionId)
	file, err := os.OpenFile(filePath, os.O_RDWR, 0777)
	if err != nil {
		return metadata, err
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return metadata, err
	}

	err = json.Unmarshal(fileContent, &metadata)
	if err != nil {
		return metadata, err
	}

	return metadata, nil
}

func (m *Metadata) SetLastUpdated(value time.Time) {
	m.LastUpdated = value
}

func (m *Metadata) SetStyleLanguage(value string) {
	m.StyleLanguage = value
}

func (m *Metadata) SetScriptLanguage(value string) {
	m.ScriptLanguage = value
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
