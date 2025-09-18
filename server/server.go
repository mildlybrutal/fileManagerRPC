package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"

	"github.com/fileManagerRPC/api"
)

// this struct binds all the RPC methods
type FileServer struct{}

// ls
func (f *FileServer) ListDirectory(path string, reply *[]string) error {
	entries, err := os.ReadDir(path)
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

// cat

func (f *FileServer) ReadFile(path string, reply *[]byte) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	*reply = content

	return nil
}

// WriteFileArgs cuz it wraps all arguments into a single struct which is in the api.

func (f *FileServer) WriteFile(args api.WriteFileArgs, reply *bool) error {
	err := os.WriteFile(args.Path, args.Content, 0644)
	if err != nil {
		*reply = false
		return err
	}

	*reply = true
	return nil
}

// "rm"
func (f *FileServer) DeleteFile(path string, reply *bool) error {
	err := os.Remove(path)
	if err != nil {
		*reply = false
		return err
	}

	*reply = true
	return nil
}

func main() {
	fs := new(FileServer)
	rpc.Register(fs)

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on port 8080")
	rpc.Accept(listener)
}
