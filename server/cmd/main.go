package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/yamato0204/cache-server-golang/cmd/di"
	"github.com/yamato0204/cache-server-golang/cmd/di/provider"
)

func main() {
	if err := start(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)

}

func start() error {

	var gameServer *provider.GameServer
	var cleanup func()
	var err error

	gameServer, cleanup, err = di.Inject()
	if err != nil {
		return err
	}
	defer cleanup()

	e := echo.New()
	gameServer.Router.RegisterRoutes(e)
	err = gameServer.Db.Migrate()
	if err != nil {
		return err
	}
	e.Logger.Fatal(e.Start(":8080"))
	return nil
}
