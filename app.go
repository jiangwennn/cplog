package main

import (
	"context"
	"fmt"
	"os/user"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (* App) storeDir() string {
	dirName := ".cplog/"
	u, err := user.Current()
	if err != nil {
		return dirName
	}
	return u.HomeDir + "/" + dirName 
}

// Greet returns a greeting for the given name
func (a *App) GreetMe(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}