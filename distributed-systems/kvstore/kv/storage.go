package kv

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"google.golang.org/protobuf/proto"
)

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
	}
	return err
}

func (s *Storage) Get(key string) (string, error) {
	var ic ItemCollection
	err := s.read(&ic)
	if err != nil {
		return "", err
	}
	for _, i := range ic.Items {
		if i.Key == key {
			return i.Value, nil
		}
	}
	return "", fmt.Errorf("no entry for %s found", key)
}

func (s *Storage) Set(key, value string) error {
	var ic ItemCollection
	err := s.read(&ic)
	if err != nil {
		return err
	}
	item := Item{Key: key, Value: value}
	ic.Items = append(ic.Items, &item)
	out, err := proto.Marshal(&ic)
	if err != nil {
		return fmt.Errorf("failed to serialize data %s: %s", &ic, err)
	}
	err = ioutil.WriteFile(s.path, out, fs.FileMode(uint32(0600)))
	if err != nil {
		return fmt.Errorf("failed to write file %s: %s", s.path, err)
	}
	return nil
}

func (s *Storage) read(ic *ItemCollection) error {
	data, err := ioutil.ReadFile(s.path)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %s", s.path, err)
	}
	return proto.Unmarshal(data, ic)
}