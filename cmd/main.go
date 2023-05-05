package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/programming02/osg/api/controller"
	"github.com/programming02/osg/config"
	"github.com/programming02/osg/pkg"
	"github.com/programming02/osg/router"
	"github.com/programming02/osg/storage"
	"net/http"
)

func main() {
	r := gin.Default()
	cfg, err := config.Load()
	if err != nil {
		return
	}
	fmt.Println(cfg)
	db, err := pkg.Connect(cfg)
	if err != nil {
		panic(err)
	}

	storages := storage.New(db)
	a := controller.NewApi(storages)
	router.AdminRouter(a, r)
	router.ProgrammerRouter(a, r)
	router.RegisterROuter(a, r)
	err = http.ListenAndServe("localhost:8080", r)
	if err != nil {
		panic(err)
	}

}
