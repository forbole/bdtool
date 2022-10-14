package repo

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func prepareFileDest() (string, error) {
	// Get home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error while getting home directory : %s", err)
	}

	// File destination
	dest := fmt.Sprintf("%s/temp_BD", home)

	// Remove temp_BD directory if exists
	if _, err := os.Stat(dest); !os.IsNotExist(err) {
		err := os.RemoveAll(dest)
		if err != nil {
			return "", fmt.Errorf("error while removing existing temp_BD directory : %s", err)
		}
	}

	return dest, nil
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func getFileNameFromPath(path string) string {
	filePathSlice := strings.Split(path, "/")
	return filePathSlice[len(filePathSlice)-1:][0]
}
