package util

import (
	"fmt"
	"path/filepath"
	"strings"
)

func GetNewImageFilename(srcFilename, suffix string) (string, error) {
	if strings.TrimSpace(srcFilename) == "" {
		return "", fmt.Errorf("filename cannot be an empty string")
	}
	if strings.TrimSpace(suffix) == "" {
		suffix = "new"
	}
	ext := filepath.Ext(srcFilename)
	fileNameWithExt := filepath.Base(srcFilename)
	fileName := strings.TrimSuffix(fileNameWithExt, ext)
	dir := strings.TrimSuffix(srcFilename, fileNameWithExt)

	return dir + fileName + "_" + suffix + ext, nil
}
