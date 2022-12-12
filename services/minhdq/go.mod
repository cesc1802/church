module minhdq

go 1.13

require (
	cloud.google.com/go/kms v1.4.0 // indirect
	github.com/Masterminds/squirrel v1.5.3
	github.com/RichardKnop/machinery v1.10.6
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/gin-gonic/gin v1.7.7
	github.com/go-chi/chi/v5 v5.0.7
	github.com/golang-jwt/jwt/v4 v4.4.1
	github.com/golang/protobuf v1.5.2
	github.com/jackc/pgx/v4 v4.16.1
	github.com/joho/godotenv v1.4.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/ory/fosite v0.42.2 // indirect
	github.com/spf13/cobra v1.5.0
	github.com/urfave/cli/v2 v2.8.1
	google.golang.org/grpc v1.47.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/olahol/melody.v1 v1.0.0-20170518105555-d52139073376
	services.core-service v0.0.1
)

replace services.core-service v0.0.1 => ../core-service

//replace gopkg.in/olahol/melody.v1 v1.0.0-20170518105555-d52139073376 => /home/minhdq/go/src/gopkg.in/olahol/melody.v1
