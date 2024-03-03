package store

import (
	"encoding/gob"
	"errors"
	"io/fs"
	"os"
)

type User struct {
	Email string
	Name  string
}

type Store struct {
	Users map[string]User
	path  string
}

func New(pathname string) *Store {
	return &Store{
		Users: make(map[string]User),
		path:  pathname,
	}
}

func FromFile(pathname string) (*Store, error) {
	store := New(pathname)
	f, err := os.Open(pathname)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return store, nil
		}
		return nil, err
	}
	defer f.Close()
	err = gob.NewDecoder(f).Decode(&store)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func (s *Store) Save() error {
	f, err := os.Create(s.path)
	if err != nil {
		return err
	}
	defer f.Close()
	return gob.NewEncoder(f).Encode(s)
}

func (s *Store) AddUser(id string, user User) {
	s.Users[id] = user
}

func (s *Store) GetUser(id string) (User, error) {
	user, ok := s.Users[id]
	if !ok {
		return User{}, errors.New("user id not found")
	}
	return user, nil
}

func (s *Store) RemoveUser(id string) {
	delete(s.Users, id)
}

func (s *Store) ListUsers() map[string]User {
	return s.Users
}
