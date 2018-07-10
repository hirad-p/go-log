package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/jb-hirad/go-log/model"
)

// Read parses a json file and returns the levels configuration for the logger
func Read(path string) []model.Level {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var levels model.Levels
	json.Unmarshal(byteValue, &levels)

	return levels.Levels
}
