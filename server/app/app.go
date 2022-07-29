package app

import (
	"github.com/dukryung/poll-server/server/poll"
	"github.com/dukryung/poll-server/server/types"
)

type App struct {
	Servers []types.Server
}

func NewApp() *App {
	app := &App{}
	pollServer := poll.NewServer()

	app.Servers = append(app.Servers, pollServer)

	return app
}

func (a *App) RunServers() {
	for _, server := range a.Servers {
		server.Run()
	}
}

func (a *App) CloseServers() {
	for _, server := range a.Servers {
		server.Close()
	}
}
