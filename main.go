package main

import (
	"github.com/programming02/osg/api/controller"
	"github.com/programming02/osg/config"
	"github.com/programming02/osg/pkg"
	"github.com/programming02/osg/repository"
	"github.com/programming02/osg/router"
	"net/http"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	db, err := pkg.Connect(cfg)
	if err != nil {
		panic(err)
	}

	s := repository.New(db)
	a := controller.NewRepo(s)
	e := router.InitRouter(a)
	http.ListenAndServe("localhost:8080", e)
}
