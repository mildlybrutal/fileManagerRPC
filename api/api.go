package api

type File struct {
	Name     string
	Contents []byte
}

type WriteFileArgs struct {
	Path    string
	Content []byte
}

type FileService interface {
	ListDirectory(path string, reply *[]string) error
	ReadFile(path string, reply *[]byte) error
	WriteFile(path string, reply *bool) error
	DeleteFile(path string, reply *bool) error
}
