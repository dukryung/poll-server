package app

import "github.com/dukryung/poll-server/server/types"

type App struct {
	Servers []types.Server
}

func NewApp() *App {

	return &App{}
}