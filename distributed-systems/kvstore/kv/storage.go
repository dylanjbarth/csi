package kv

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
)

type DataContainer map[string]string

type Storage struct {
	path string
}

func NewStorage(path string, clean bool) *Storage {
	s := Storage{path}
	s.initStorage(clean)
	return &s
}

func (s *Storage) initStorage(clean bool) error {
	_, err := os.Stat(s.path)
	if os.IsNotExist(err) || clean {
		_, err = os.Create(s.path)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(s.path, []byte("{}"), fs.FileMode(uint32(0600)))
	}
	return err
}

func (s *Storage) Get(key string) (string, error) {
	var container DataContainer
	err := s.read(&container)
	if err != nil {
		return "", err
	}
	return container[key], nil
}

func (s *Storage) Set(key, value string) error {
	var container DataContainer
	err := s.read(&container)
	if err != nil {
		return err
	}
	container[key] = value
	out, err := json.Marshal(container)
	if err != nil {
		return fmt.Errorf("failed to serialize data %s: %s", container, err)
	}
	err = ioutil.WriteFile(s.path, out, fs.FileMode(uint32(0600)))
	if err != nil {
		return fmt.Errorf("failed to write file %s: %s", s.path, err)
	}
	return nil
}

func (s *Storage) read(container *DataContainer) error {
	data, err := ioutil.ReadFile(s.path)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %s", s.path, err)
	}
	return json.Unmarshal(data, &container)
}
