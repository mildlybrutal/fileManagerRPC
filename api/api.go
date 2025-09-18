package api

// represents a file with its name and conten
type File struct {
	Name     string
	Contents []byte
}

// this is used for the WriteFile RPC call.
type WriteFileArgs struct {
	Path    string
	Content []byte
}

// defn of all rpc methods I used
type FileService interface {
	ListDirectory(path string, reply *[]string) error
	ReadFile(path string, reply *[]byte) error
	WriteFile(path string, reply *bool) error
	DeleteFile(path string, reply *bool) error
}
