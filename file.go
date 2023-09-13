package utils

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

// FileExists check if file exists
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// IsDirectory check if path is directory
func IsDirectory(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
}

// FindFile find files in directory with pattern
func FindFile(dir string, pattern string) []string {
	var files []string
	filepath.Walk(dir, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(pattern, f.Name())
			if err == nil && r {
				files = append(files, path)
			}
		}
		return nil
	})
	return files
}

// ClearDirectory delete all files and sub-directory in directory
func ClearDirectory(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

// GetSubDirectory get list of sub directories
func GetSubDirectory(dir string) ([]string, error) {
	var res []string
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if f.IsDir() {
			res = append(res, f.Name())
		}
	}
	return res, nil
}

// CreateDirectory create nested directory
func CreateDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// DetectMime detect file mime info from content
func DetectMime(data []byte) *mimetype.MIME {
	return mimetype.Detect(data)
}

// Extension get file extension
func Extension(file string) string {
	return strings.ToLower(filepath.Ext(file))
}

// NumberedFile generate unique numbered file until 10000000, e.g. file.txt file-1.txt, file-2.txt
func NumberedFile(dir, name, file string) (string, error) {
	// extract extension
	ext := strings.TrimSpace(strings.TrimLeft(filepath.Ext(file), "."))
	if ext != "" {
		ext = "." + ext
	}
	// check if current file valid
	if exists, err := FileExists(path.Join(dir, name+ext)); err != nil {
		return "", err
	} else if !exists {
		return name + ext, nil
	} else {
		for i := 1; i < 10000000; i++ {
			_name := fmt.Sprintf("%s-%d%s", name, i, ext)
			if exists, err := FileExists(path.Join(dir, _name)); err != nil {
				return "", err
			} else if !exists {
				return _name, nil
			}
		}
	}
	return "", errors.New("try 10000000 number failed")
}
