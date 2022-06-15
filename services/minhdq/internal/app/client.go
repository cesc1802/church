package app

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"minhdq/internal/authentication"
)

func NewAuthenClient(ctx context.Context, raddr string, laddr string) error {
	conn, err := grpc.Dial(raddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to register server: %v", err)
	}

	defer conn.Close()
	c := authentication.NewResgisterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.Resgister(ctx, &authentication.RegisterModel{
		LoginID:   "1",
		Password:  "Hello",
		LastName:  "Duong",
		FirstName: "Minh",
	})

	if err != nil {
		log.Fatalf("could not resgister: %v", err)
	}

	log.Printf("Register success")

	conn, err = grpc.Dial(laddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

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
