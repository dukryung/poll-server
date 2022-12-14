package main

import (
	"github.com/dukryung/poll-server/server/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {


	app := app.NewApp()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	app.RunServers()

	<-quit
	app.CloseServers()
}
