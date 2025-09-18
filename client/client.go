package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"

	"github.com/fileManagerRPC/api"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Error connecting: ", err)
		return
	}

	defer client.Close()

	fmt.Println("RPC File Manager Client connected.")
	fmt.Println("Available commands: ls <path>, cat <path>, upload <local_file> <remote_path>, download <remote_file> <local_path>, rm <path>, exit")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		args := parts[1:]

		switch command {
		case "ls":
			if len(args) != 1 {
				fmt.Println("Usage: ls <path>")
				continue
			}
			var reply []string

			err := client.Call("FileServer.ListDirectory", args[0], &reply)

			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}

			for _, file := range reply {
				fmt.Println(file)
			}
		case "cat":
			if len(args) != 1 {
				fmt.Println("Usage: cat <path>")
				continue
			}

			var content []byte

			err := client.Call("FileServer.ReadFile", args[0], &content)

			if err != nil {
				fmt.Printf("Error: %v\n", err)
				continue
			}

			fmt.Println(string(content))
		case "upload":
			if len(args) != 2 {
				fmt.Println("Usage: upload <local_file> <remote_path>")
				continue
			}
			localPath := args[0]
			remotePath := args[1]

			content, err := os.ReadFile(localPath)

			if err != nil {
				fmt.Printf("Error reading local file: %v\n", err)
				continue
			}

			var success bool

			args := api.WriteFileArgs{Path: remotePath, Content: content}

			err = client.Call("FileServer.WriteFile", args, &success)

			if err != nil {
				fmt.Printf("Error uploading file: %v\n", err)
				continue
			}

			if success {
				fmt.Println("File uploaded successfully.")
			} else {
				fmt.Println("File upload failed.")
			}
		case "rm":
			if len(args) != 1 {
				fmt.Println("Usage: rm <path>")
				continue
			}
			var success bool

			err := client.Call("FileServer.DeleteFile", args[0], &success)

			if err != nil {
				fmt.Printf("Error deleting file: %v\n", err)
				continue
			}
			if success {
				fmt.Println("File deleted successfully.")
			} else {
				fmt.Println("File deletion failed.")
			}
		case "download":
			if len(args) != 2 {
				fmt.Println("Usage: download <remote_file> <local_path>")
				continue
			}

			remotePath := args[0]
			localPath := args[1]

			var content []byte

			err := client.Call("FileServer.ReadFile", remotePath, &content)
			if err != nil {
				fmt.Printf("Error downloading file: %v\n", err)
				continue
			}

			err = os.WriteFile(localPath, content, 0644)
			if err != nil {
				fmt.Printf("Error writing to local file: %v\n", err)
				continue
			}
			fmt.Println("File downloaded successfully.")

		case "exit":
			fmt.Println("Exiting.")
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}
