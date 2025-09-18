package main

import (
	"os"
	"github.com/fileManagerRPC/api"
)

type FileServer struct{}

func (f *FileServer) ListDirectory(path string, reply *[]string) error {
	entries, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var files []string

	for _, entry := range entries {
		files = append(files, entry.Name())
	}

	*reply = files
	return nil
}

func (f *FileServer) ReadFile(path string, reply *[]byte) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	*reply = content

	return nil
}

func (f *FileServer) WriteFile(file api.File, reply *bool) error {
	err := os.WriteFile(file.Name, file.Content, 0644)
	if err != nil {
		*reply = false
		return err
	}

	*reply = true
	return nil
}

func main() {

}
