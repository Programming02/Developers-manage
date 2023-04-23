package controller

import "github.com/programming02/osg/repository"

type Api struct {
	Repo repository.Repository
}

func NewRepo(repo repository.Repository) *Api {
	return &Api{
		Repo: repo,
	}
}
