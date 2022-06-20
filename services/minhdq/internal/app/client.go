package app

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"log"
	"time"

	"minhdq/internal/authentication"
)

func NewAuthenClient(ctx context.Context, raddr string, laddr string) error {
	conn, err := grpc.Dial("localhost:8888", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to register server: %v", err)
	}

	header := metadata.New(map[string]string{"content-type": "application/grpc"})

	defer conn.Close()
	c := authentication.NewResgisterClient(conn)
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()
	_, err = c.Resgister(ctx, &authentication.RegisterModel{
		LoginID:   "1",
		Password:  "Hello",
		LastName:  "Duong",
		FirstName: "Minh",
	}, grpc.Header(&header))

	if err != nil {
		log.Fatalf("could not resgister: %v", err)
	}

	log.Printf("Register success")

	conn, err = grpc.Dial("localhost:8889", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("did not connect to login server: %v", err)
	}

	cd := authentication.NewLoginClient(conn)

	log.Printf("Login with LoginID 1")

	r, err := cd.Login(ctx, &authentication.LoginModel{
		LoginID:  "1",
		Password: "Hello",
	})
	if err != nil {
		log.Fatalf("could not login: %v", err)
	}

	log.Printf("Login success")

	log.Printf("JWT return: %s", r.GetJWT())

	log.Printf("Authorization")

	_, err = cd.CheckAuthen(ctx, r)

	if err != nil {
		log.Fatalf("could not authori: %v", err)
	}

	log.Printf("Author success")

	return nil
}
