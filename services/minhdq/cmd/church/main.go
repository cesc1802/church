package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sort"
	"sync"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"

	"minhdq/internal/app"
	"minhdq/internal/config"
	"minhdq/internal/persistence"
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
		{
			Name:   "authen",
			Usage:  "Serve Authentication server",
			Action: AuthenServer,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "laddr",
					Aliases: []string{"login-address"},
					Value:   "0.0.0.0:8081",
				},
				&cli.StringFlag{
					Name:    "raddr",
					Aliases: []string{"register-address"},
					Value:   "0.0.0.0:8082",
				},
			},
		},
		{
			Name:   "client",
			Usage:  "Client Authentication server",
			Action: AuthenClient,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "laddr",
					Aliases: []string{"login-address"},
					Value:   "127.0.0.1:8081",
				},
				&cli.StringFlag{
					Name:    "raddr",
					Aliases: []string{"register-address"},
					Value:   "127.0.0.2:8082",
				},
			},
		},
		{
			Name:   "chat",
			Usage:  "Chatting  server",
			Action: ChattingDemo,
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

	err := persistence.LoadUserGrouprRespositorySql(ctx)
	if err != nil {
		return err
	}

	return app.Serve(ctx, c.String("addr"))
}

func ChattingDemo(c *cli.Context) error {
	err := persistence.LoadRoomChatRepoRespositoryMem(c.Context)
	if err != nil {
		return err
	}
	return app.GetChatCommand().Execute()
}

func AuthenServer(c *cli.Context) error {
	ctx := c.Context
	err := persistence.LoadAccountRespositoryMem(ctx)
	if err != nil {
		return err
	}

	errChan := make(chan error, 1)
	defer close(errChan)

	go func() {
		errChan <- app.ServeLoginServer(ctx, c.String("laddr"))
	}()

	go func() {
		errChan <- app.ServeRegisServer(ctx, c.String("raddr"))
	}()

	select {
	case err := <-errChan:
		return err
	}
}

func AuthenClient(c *cli.Context) error {
	ctx := c.Context
	return app.NewAuthenClient(ctx, c.String("raddr"), c.String("laddr"))
}
