package service

import (
	"context"
	"fmt"
	"go.uber.org/dig"
	"meme_coin_api/service/api"
	"meme_coin_api/service/controller/memeCoinCtrl"
	"meme_coin_api/service/internal/config"
	"meme_coin_api/service/internal/database"
	"meme_coin_api/service/internal/flags"
	"net/http"
)

func MemeCoin() Service {
	once.Do(func() {
		srv = &memeCoin{}
	})

	return srv
}

type memeCoin struct{}

func (srv *memeCoin) Run() {
	flags.Parse()

	container := dig.New()

	srv.provideConfig(container)

	srv.provideService(container)

	srv.provideController(container)

	srv.invokeApiRoutes(container)

	if err := container.Invoke(srv.run); err != nil {
		panic(err)
	}
}

func (srv *memeCoin) provideConfig(container *dig.Container) {
	if err := container.Provide(config.NewMemeCoin); err != nil {
		panic(err)
	}
}

func (srv *memeCoin) provideService(container *dig.Container) {
	if err := container.Provide(func() context.Context {
		return context.TODO()
	}); err != nil {
		panic(err)
	}

	if err := container.Provide(database.NewMemeCoin); err != nil {
		panic(err)
	}

	if err := container.Provide(api.NewServer); err != nil {
		panic(err)
	}

	if err := container.Provide(api.NewGinEngine); err != nil {
		panic(err)
	}

	if err := container.Provide(api.NewRouterRoot); err != nil {
		panic(err)
	}
}

func (srv *memeCoin) provideController(container *dig.Container) {
	if err := container.Provide(memeCoinCtrl.New); err != nil {
		panic(err)
	}
}

func (srv *memeCoin) invokeApiRoutes(container *dig.Container) {
	if err := container.Invoke(api.NewBasic); err != nil {
		panic(err)
	}

	if err := container.Invoke(api.NewMemeCoin); err != nil {
		panic(err)
	}
}

func (srv *memeCoin) run(server *http.Server) {
	fmt.Printf("Meme Coin API starts at %s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
