package fileio

import "os"

type IFileIO interface {
	WriteToFile(path string, data []byte) error
	ReadFromFile(path string) ([]byte, error)
}

type FileIO struct {
	//FileName string
}

func NewFileIO() *FileIO {
	return &FileIO{}
}

func (f *FileIO) WriteToFile(path string, data []byte) error {
	// we could have opened the file for writing, deferred closing it, and written the
	// byte data to it, but os.WriteFile does all this with one function call. It also creates
	// the file if it doesn’t exist, and truncates it if it does
	// 0o600 is the file permission, which is octal for 0600, or 0b1100000000 in binary
	// files has some metadata bits governing permissions who can read, write, and execute the file
	// 0o is the prefix for octal numbers, 6 is the sum of 4 (read) and 2 (write) for the owner of the file
	// 0 is the sum of 0 (read), 0 (write), and 0 (execute) for the group of the file,
	// and the last 0 is the sum of 0 (read), 0 (write), and 0 (execute) for everyone else
	err := os.WriteFile(path, data, 0o600)
	if err != nil {
		return err
	}
	return os.Chmod(path, 0o600)
}

func (f *FileIO) ReadFromFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}
