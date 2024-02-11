package fileio

import (
	"os"
	"testing"
)

func TestWriteToFile_WriteGivenDataToFile(t *testing.T) {
	fileIO := NewFileIO()
	path := t.TempDir() + "/test.txt"
	want := []byte("Hello World!")
	err := fileIO.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != string(want) {
		t.Fatalf("got %s, want %s", got, want)
	}

}

func TestWriteToFile_CheckPermissionAfterWrting(t *testing.T) {
	path := t.TempDir() + "/test.txt"
	want := []byte("Hello World!")
	err := os.WriteFile(path, []byte(want), 0o644)
	if err != nil {
		t.Fatal(err)
	}
	fileIO := NewFileIO()
	err = fileIO.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	state, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if state.Mode().Perm() != 0o600 {
		t.Fatalf("got %v, want %v", state.Mode().Perm(), 0o600)
	}

}
