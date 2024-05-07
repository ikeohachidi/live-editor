package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

type Session struct {
	Lang    string `json:"lang"`
	Content string `json:"content"`
}

type SessionContent struct {
	CSS  string `json:"css"`
	JS   string `json:"js"`
	HTML string `json:"html"`
}

func StartSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	sessionId := uuid.New().String()

	folderPath := fmt.Sprintf("./session_files/%v", sessionId)

	err := os.MkdirAll(folderPath, 0777)
	if err != nil {
		log.Errorf("error creating session folder: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	writeMetadata(folderPath, Metadata{LastUpdated: time.Now()})

	file, err := os.Create(fmt.Sprintf("%v/%v.html", folderPath, sessionId))
	if err != nil {
		log.Errorf("error creating initial html: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var newSession struct {
		SessionId string `json:"sessionId"`
	}
	newSession.SessionId = sessionId

	json.NewEncoder(w).Encode(newSession)
}

func processSass(session Session, sessionId string, folderPath string) error {
	filePath := fmt.Sprintf("%v/%v.css", folderPath, sessionId)

	cmd := exec.Command("sass", "--stdin", "--no-source-map", "--no-error-css", filePath)

	cmd.Stdin = strings.NewReader(session.Content)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error running sass build command: %v", err)
	}

	writeMetadata(folderPath, Metadata{LastUpdated: time.Now()})

	return nil
}

func processTypescript(session Session, sessionId string, folderPath string) error {
	filePath := fmt.Sprintf("%v/%v.js", folderPath, sessionId)
	tsFilePath := fmt.Sprintf("%v/%v.ts", folderPath, sessionId)

	// couldn't find a way to get typescript from stdin
	// so we'll create the file, compile it, then delete it
	err := os.WriteFile(tsFilePath, []byte(session.Content), 0777)
	if err != nil {
		return fmt.Errorf("error creating ts file: %v", err)
	}

	cmd := exec.Command("tsc", tsFilePath, "--outfile", filePath)
	err = cmd.Run()
	if err != nil {
		os.Remove(tsFilePath)
		return fmt.Errorf("error compiling ts file: %v", err)
	}
	os.Remove(tsFilePath)

	writeMetadata(folderPath, Metadata{LastUpdated: time.Now()})

	return nil
}

func UpdateSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	sessionId := r.PathValue("id")
	if sessionId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var body Session

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Errorf("error reading session update request body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	folderPath := fmt.Sprintf("./session_files/%v", sessionId)

	if body.Lang == "ts" {
		err = processTypescript(body, sessionId, folderPath)
		if err != nil {
			log.Error(err)
		}
		return
	}

	if body.Lang == "scss" || body.Lang == "sass" {
		err = processSass(body, sessionId, folderPath)
		if err != nil {
			log.Error(err)
		}
		return
	}

	filePath := fmt.Sprintf("%v/%v.%v", folderPath, sessionId, body.Lang)

	err = os.WriteFile(filePath, []byte(body.Content), 0777)
	if err != nil {
		log.Errorf("error writing file: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

const htmlTemplate = `
<!DOCTYPE html>
<html>
    <head>
		<link href="%[1]v/f/%[2]v/%[2]v.css" rel="stylesheet"></link>
	</head>
    <body>
		%[3]v
		<script src="%[1]v/f/%[2]v/%[2]v.js"></script>
    </body>
</html>
`

func FetchSession(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	sessionId := r.PathValue("id")
	if sessionId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	filePath := fmt.Sprintf("./session_files/%[1]v/%[1]v.html", sessionId)

	file, err := os.OpenFile(filePath, os.O_RDONLY, 0777)
	if err != nil {
		log.Errorf("error reading file from storage: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Errorf("error reading file content: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	page := fmt.Sprintf(htmlTemplate, os.Getenv("API"), sessionId, string(fileContent))

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(page))
}

func FetchSessionFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	sessionId := r.PathValue("sessionId")
	if sessionId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// find all files in that session
	folderPath := fmt.Sprintf("./session_files/%v", sessionId)
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		log.Errorf("error reading session directory: %v", err)
		w.WriteHeader(http.StatusGone)
		return
	}

	var sessionContent SessionContent

	for _, entry := range entries {
		filePath := fmt.Sprintf("./session_files/%v/%v", sessionId, entry.Name())
		extension := filepath.Ext(entry.Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			log.Errorf("error reading session file content: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if extension == ".css" {
			sessionContent.CSS = string(content)
		}
		if extension == ".js" {
			sessionContent.JS = string(content)
		}
		if extension == ".html" {
			sessionContent.HTML = string(content)
		}
	}

	c, err := json.Marshal(sessionContent)
	if err != nil {
		log.Errorf("error marshalling json content: %v", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Write(c)
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	sessionId := r.PathValue("sessionId")
	if sessionId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fileParam := r.PathValue("file")
	if fileParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	filePath := fmt.Sprintf("./session_files/%v/%v", sessionId, fileParam)
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0777)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Errorf("error reading file content: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	extension := filepath.Ext(filePath)
	if extension == ".js" {
		w.Header().Add("content-type", "application/javascript")
	}

	if extension == ".css" {
		w.Header().Add("content-type", "text/css; charset=utf-8")
	}

	w.Write(fileContent)
}

func HandleOptionsRequest(w http.ResponseWriter, r *http.Request) {
	accessControlRequest := r.Header.Get("access-control-request-method")
	accessControlRequestHeaders := r.Header.Get("access-control-request-headers")
	if accessControlRequestHeaders != "" {
		w.Header().Set("access-control-allow-headers", accessControlRequestHeaders)
	}

	w.Header().Set("access-control-allow-methods", accessControlRequest)
	w.Header().Set("access-control-allow-origin", "*")
	w.WriteHeader(http.StatusNoContent)
}
