package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/programming02/osg/api/controller"
	"github.com/programming02/osg/config"
	"github.com/programming02/osg/pkg"
	"github.com/programming02/osg/router"
	"github.com/programming02/osg/storage"
	"net/http"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		return
	}
	fmt.Println(cfg)

	db, err := pkg.Connect(cfg)
	if err != nil {
		panic(err)
	}
	storage := storage.New(db)
	a := controller.NewApi(storage)
	e := router.InitRouter(a)
	http.ListenAndServe("localhost:8080", e)

}
