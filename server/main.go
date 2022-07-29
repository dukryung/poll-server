package main

import (
	"fmt"
	"os"
	"os/signal"
	"poll-server/server/app"
	"syscall"
)

func main() {
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!")

	app := app.NewApp()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	app.RunServers()

	<-quit
	app.CloseServers()
}
