package main

import (
	"context"
	"github.com/ory/fosite"
	"github.com/urfave/cli/v2"
	"log"
	"oauth/app"
	"oauth/config"
	"oauth/persistence"
	"oauth/service"
	"oauth/ulti"
	"os"
	"os/signal"
	"sort"
	"sync"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	app := cli.NewApp()
	app.Name = "Church API User Group"
	app.Usage = "CRUD User Group"
	app.Version = "Pre_Beta"

	app.Compiled = time.Now()
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "env",
			Aliases: []string{"e"},
			Value:   "./configs/.env",
			Usage:   "set path to enviroment file",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:   "serve",
			Usage:  "Serve server",
			Action: Serve,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "addr",
					Aliases: []string{"address"},
					Value:   "0.0.0.0:8080",
				},
			},
		},
	}

	app.Before = func(c *cli.Context) error {
		return godotenv.Load(c.String("env"))
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	ctx, cancel := context.WithCancel(context.Background())

	endSignal := make(chan os.Signal, 1)
	signal.Notify(endSignal, syscall.SIGINT, syscall.SIGTERM)

	wg := &sync.WaitGroup{}

	errChan := make(chan error, 1)

	wg.Add(1)

	go func(ctx context.Context, errChan chan error) {
		defer wg.Done()
		err := app.RunContext(ctx, os.Args)
		errChan <- err
	}(ctx, errChan)

	select {
	case sign := <-endSignal:
		log.Printf("Shutting down. reason:", sign)
	case err := <-errChan:
		if err == nil {
			break
		}
		log.Println("encountered error", err)
		return
	}

	cancel()
	wg.Wait()
}

func Serve(c *cli.Context) error {
	if err := config.LoadFromEnv(); err != nil {
		return err
	}

	ctx := c.Context

	err := persistence.LoadUserRepoRespositoryMem(ctx)
	if err != nil {
		return err
	}

	err = persistence.LoadClientRepoRespositoryMem(ctx)
	if err != nil {
		return err
	}

	storage := service.NewStorage()
	privateKey, err := ulti.PrivateKey(ctx)

	if err != nil {
		return err
	}

	secret := []byte("why_cant_this_be_empty")

	configOauth := &fosite.Config{
		AccessTokenLifespan: time.Minute * 30,
		GlobalSecret:        secret,
		// ...
	}

	service.InitProvider(ctx, configOauth, storage, privateKey)

	return app.Serve(ctx, c.String("addr"))
}
