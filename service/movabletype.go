package service

import (
	"bytes"
	"io/ioutil"

	"github.com/yamadatt/movabletype"
)

type MovableType interface {
	Parse(path string) ([]*movabletype.Entry, error)
}

type MovableTypeImpl struct {
}

func NewMovableType() MovableType {
	return &MovableTypeImpl{}
}

func (s *MovableTypeImpl) Parse(path string) ([]*movabletype.Entry, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(data)
	return movabletype.Parse(reader)
}
