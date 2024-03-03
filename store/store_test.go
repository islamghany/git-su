package store

import (
	"os"
	"testing"
)

func TestStoreNotExsitingFile(t *testing.T) {
	// TestStoreNotExsitingFile tests the FromFile function with a non-existing file
	t.Run("TestStoreNotExsitingFile", func(t *testing.T) {
		store, err := FromFile("not-existing-file")
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
		if len(store.Users) != 0 {
			t.Errorf("Expected empty store, but got %v", store.Users)
		}
	})

	// ErrorsWhenPathUnreadable
	t.Run("ErrorsWhenPathUnreadable", func(t *testing.T) {
		path := t.TempDir() + "/unreadable-file"
		if _, err := os.Create(path); err != nil {
			t.Fatal(err)
		}
		if err := os.Chmod(path, 0); err != nil {
			t.Fatal(err)
		}
		_, err := FromFile(path)
		if err == nil {
			t.Error("Expected error, but got nil")
		}
	})
}
