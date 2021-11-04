package kv

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

type Storage struct {
	path string
}

func NewStorage(path string, clean bool) *Storage {
	s := Storage{path}
	err := s.initStorage(clean)
	if err != nil {
		log.Fatalf("failed to init storage; %s", err)
	}
	return &s
}

func (s *Storage) initStorage(clean bool) error {
	log.Printf("Initializing storage file %s", s.path)
	_, err := os.Stat(s.path)
	if os.IsNotExist(err) || clean {
		log.Printf("path doesn't exist. attempting to create it")
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
	match := -1
	for idx, i := range ic.Items {
		if i.Key == key {
			match = idx
			break
		}
	}
	if match >= 0 {
		ic.Items[match].Value = value
	} else {
		item := Item{Key: key, Value: value}
		ic.Items = append(ic.Items, &item)
	}
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
