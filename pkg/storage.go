package pkg

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func StoreImage(path string) error {
	allowedExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
		".webp": true,
	}

	ext := strings.ToLower(filepath.Ext(path))
	if !allowedExtensions[ext] {
		extensions := make([]string, 0, len(allowedExtensions))
		for k := range allowedExtensions {
			extensions = append(extensions, k)
		}
		return fmt.Errorf("file extension not allowed. Allowed extensions: %v", extensions)
	}

	srcFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	destDir := "/storage/urlPathnya"
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	destPath := filepath.Join(destDir, filepath.Base(path))
	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destFile.Close()

	buffer := make([]byte, 1024*1024)
	if _, err := io.CopyBuffer(destFile, srcFile, buffer); err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	return nil
}

func StoreFile(path string) error {
	allowedExtensions := map[string]bool{
		".pdf":  true,
		".zip":  true,
		".mp4":  true,
		".mov":  true,
		".avi":  true,
		".doc":  true,
		".docx": true,
		".xls":  true,
		".xlsx": true,
	}

	ext := strings.ToLower(filepath.Ext(path))
	if !allowedExtensions[ext] {
		extensions := make([]string, 0, len(allowedExtensions))
		for k := range allowedExtensions {
			extensions = append(extensions, k)
		}
		return fmt.Errorf("file extension not allowed. Allowed extensions: %v", extensions)
	}

	srcFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	destDir := "/storage/urlPathnya"
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	destPath := filepath.Join(destDir, filepath.Base(path))
	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destFile.Close()

	buffer := make([]byte, 1024*1024)
	if _, err := io.CopyBuffer(destFile, srcFile, buffer); err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	return nil
}
