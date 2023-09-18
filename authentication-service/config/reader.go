package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// read file from a disk
func ReadFile(fileName string) *[]byte {
	// identify current file directory
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	content, err := os.ReadFile(dir + "/ops/" + fileName)
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}

	return &content
}