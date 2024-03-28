package services

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/ikeohachidi/live-editor/routes"
)

// find all stale folders and delete them
func PruneStaleSessions(d time.Duration) {
	ticker := time.NewTicker(d)
	defer ticker.Stop()

	for range ticker.C {
		findAndClean(d)
	}
}

func findAndClean(d time.Duration) error {
	folders, err := os.ReadDir("./session_files")
	if err != nil {
		return err
	}

	// this assumes the folder structure in session_files
	// is session_files/<session>/
	for _, folder := range folders {

		// find metadata.json file in folder
		metadataContent, err := os.ReadFile(fmt.Sprintf("./session_files/%v/metadata.json", folder.Name()))
		if err != nil {
			return err
		}

		canDelete, err := isStale(metadataContent, d)
		if err != nil {
			return err
		}

		if canDelete {
			os.RemoveAll(fmt.Sprintf("./session_files/%v", folder.Name()))
		}
	}

	return nil
}

func isStale(data []byte, d time.Duration) (bool, error) {
	var metadata routes.Metadata
	err := json.Unmarshal(data, &metadata)
	if err != nil {
		return false, err
	}

	now := time.Now()
	elapsed := now.Sub(metadata.LastUpdated)
	if elapsed < d {
		return false, nil
	}

	return true, nil
}
