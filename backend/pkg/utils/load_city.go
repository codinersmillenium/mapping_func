package utils

import (
	"encoding/json"
	"os"
)

func LoadCities(path string) (map[string]string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data map[string]string
	err = json.Unmarshal(f, &data)
	return data, err
}
