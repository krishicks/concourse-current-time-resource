package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/krishicks/concourse-current-time-resource/models"
)

func main() {
	var request models.CheckRequest
	err := json.NewDecoder(os.Stdin).Decode(&request)
	if err != nil {
		log.Fatalf("parse error: %s", err)
	}

	versions := []models.Version{}

	previousTime := request.Version.Time
	if previousTime != "" {
		versions = append(versions, models.Version{Time: previousTime})
	}

	currentTime := time.Now().UTC()

	var timeStr string
	if request.Source.Format != "" {
		timeStr = currentTime.Format(request.Source.Format)
	} else {
		timeStr = currentTime.String()
	}

	versions = append(versions, models.Version{Time: timeStr})

	err = json.NewEncoder(os.Stdout).Encode(versions)
	if err != nil {
		log.Fatalf("encode error: %s", err)
	}
}
