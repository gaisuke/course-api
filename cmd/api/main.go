package main

import (
	"course-api/pkg/db/mysql"

	admin "course-api/internal/admin/injector"
	forgotPassword "course-api/internal/forgot_password/injector"
	oauth "course-api/internal/oauth/injector"
	register "course-api/internal/register/injector"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := mysql.DB()

	admin.InitializedService(db).Route(&r.RouterGroup)
	register.InitializedService(db).Route(&r.RouterGroup)
	oauth.InitializedService(db).Route(&r.RouterGroup)
	forgotPassword.InitializedService(db).Route(&r.RouterGroup)

	r.Run()
}
