package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/krishicks/concourse-current-time-resource/models"
)

func main() {
	var request models.InRequest
	err := json.NewDecoder(os.Stdin).Decode(&request)
	if err != nil {
		log.Fatalf("reading request: %s", err)
	}

	if len(os.Args) < 2 {
		log.Fatalf("usage: %s <destDir>", os.Args[0])
	}

	destDir := os.Args[1]
	err = os.MkdirAll(destDir, 0755)
	if err != nil {
		log.Fatalf("creating %s: %s", destDir, err)
	}

	destFile := filepath.Join(destDir, "time")
	err = ioutil.WriteFile(destFile, []byte(request.Version.Time), 0755)
	if err != nil {
		log.Fatalf("creating %s: %s", destFile, err)
	}

	response := models.InResponse{
		Version: request.Version,
	}

	err = json.NewEncoder(os.Stdout).Encode(response)
	if err != nil {
		log.Fatalf("encoding response: %s", err)
	}
}
