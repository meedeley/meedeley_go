package pkg

import (
	"log"
	"os"
	"path/filepath"
)

func BasePath(path string) string {
	rootPath, err := os.Getwd()

	if err != nil {
		log.Printf("Cannot find file %v", err)
	}

	return filepath.Join(rootPath, path)
}

func SetupLogger() (*os.File, error) {

	path := os.Getenv("LOG_PATH")
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	return file, nil
}
