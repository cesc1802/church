module services.rbac-service

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7 // indirect
	github.com/rubenv/sql-migrate v0.0.0-20211023115951-9f02b1e13857
	github.com/spf13/cobra v1.2.1
	github.com/ziutek/mymysql v1.5.4 // indirect
	gorm.io/gorm v1.22.4
	services.core-service v0.0.1
)

replace services.core-service v0.0.1 => ../core-service
