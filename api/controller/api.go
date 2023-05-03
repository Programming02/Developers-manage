package controller

import "github.com/programming02/osg/storage"

type Api struct {
	storage storage.IStorage
}

func NewApi(iStorage storage.IStorage) *Api {
	return &Api{
		storage: iStorage,
	}
}
