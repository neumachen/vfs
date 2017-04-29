package vfs

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Store ...
type Store interface {
	Get(key string) (string, error)
}

type store struct {
	files map[string]string
}

// Get takes a schema string and searches the *Schemas and returns the json
// schema if found. It returns an error if the schema is not located.
func (s *store) Get(file string) (string, error) {
	f, ok := s.files[file]
	if !ok {
		return "", fmt.Errorf("file: '%s' could not be found", file)
	}

	return f, nil
}

// LoadFiles loads the file(s) in a given folder from the given string that is
// passed. If the path supplied is a directory it will walk through the entire
// directory and load the json files into the *Store.  If the file is not a
// directory it will simply load that single file into the Store. If an error
// occurs anytime during the process an error is returned.
func LoadFiles(folder string) (Store, error) {
	fdir, err := os.Open(folder)
	if err != nil {
		return nil, err
	}
	defer fdir.Close()

	finfo, err := fdir.Stat()

	if err != nil {
		return nil, err
	}

	s := &store{files: make(map[string]string)}

	switch mode := finfo.Mode(); {
	case mode.IsDir():
		files, err := lsFiles(fdir.Name())
		if err != nil {
			return nil, err
		}
		for _, f := range files {
			file, err := os.Open(f)
			if err != nil {
				return nil, err
			}
			err = loadFile(s, file)
			defer file.Close()
			if err != nil {
				return nil, err
			}
		}
	case mode.IsRegular():
		err := loadFile(s, fdir)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

// loadFile ...
func loadFile(s *store, file *os.File) error {
	base := filepath.Base(file.Name())

	// store the file in the files using the base name only (not
	// path/filename)
	s.files[base] = scanFile(file)

	return nil
}

// lsFiles takes a searchDir string and RECURSIVELY collects the file names
// inside the directory.
func lsFiles(searchDir string) ([]string, error) {
	fileList := []string{}
	_ = filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})

	return fileList, nil
}

// scanFile ...
func scanFile(file *os.File) string {
	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))

	var err error
	var eol bool
	var chunk []byte
	var strArray []string

	for {
		if chunk, eol, err = reader.ReadLine(); err != nil {
			break
		}

		buffer.Write(chunk)
		if !eol {
			strArray = append(strArray, buffer.String())
			buffer.Reset()
		}
	}

	return strings.Join(strArray[:], "\n")
}
