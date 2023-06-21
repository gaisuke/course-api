package main

import (
	"course-api/pkg/db/mysql"

	oauth "course-api/internal/oauth/injector"
	register "course-api/internal/register/injector"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := mysql.DB()

	register.InitializedService(db).Route(&r.RouterGroup)
	oauth.InitializedService(db).Route(&r.RouterGroup)

	r.Run()
}
