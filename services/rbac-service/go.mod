module services.rbac-service

go 1.16

require (
	github.com/getkin/kin-openapi v0.97.0
	github.com/gin-gonic/gin v1.7.7
	github.com/golang-jwt/jwt/v4 v4.4.1
	github.com/golang/protobuf v1.5.2
	github.com/rubenv/sql-migrate v0.0.0-20211023115951-9f02b1e13857
	github.com/spf13/cobra v1.2.1
	github.com/ziutek/mymysql v1.5.4 // indirect
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
	gorm.io/gorm v1.22.4
	services.core-service v0.0.1
)

replace services.core-service v0.0.1 => ../core-service
