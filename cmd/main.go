package main

import (
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/programming02/osg/api/controller"
	"github.com/programming02/osg/config"
	"github.com/programming02/osg/pkg"
	"github.com/programming02/osg/router"
	"net/http"
)

func main() {

	cfg := config.Load()

	db, err := pkg.Connect(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	a := controller.NewApi(controller.NewAdminService(db), controller.NewProgrammerService(db), controller.NewRegisterService(db))
	e := router.InitRouter(a)
	http.ListenAndServe("localhost:8080", e)
}
