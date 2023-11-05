package utils

import "os"

// Saves string to file (in /tmp folder) and returns path to said file
func SaveToFile(dataString, name string) (string, error) {
	path := "/tmp/" + name
	err := os.WriteFile(path, []byte(dataString), 0777)
	if err != nil {
		return "", err
	}
	return path, err
}
