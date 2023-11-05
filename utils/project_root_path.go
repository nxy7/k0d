package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetProjectRootPath(path string) string {
	path, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	for _, d := range entries {
		if d.Name() == ".git" {
			return path
		}
	}

	pathSplit := strings.Split(path, "/")
	newPath := strings.Join(pathSplit[0:len(pathSplit)-1], "/")
	if len(newPath) == 0 {
		panic(fmt.Errorf("Working directory is not a repository and there's no git repo in any parent folder"))
	}

	return GetProjectRootPath(newPath)
}
